package model

import (
	pb "inventory/proto"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	ID         string
	Status     pb.Status
	Items      []*Item
	CreatedAt  time.Time
	CapturedAt time.Time
}

type Item struct {
	ID         string
	OrderID    string
	Order      *Order
	ProductID  string
	Quantity   uint32
	Fulfilable bool
}

type Topup struct {
	ID        string
	Quantity  int32
	ProductID string
	CreatedAt time.Time
}

type Product struct {
	ID             string
	UnitsInStock   int32
	UnitsReserved  int32
	UnitsAvailable int32
}

// CreateOrder in the database
func CreateOrder(db *gorm.DB, o *Order) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		// get the units available for all the products
		productIDs := make([]string, len(o.Items))
		for i, item := range o.Items {
			productIDs[i] = item.ProductID
		}
		counts, err := unitsAvailableForProducts(tx, productIDs)
		if err != nil {
			return err
		}

		// loop over the items and only mark them as fulfilable if there is more units available than
		// there is in the order
		for _, i := range o.Items {
			i.Fulfilable = counts[i.ProductID] >= int32(i.Quantity)
		}

		// write the order & the items to the store
		if err := db.Create(o).Error; err != nil {
			return err
		}

		return nil
	})
	// ignore duplicate index constraint error (the order already exists)
	if err != nil && strings.Contains(err.Error(), "orders_pkey") {
		return nil
	}
	return err
}

// ReadOrder from the database
func ReadOrder(db *gorm.DB, id string) (*Order, error) {
	var o Order
	if err := db.Where(&Order{ID: id}).First(&o).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

// Pick an order
func Pick(db *gorm.DB, id string) error {
	return db.Where(&Order{ID: id}).Updates(&Order{Status: pb.Status_STATUS_PICKED}).Error
}

// Void an order
func Void(db *gorm.DB, id string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// update the orders status to void
		if err := db.Exec("UPDATE orders SET status = ? WHERE id = ?", int32(pb.Status_STATUS_VOIDED), id).Error; err == gorm.ErrRecordNotFound {
			return nil
		} else if err != nil {
			return err
		}

		// load the items
		var items []Item
		if err := tx.Raw("SELECT product_id FROM items WHERE order_id = ? FOR UPDATE", id).Scan(&items).Error; err != nil {
			return err
		}

		// reconcile each product id
		for _, i := range items {
			if err := reconcile(tx, i.ProductID); err != nil {
				return err
			}
		}

		return nil
	})
}

// CreateTopups multiple products
func CreateTopups(db *gorm.DB, topups []*pb.Topup) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, t := range topups {
			result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&Topup{ID: t.Id, ProductID: t.ProductId, Quantity: t.Quantity})
			if result.Error != nil {
				return result.Error
			} else if result.RowsAffected == 0 {
				continue
			}

			if err := reconcile(tx, t.ProductId); err != nil {
				return err
			}
		}
		return nil
	})
}

// OverrideInventory for a product
func OverrideInventory(db *gorm.DB, id, productID string, quanity int32) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// get the units available for this product
		prev, err := unitsInStock(tx, productID)
		if err != nil {
			return err
		}

		// calculate the difference
		change := quanity - prev
		if change == 0 {
			return nil
		}

		// create the topup to reconcile the difference
		result := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&Topup{ID: id, ProductID: productID, Quantity: change})
		if result.Error != nil {
			return result.Error
		} else if result.RowsAffected == 0 {
			return nil
		}

		return reconcile(tx, productID)
	})
}

func ListProducts(db *gorm.DB, ids []string) ([]*Product, error) {
	var products []*Product

	err := db.Transaction(func(tx *gorm.DB) error {
		// get all the items which have been picked
		var pickedQuery *gorm.DB
		if ids == nil {
			pickedQuery = tx.Raw("SELECT SUM(i.quantity), i.product_id FROM items i, orders o WHERE o.id = i.order_id AND o.status = ? GROUP BY i.product_id", int(pb.Status_STATUS_PICKED))
		} else {
			pickedQuery = tx.Raw("SELECT SUM(i.quantity), i.product_id FROM items i, orders o WHERE o.id = i.order_id AND o.status = ? AND i.product_id IN (?) GROUP BY i.product_id", int(pb.Status_STATUS_PICKED), ids)
		}
		var unitsPicked []struct {
			ProductID string
			Sum       int32
		}
		if err := pickedQuery.Scan(&unitsPicked).Error; err != nil {
			return err
		}

		// get all the items which have been reserved
		var reservedQuery *gorm.DB
		if ids == nil {
			reservedQuery = tx.Raw("SELECT SUM(i.quantity), i.product_id FROM items i, orders o WHERE o.id = i.order_id AND o.status = ? GROUP BY i.product_id", int(pb.Status_STATUS_RESERVED))
		} else {
			reservedQuery = tx.Raw("SELECT SUM(i.quantity), i.product_id FROM items i, orders o WHERE o.id = i.order_id AND o.status = ? AND i.product_id IN (?) GROUP BY i.product_id", int(pb.Status_STATUS_RESERVED), ids)
		}
		var unitsReserved []struct {
			ProductID string
			Sum       int32
		}
		if err := reservedQuery.Scan(&unitsReserved).Error; err != nil {
			return err
		}

		// get all the topups
		var topupsQuery *gorm.DB
		if ids == nil {
			topupsQuery = tx.Raw("SELECT SUM(quantity), product_id FROM topups GROUP BY product_id")
		} else {
			topupsQuery = tx.Raw("SELECT SUM(quantity), product_id FROM topups WHERE product_id IN (?) GROUP BY product_id", ids)
		}
		var topups []struct {
			ProductID string
			Sum       int32
		}
		if err := topupsQuery.Scan(&topups).Error; err != nil {
			return err
		}

		// aggregate the data
		var productsByID map[string]*Product
		if ids == nil {
			productsByID = map[string]*Product{}
		} else {
			productsByID = make(map[string]*Product, len(ids))
		}
		for _, t := range topups {
			productsByID[t.ProductID] = &Product{
				UnitsInStock: t.Sum,
			}
		}
		for _, r := range unitsPicked {
			if p, ok := productsByID[r.ProductID]; ok {
				p.UnitsInStock -= r.Sum
			} else {
				productsByID[r.ProductID] = &Product{
					UnitsInStock: -r.Sum,
				}
			}
		}
		for _, r := range unitsReserved {
			if p, ok := productsByID[r.ProductID]; ok {
				p.UnitsReserved = r.Sum
			} else {
				productsByID[r.ProductID] = &Product{
					UnitsReserved: r.Sum,
				}
			}
		}

		// populate any missing products (there could be none topped up, picked or reserved)
		for _, id := range ids {
			if _, ok := productsByID[id]; !ok {
				productsByID[id] = &Product{ID: id}
			}
		}

		// serialize the data
		products = make([]*Product, 0, len(productsByID))
		for id, product := range productsByID {
			products = append(products, &Product{
				ID:             id,
				UnitsInStock:   product.UnitsInStock,
				UnitsReserved:  product.UnitsReserved,
				UnitsAvailable: product.UnitsInStock - product.UnitsReserved,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return products, nil
}

func FulfilableOrders(db *gorm.DB) ([]string, error) {
	query := db.Raw("SELECT order_id, every(fulfilable), COUNT(order_id), MIN(created_at) FROM items i INNER JOIN orders o ON i.order_id = o.id WHERE status = ? GROUP BY order_id HAVING every(fulfilable) = TRUE ORDER BY MIN(created_at) ASC", int(pb.Status_STATUS_RESERVED))
	var result []struct {
		OrderID string
	}
	if err := query.Scan(&result).Error; err != nil {
		return nil, err
	}
	ids := make([]string, len(result))
	for i, row := range result {
		ids[i] = row.OrderID
	}
	return ids, nil
}

func UnfulfilableOrders(db *gorm.DB) ([]string, error) {
	query := db.Raw("SELECT order_id, every(fulfilable), COUNT(order_id), MIN(created_at) FROM items i INNER JOIN orders o ON i.order_id = o.id WHERE status = ? GROUP BY order_id HAVING every(fulfilable) = FALSE ORDER BY MIN(created_at) ASC", int(pb.Status_STATUS_RESERVED))
	var result []struct {
		OrderID string
	}
	if err := query.Scan(&result).Error; err != nil {
		return nil, err
	}
	ids := make([]string, len(result))
	for i, row := range result {
		ids[i] = row.OrderID
	}
	return ids, nil
}

func unitsAvailableForProducts(tx *gorm.DB, productIDs []string) (map[string]int32, error) {
	// get all the topups for this product
	var unitsIn []struct {
		Sum       int32
		ProductID string
	}
	if err := tx.Raw("SELECT SUM(quantity), product_id FROM topups WHERE product_id IN (?) GROUP BY product_id", productIDs).Scan(&unitsIn).Error; err != nil {
		return nil, err
	}

	// get all the items which have already been fulfiled
	var unitsOut []struct {
		Sum       int32
		ProductID string
	}
	if err := tx.Raw("SELECT SUM(i.quantity), product_id FROM items i, orders o WHERE o.id = i.order_id AND o.status IN (?) AND i.product_id IN (?) GROUP BY product_id",
		[]int{int(pb.Status_STATUS_PICKED), int(pb.Status_STATUS_RESERVED)}, productIDs).Scan(&unitsOut).Error; err != nil {
		return nil, err
	}

	// construct the result
	result := make(map[string]int32, len(productIDs))
	for _, row := range unitsIn {
		result[row.ProductID] += row.Sum
	}
	for _, row := range unitsOut {
		result[row.ProductID] -= row.Sum
	}
	return result, nil
}

func unitsInStock(tx *gorm.DB, productID string) (int32, error) {
	// get all the items which have already been fulfiled
	var unitsOut struct {
		Sum int32
	}
	if err := tx.Raw("SELECT SUM(i.quantity) FROM items i, orders o WHERE o.id = i.order_id AND o.status = ? AND i.product_id = ?",
		int(pb.Status_STATUS_PICKED), productID).Scan(&unitsOut).Error; err != nil {
		return 0, err
	}

	// get all the topups for this product
	var unitsIn struct {
		Sum int32
	}
	if err := tx.Raw("SELECT SUM(quantity) FROM topups WHERE product_id = ?", productID).Scan(&unitsIn).Error; err != nil {
		return 0, err
	}

	// calculate the number of units remaining
	return unitsIn.Sum - unitsOut.Sum, nil
}

func reconcile(tx *gorm.DB, productID string) error {
	// get all the items which need to be fulfilled
	var itemsToFulfil []Item
	if err := tx.Raw("SELECT i.id, i.quantity FROM items i, orders o WHERE o.id = i.order_id AND o.status NOT IN (?) AND i.product_id = ? ORDER BY o.created_at ASC", []int{
		int(pb.Status_STATUS_PICKED), int(pb.Status_STATUS_VOIDED),
	}, productID).Scan(&itemsToFulfil).Error; err != nil {
		return err
	}

	// get the units available for this product
	remaining, err := unitsInStock(tx, productID)
	if err != nil {
		return err
	}

	fulfilableItems := []string{}
	unfulfilableItems := []string{}
	for _, i := range itemsToFulfil {
		remaining -= int32(i.Quantity)
		if remaining >= 0 {
			fulfilableItems = append(fulfilableItems, i.ID)
		} else {
			unfulfilableItems = append(unfulfilableItems, i.ID)
		}
	}

	// update all the items
	if len(fulfilableItems) > 0 {
		if err := tx.Exec("UPDATE items SET fulfilable = TRUE WHERE id IN (?)", fulfilableItems).Error; err != nil {
			return err
		}
	}
	if len(unfulfilableItems) > 0 {
		if err := tx.Exec("UPDATE items SET fulfilable = FALSE WHERE id IN (?)", unfulfilableItems).Error; err != nil {
			return err
		}
	}

	return nil
}
