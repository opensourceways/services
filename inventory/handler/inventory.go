package handler

import (
	"context"

	"inventory/model"
	pb "inventory/proto"

	"github.com/micro/micro/v3/service/errors"
	"gorm.io/gorm"
)

var (
	ErrMissingOrderID        = errors.BadRequest("inventory.MissingOrderID", "Missing OrderID")
	ErrMissingItems          = errors.BadRequest("inventory.MissingItems", "One or more items are required")
	ErrMissingItemID         = errors.BadRequest("inventory.MissingItemID", "One of the items is missing an ID")
	ErrMissingItemProductID  = errors.BadRequest("inventory.MissingItemProductID", "One of the items is missing a product ID")
	ErrMissingItemQuantity   = errors.BadRequest("inventory.MissingItemQuantity", "Item quantity must be greater than zero")
	ErrMissingTopups         = errors.BadRequest("inventory.MissingTopups", "One or more topups are required")
	ErrMissingTopupID        = errors.BadRequest("inventory.MissingTopupID", "One of the topups are missing an ID")
	ErrMissingTopupQuantity  = errors.BadRequest("inventory.MissingTopupQuantity", "One of the topups are missing a quanity")
	ErrMissingTopupProductID = errors.BadRequest("inventory.MissingTopupProductID", "One of the topups are missing a product ID")
	ErrMissingID             = errors.BadRequest("inventory.MissingID", "Missing ID")
	ErrMissingProductID      = errors.BadRequest("inventory.MissingProductID", "Missing ProductID")
)

type Inventory struct {
	DB *gorm.DB
}

// Reserve inventory required to fulfil and order
func (i *Inventory) Reserve(ctx context.Context, req *pb.ReserveRequest, rsp *pb.ReserveResponse) error {
	// validate the request
	if len(req.OrderId) == 0 {
		return ErrMissingOrderID
	} else if len(req.Items) == 0 {
		return ErrMissingItems
	}
	for _, i := range req.Items {
		if err := validateItem(i); err != nil {
			return err
		}
	}

	// write the order to the database
	items := make([]*model.Item, len(req.Items))
	for i, it := range req.Items {
		items[i] = &model.Item{
			ID:        it.Id,
			Quantity:  uint32(it.Quantity),
			ProductID: it.ProductId,
		}
	}
	return model.CreateOrder(i.DB, &model.Order{
		ID:     req.OrderId,
		Status: pb.Status_STATUS_RESERVED,
		Items:  items,
	})
}

// Pick the inventory which was used to fulfil an order
func (i *Inventory) Pick(ctx context.Context, req *pb.PickRequest, rsp *pb.PickResponse) error {
	// validate the request
	if len(req.OrderId) == 0 {
		return ErrMissingOrderID
	}

	// check to see if the order already exists
	if _, err := model.ReadOrder(i.DB, req.OrderId); err != nil && err != gorm.ErrRecordNotFound {
		return errors.InternalServerError("inventory.Database", "Error connecting to the database")
	} else if err == nil {
		if cerr := model.Pick(i.DB, req.OrderId); cerr != nil {
			return errors.InternalServerError("inventory.Database", "Error connecting to the database")
		}
		return nil
	}

	// order needs to be created, validate the items
	if len(req.Items) == 0 {
		return ErrMissingItems
	}
	for _, i := range req.Items {
		if err := validateItem(i); err != nil {
			return err
		}
	}

	// write the order to the database
	items := make([]*model.Item, len(req.Items))
	for i, it := range req.Items {
		items[i] = &model.Item{
			ID:        it.Id,
			Quantity:  uint32(it.Quantity),
			ProductID: it.ProductId,
		}
	}
	return model.CreateOrder(i.DB, &model.Order{
		ID:     req.OrderId,
		Status: pb.Status_STATUS_PICKED,
		Items:  items,
	})
}

// Void an order
func (i *Inventory) Void(ctx context.Context, req *pb.VoidRequest, rsp *pb.VoidResponse) error {
	// validate the request
	if len(req.OrderId) == 0 {
		return ErrMissingOrderID
	}

	return model.Void(i.DB, req.OrderId)
}

// UnfulfilableOrders lists all the orders which cannot be fulfilled with the current inventory
func (i *Inventory) UnfulfilableOrders(ctx context.Context, req *pb.UnfulfilableOrdersRequest, rsp *pb.UnfulfilableOrdersResponse) error {
	ids, err := model.UnfulfilableOrders(i.DB)
	if err != nil {
		return err
	}
	rsp.Ids = ids
	return nil
}

// FulfilableOrders lists all the orders which can be fulfilled with the current inventory
func (i *Inventory) FulfilableOrders(ctx context.Context, req *pb.FulfilableOrdersRequest, rsp *pb.FulfilableOrdersResponse) error {
	ids, err := model.FulfilableOrders(i.DB)
	if err != nil {
		return err
	}
	rsp.Ids = ids
	return nil
}

func (i *Inventory) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	products, err := model.ListProducts(i.DB, req.Ids)
	if err != nil {
		return err
	}
	rsp.Products = make([]*pb.Product, len(products))
	for i, p := range products {
		rsp.Products[i] = &pb.Product{
			Id:             p.ID,
			UnitsInStock:   p.UnitsInStock,
			UnitsReserved:  p.UnitsReserved,
			UnitsAvailable: p.UnitsAvailable,
		}
	}
	return nil
}

func (i *Inventory) Topup(ctx context.Context, req *pb.TopupRequest, rsp *pb.TopupResponse) error {
	// validate the request
	if len(req.Topups) == 0 {
		return ErrMissingTopups
	}
	for _, t := range req.Topups {
		if err := validateTopup(t); err != nil {
			return err
		}
	}

	return model.CreateTopups(i.DB, req.Topups)
}

func (i *Inventory) Override(ctx context.Context, req *pb.OverrideRequest, rsp *pb.OverrideResponse) error {
	// validate the request
	if len(req.Id) == 0 {
		return ErrMissingID
	} else if len(req.ProductId) == 0 {
		return ErrMissingProductID
	}

	return model.OverrideInventory(i.DB, req.Id, req.ProductId, req.Quantity)
}

func validateItem(i *pb.Item) error {
	if len(i.Id) == 0 {
		return ErrMissingItemID
	} else if len(i.ProductId) == 0 {
		return ErrMissingItemProductID
	} else if i.Quantity <= 0 {
		return ErrMissingItemQuantity
	}
	return nil
}

func validateTopup(p *pb.Topup) error {
	if len(p.Id) == 0 {
		return ErrMissingTopupID
	} else if len(p.ProductId) == 0 {
		return ErrMissingTopupProductID
	} else if p.Quantity <= 0 {
		return ErrMissingTopupQuantity
	}
	return nil
}
