package handler_test

import (
	"context"
	"inventory/handler"
	"inventory/model"
	pb "inventory/proto"
	"strings"
	"testing"

	"github.com/bradfitz/slice"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	orderIDOne   = "order-1"
	orderIDTwo   = "order-2"
	orderIDThree = "order-3"
	productIDOne = "product-1"
	productIDTwo = "product-2"
)

func testHandler(t *testing.T) *handler.Inventory {
	// connect to the database
	db, err := gorm.Open(postgres.Open("postgresql://postgres@localhost:5432/inventory?sslmode=disable"), nil)
	if err != nil {
		t.Fatal(err)
	}

	// migrate the database
	if err := db.AutoMigrate(&model.Order{}, &model.Item{}, &model.Topup{}); err != nil {
		t.Fatal(err)
	}

	// drop all the data from the tables
	var tables []string
	if err := db.Table("pg_tables").
		Where("schemaname = 'public' and tablename != 'schema_migrations'").
		Pluck("tablename", &tables).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("TRUNCATE TABLE " + strings.Join(tables, ",") + " CASCADE").Error; err != nil {
		t.Fatal(err)
	}

	return &handler.Inventory{DB: db}
}

func TestReserve(t *testing.T) {
	h := testHandler(t)

	t.Run("MissingOrderID", func(t *testing.T) {
		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.ReserveResponse{})
		assert.Equal(t, handler.ErrMissingOrderID, err)
	})

	t.Run("MissingItems", func(t *testing.T) {
		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: uuid.New().String(),
		}, &pb.ReserveResponse{})
		assert.Equal(t, handler.ErrMissingItems, err)
	})

	t.Run("InvalidItems", func(t *testing.T) {
		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: uuid.New().String(),
			Items: []*pb.Item{
				&pb.Item{
					Id:       uuid.New().String(),
					Quantity: 1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.Equal(t, handler.ErrMissingItemProductID, err)
	})

	t.Run("Valid", func(t *testing.T) {
		run := func() {
			err := h.Reserve(context.TODO(), &pb.ReserveRequest{
				OrderId: orderIDOne,
				Items: []*pb.Item{
					&pb.Item{
						Id:        uuid.New().String(),
						ProductId: productIDOne,
						Quantity:  1,
					},
					&pb.Item{
						Id:        uuid.New().String(),
						ProductId: productIDTwo,
						Quantity:  2,
					},
				},
			}, &pb.ReserveResponse{})
			assert.NoError(t, err)
		}

		// run the Reserve function twice to ensure it is retry safe
		run()
		run()
	})
}

func TestPick(t *testing.T) {
	h := testHandler(t)

	t.Run("MissingOrderID", func(t *testing.T) {
		err := h.Pick(context.TODO(), &pb.PickRequest{
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.PickResponse{})
		assert.Equal(t, handler.ErrMissingOrderID, err)
	})

	t.Run("MissingItems", func(t *testing.T) {
		err := h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: uuid.New().String(),
		}, &pb.PickResponse{})
		assert.Equal(t, handler.ErrMissingItems, err)
	})

	t.Run("InvalidItems", func(t *testing.T) {
		err := h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: uuid.New().String(),
			Items: []*pb.Item{
				&pb.Item{
					Id:       uuid.New().String(),
					Quantity: 1,
				},
			},
		}, &pb.PickResponse{})
		assert.Equal(t, handler.ErrMissingItemProductID, err)
	})

	t.Run("Reserved", func(t *testing.T) {
		// firstly, reserve the order
		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		// then Pick the inventory for the order
		err = h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: orderIDOne,
		}, &pb.PickResponse{})
		assert.NoError(t, err)
	})

	t.Run("Unreserved", func(t *testing.T) {
		err := h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.PickResponse{})
		assert.NoError(t, err)
	})
}

func TestVoid(t *testing.T) {
	h := testHandler(t)

	t.Run("MissingOrderID", func(t *testing.T) {
		err := h.Void(context.TODO(), &pb.VoidRequest{}, &pb.VoidResponse{})
		assert.Equal(t, handler.ErrMissingOrderID, err)
	})

	t.Run("Reserved", func(t *testing.T) {
		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		// then void the order
		err = h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: orderIDOne,
		}, &pb.PickResponse{})
		assert.NoError(t, err)
	})

	t.Run("Pickd", func(t *testing.T) {
		err := h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: orderIDTwo,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.PickResponse{})
		assert.NoError(t, err)

		// then void the order
		err = h.Void(context.TODO(), &pb.VoidRequest{
			OrderId: orderIDOne,
		}, &pb.VoidResponse{})
		assert.NoError(t, err)
	})

	t.Run("NotFound", func(t *testing.T) {
		err := h.Void(context.TODO(), &pb.VoidRequest{
			OrderId: uuid.New().String(),
		}, &pb.VoidResponse{})
		assert.NoError(t, err)
	})
}

func TestTopup(t *testing.T) {
	h := testHandler(t)

	t.Run("MissingTopups", func(t *testing.T) {
		err := h.Topup(context.TODO(), &pb.TopupRequest{}, &pb.TopupResponse{})
		assert.Equal(t, handler.ErrMissingTopups, err)
	})

	t.Run("InvalidTopups", func(t *testing.T) {
		err := h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:       uuid.New().String(),
					Quantity: 1,
				},
			},
		}, &pb.TopupResponse{})
		assert.Equal(t, handler.ErrMissingTopupProductID, err)
	})

	t.Run("Valid", func(t *testing.T) {
		run := func() {
			err := h.Topup(context.TODO(), &pb.TopupRequest{
				Topups: []*pb.Topup{
					&pb.Topup{
						Id:        uuid.New().String(),
						ProductId: productIDOne,
						Quantity:  10,
					},
					&pb.Topup{
						Id:        uuid.New().String(),
						ProductId: productIDTwo,
						Quantity:  20,
					},
				},
			}, &pb.TopupResponse{})
			assert.NoError(t, err)
		}

		// run the Topup function twice to ensure it is retry safe
		run()
		run()
	})
}

func TestOverride(t *testing.T) {
	h := testHandler(t)

	t.Run("MissingID", func(t *testing.T) {
		err := h.Override(context.TODO(), &pb.OverrideRequest{
			ProductId: productIDOne,
		}, &pb.OverrideResponse{})
		assert.Equal(t, handler.ErrMissingID, err)
	})

	t.Run("MissingProductID", func(t *testing.T) {
		err := h.Override(context.TODO(), &pb.OverrideRequest{
			Id: uuid.New().String(),
		}, &pb.OverrideResponse{})
		assert.Equal(t, handler.ErrMissingProductID, err)
	})

	t.Run("Valid", func(t *testing.T) {
		id := uuid.New().String()

		run := func() {
			err := h.Override(context.TODO(), &pb.OverrideRequest{
				Id: id, ProductId: productIDOne, Quantity: 10,
			}, &pb.OverrideResponse{})
			assert.NoError(t, err)
		}

		// run the Override function twice to ensure it is retry safe
		run()
		run()
	})
}

func TestList(t *testing.T) {
	t.Run("NoProducts", func(t *testing.T) {
		h := testHandler(t)
		var rsp pb.ListResponse
		err := h.List(context.TODO(), &pb.ListRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Products)
	})

	t.Run("TopupOnly", func(t *testing.T) {
		h := testHandler(t)

		err := h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		var rsp pb.ListResponse
		err = h.List(context.TODO(), &pb.ListRequest{}, &rsp)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp.Products, "Expected products")
		p := rsp.Products[0]
		assert.NotNilf(t, p, "Expected one product")
		assert.Equal(t, productIDOne, p.Id)
		assert.Equal(t, int32(100), p.UnitsInStock)
		assert.Equal(t, int32(100), p.UnitsAvailable)
		assert.Equal(t, int32(0), p.UnitsReserved)
	})

	t.Run("TopupAndReserve", func(t *testing.T) {
		h := testHandler(t)

		err := h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  150,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		err = h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		var rsp pb.ListResponse
		err = h.List(context.TODO(), &pb.ListRequest{}, &rsp)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp.Products, "Expected products")
		slice.Sort(rsp.Products, func(i, j int) bool {
			return rsp.Products[i].Id < rsp.Products[j].Id
		})

		p1 := rsp.Products[0]
		assert.NotNilf(t, p1, "Expected two products")
		assert.Equal(t, productIDOne, p1.Id)
		assert.Equal(t, int32(100), p1.UnitsInStock)
		assert.Equal(t, int32(99), p1.UnitsAvailable)
		assert.Equal(t, int32(1), p1.UnitsReserved)

		p2 := rsp.Products[1]
		assert.NotNilf(t, p2, "Expected two products")
		assert.Equal(t, productIDTwo, p2.Id)
		assert.Equal(t, int32(150), p2.UnitsInStock)
		assert.Equal(t, int32(148), p2.UnitsAvailable)
		assert.Equal(t, int32(2), p2.UnitsReserved)
	})

	t.Run("TopupReserveAndPick", func(t *testing.T) {
		h := testHandler(t)

		err := h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  150,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		err = h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		err = h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: orderIDTwo,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  2,
				},
			},
		}, &pb.PickResponse{})
		assert.NoError(t, err)

		var rsp pb.ListResponse
		err = h.List(context.TODO(), &pb.ListRequest{}, &rsp)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp.Products, "Expected products")
		slice.Sort(rsp.Products, func(i, j int) bool {
			return rsp.Products[i].Id < rsp.Products[j].Id
		})

		p1 := rsp.Products[0]
		assert.NotNilf(t, p1, "Expected two products")
		assert.Equal(t, productIDOne, p1.Id)
		assert.Equal(t, int32(100), p1.UnitsInStock)
		assert.Equal(t, int32(99), p1.UnitsAvailable)
		assert.Equal(t, int32(1), p1.UnitsReserved)

		p2 := rsp.Products[1]
		assert.NotNilf(t, p2, "Expected two products")
		assert.Equal(t, productIDTwo, p2.Id)
		assert.Equal(t, int32(148), p2.UnitsInStock)
		assert.Equal(t, int32(147), p2.UnitsAvailable)
		assert.Equal(t, int32(1), p2.UnitsReserved)
	})
}

func TestUnfulfilableOrders(t *testing.T) {
	t.Run("NoOrders", func(t *testing.T) {
		h := testHandler(t)
		var rsp pb.UnfulfilableOrdersResponse
		err := h.UnfulfilableOrders(context.TODO(), &pb.UnfulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Ids)
	})

	t.Run("UnfulfilableOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		var rsp pb.UnfulfilableOrdersResponse
		err = h.UnfulfilableOrders(context.TODO(), &pb.UnfulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Len(t, rsp.Ids, 1)
		assert.Contains(t, rsp.Ids, orderIDOne)
	})

	t.Run("PartiallyUnfulfilableOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		err = h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		var rsp pb.UnfulfilableOrdersResponse
		err = h.UnfulfilableOrders(context.TODO(), &pb.UnfulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Len(t, rsp.Ids, 1)
		assert.Contains(t, rsp.Ids, orderIDOne)
	})

	t.Run("FulfilableOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		err = h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		var rsp pb.UnfulfilableOrdersResponse
		err = h.UnfulfilableOrders(context.TODO(), &pb.UnfulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Ids)
	})

	t.Run("FulfilledOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		err = h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.PickResponse{})
		assert.NoError(t, err)

		var rsp pb.UnfulfilableOrdersResponse
		err = h.UnfulfilableOrders(context.TODO(), &pb.UnfulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Ids)
	})
}
func TestFulfilableOrders(t *testing.T) {
	t.Run("NoOrders", func(t *testing.T) {
		h := testHandler(t)
		var rsp pb.FulfilableOrdersResponse
		err := h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Ids)
	})

	t.Run("UnfulfilableOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		var rsp pb.FulfilableOrdersResponse
		err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Ids)
	})

	t.Run("PartiallyUnfulfilableOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		err = h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		var rsp pb.FulfilableOrdersResponse
		err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Ids)
	})

	t.Run("FulfilableOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		err = h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		var rsp pb.FulfilableOrdersResponse
		err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &rsp)
		assert.Len(t, rsp.Ids, 1)
		assert.Contains(t, rsp.Ids, orderIDOne)
	})

	t.Run("FulfilableOrderAfterTopup", func(t *testing.T) {
		h := testHandler(t)

		err := h.Reserve(context.TODO(), &pb.ReserveRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.ReserveResponse{})
		assert.NoError(t, err)

		err = h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		var rsp pb.FulfilableOrdersResponse
		err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &rsp)
		assert.Len(t, rsp.Ids, 1)
		assert.Contains(t, rsp.Ids, orderIDOne)
	})

	t.Run("FulfilledOrder", func(t *testing.T) {
		h := testHandler(t)

		err := h.Topup(context.TODO(), &pb.TopupRequest{
			Topups: []*pb.Topup{
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  100,
				},
				&pb.Topup{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  100,
				},
			},
		}, &pb.TopupResponse{})
		assert.NoError(t, err)

		err = h.Pick(context.TODO(), &pb.PickRequest{
			OrderId: orderIDOne,
			Items: []*pb.Item{
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDOne,
					Quantity:  1,
				},
				&pb.Item{
					Id:        uuid.New().String(),
					ProductId: productIDTwo,
					Quantity:  1,
				},
			},
		}, &pb.PickResponse{})
		assert.NoError(t, err)

		var rsp pb.FulfilableOrdersResponse
		err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &rsp)
		assert.NoError(t, err)
		assert.Empty(t, rsp.Ids)
	})
}

func TestEndToEnd(t *testing.T) {
	h := testHandler(t)

	// inventory is topped up with 10 units for each of the two products
	err := h.Topup(context.TODO(), &pb.TopupRequest{
		Topups: []*pb.Topup{
			&pb.Topup{
				Id:        uuid.New().String(),
				ProductId: productIDOne,
				Quantity:  10,
			},
			&pb.Topup{
				Id:        uuid.New().String(),
				ProductId: productIDTwo,
				Quantity:  10,
			},
		},
	}, &pb.TopupResponse{})
	assert.NoError(t, err)

	// four orders are received
	err = h.Reserve(context.TODO(), &pb.ReserveRequest{
		OrderId: "one",
		Items: []*pb.Item{
			&pb.Item{
				Id:        uuid.New().String(),
				ProductId: productIDOne,
				Quantity:  3,
			},
		},
	}, &pb.ReserveResponse{})
	assert.NoError(t, err)

	err = h.Reserve(context.TODO(), &pb.ReserveRequest{
		OrderId: "two",
		Items: []*pb.Item{
			&pb.Item{
				Id:        uuid.New().String(),
				ProductId: productIDOne,
				Quantity:  6,
			},
			&pb.Item{
				Id:        uuid.New().String(),
				ProductId: productIDTwo,
				Quantity:  2,
			},
		},
	}, &pb.ReserveResponse{})
	assert.NoError(t, err)

	err = h.Reserve(context.TODO(), &pb.ReserveRequest{
		OrderId: "three",
		Items: []*pb.Item{
			&pb.Item{
				Id:        uuid.New().String(),
				ProductId: productIDOne,
				Quantity:  2,
			},
		},
	}, &pb.ReserveResponse{})
	assert.NoError(t, err)

	err = h.Reserve(context.TODO(), &pb.ReserveRequest{
		OrderId: "four",
		Items: []*pb.Item{
			&pb.Item{
				Id:        uuid.New().String(),
				ProductId: productIDTwo,
				Quantity:  2,
			},
		},
	}, &pb.ReserveResponse{})
	assert.NoError(t, err)

	// the inventory should update correctly
	var rsp pb.ListResponse
	err = h.List(context.TODO(), &pb.ListRequest{}, &rsp)
	assert.NoError(t, err)
	assert.NotNilf(t, rsp.Products, "Expected products")
	slice.Sort(rsp.Products, func(i, j int) bool {
		return rsp.Products[i].Id < rsp.Products[j].Id
	})

	p1 := rsp.Products[0]
	assert.NotNilf(t, p1, "Expected two products")
	assert.Equal(t, productIDOne, p1.Id)
	assert.Equal(t, int32(10), p1.UnitsInStock)
	assert.Equal(t, int32(-1), p1.UnitsAvailable)
	assert.Equal(t, int32(11), p1.UnitsReserved)

	p2 := rsp.Products[1]
	assert.NotNilf(t, p2, "Expected two products")
	assert.Equal(t, productIDTwo, p2.Id)
	assert.Equal(t, int32(10), p2.UnitsInStock)
	assert.Equal(t, int32(6), p2.UnitsAvailable)
	assert.Equal(t, int32(4), p2.UnitsReserved)

	// only orders where there is enough inventory should be returned as fulfillable, the rest should
	// be returned as unfulfilable
	var fRsp pb.FulfilableOrdersResponse
	err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &fRsp)
	assert.Len(t, fRsp.Ids, 3)
	assert.Contains(t, fRsp.Ids, "one")
	assert.Contains(t, fRsp.Ids, "two")
	assert.Contains(t, fRsp.Ids, "four")

	var ufRsp pb.UnfulfilableOrdersResponse
	err = h.UnfulfilableOrders(context.TODO(), &pb.UnfulfilableOrdersRequest{}, &ufRsp)
	assert.Len(t, ufRsp.Ids, 1)
	assert.Contains(t, ufRsp.Ids, "three")

	// picking some orders updates the inventory and removes the orders from the fulfilment endpoint
	err = h.Pick(context.TODO(), &pb.PickRequest{OrderId: "one"}, &pb.PickResponse{})
	assert.NoError(t, err)
	err = h.Pick(context.TODO(), &pb.PickRequest{OrderId: "two"}, &pb.PickResponse{})
	assert.NoError(t, err)

	err = h.List(context.TODO(), &pb.ListRequest{}, &rsp)
	assert.NoError(t, err)
	assert.NotNilf(t, rsp.Products, "Expected products")
	slice.Sort(rsp.Products, func(i, j int) bool {
		return rsp.Products[i].Id < rsp.Products[j].Id
	})

	p1 = rsp.Products[0]
	assert.NotNilf(t, p1, "Expected two products")
	assert.Equal(t, productIDOne, p1.Id)
	assert.Equal(t, int32(1), p1.UnitsInStock)
	assert.Equal(t, int32(-1), p1.UnitsAvailable)
	assert.Equal(t, int32(2), p1.UnitsReserved)

	p2 = rsp.Products[1]
	assert.NotNilf(t, p2, "Expected two products")
	assert.Equal(t, productIDTwo, p2.Id)
	assert.Equal(t, int32(8), p2.UnitsInStock)
	assert.Equal(t, int32(6), p2.UnitsAvailable)
	assert.Equal(t, int32(2), p2.UnitsReserved)

	err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &fRsp)
	assert.Len(t, fRsp.Ids, 1)
	assert.Contains(t, fRsp.Ids, "four")

	// now overriding the inventory for product one should move it to fulfilable
	err = h.Override(context.TODO(), &pb.OverrideRequest{
		Id: uuid.New().String(), ProductId: productIDOne, Quantity: 10,
	}, &pb.OverrideResponse{})
	assert.NoError(t, err)

	err = h.FulfilableOrders(context.TODO(), &pb.FulfilableOrdersRequest{}, &fRsp)
	assert.Len(t, fRsp.Ids, 2)
	assert.Contains(t, fRsp.Ids, "three")
	assert.Contains(t, fRsp.Ids, "four")

	err = h.UnfulfilableOrders(context.TODO(), &pb.UnfulfilableOrdersRequest{}, &ufRsp)
	assert.Len(t, ufRsp.Ids, 0)

	err = h.List(context.TODO(), &pb.ListRequest{}, &rsp)
	assert.NoError(t, err)
	assert.NotNilf(t, rsp.Products, "Expected products")
	slice.Sort(rsp.Products, func(i, j int) bool {
		return rsp.Products[i].Id < rsp.Products[j].Id
	})

	p1 = rsp.Products[0]
	assert.NotNilf(t, p1, "Expected two products")
	assert.Equal(t, productIDOne, p1.Id)
	assert.Equal(t, int32(10), p1.UnitsInStock)
	assert.Equal(t, int32(8), p1.UnitsAvailable)
	assert.Equal(t, int32(2), p1.UnitsReserved)

	p2 = rsp.Products[1]
	assert.NotNilf(t, p2, "Expected two products")
	assert.Equal(t, productIDTwo, p2.Id)
	assert.Equal(t, int32(8), p2.UnitsInStock)
	assert.Equal(t, int32(6), p2.UnitsAvailable)
	assert.Equal(t, int32(2), p2.UnitsReserved)

	// picking the final two orders should set the units reserved to zero
	err = h.Pick(context.TODO(), &pb.PickRequest{OrderId: "three"}, &pb.PickResponse{})
	assert.NoError(t, err)
	err = h.Pick(context.TODO(), &pb.PickRequest{OrderId: "four"}, &pb.PickResponse{})
	assert.NoError(t, err)

	err = h.List(context.TODO(), &pb.ListRequest{}, &rsp)
	assert.NoError(t, err)
	assert.NotNilf(t, rsp.Products, "Expected products")
	slice.Sort(rsp.Products, func(i, j int) bool {
		return rsp.Products[i].Id < rsp.Products[j].Id
	})

	p1 = rsp.Products[0]
	assert.NotNilf(t, p1, "Expected two products")
	assert.Equal(t, productIDOne, p1.Id)
	assert.Equal(t, int32(8), p1.UnitsInStock)
	assert.Equal(t, int32(8), p1.UnitsAvailable)
	assert.Equal(t, int32(0), p1.UnitsReserved)

	p2 = rsp.Products[1]
	assert.NotNilf(t, p2, "Expected two products")
	assert.Equal(t, productIDTwo, p2.Id)
	assert.Equal(t, int32(6), p2.UnitsInStock)
	assert.Equal(t, int32(6), p2.UnitsAvailable)
	assert.Equal(t, int32(0), p2.UnitsReserved)
}
