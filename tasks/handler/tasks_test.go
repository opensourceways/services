package handler_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"tasks/handler"
	"tasks/model"
	pb "tasks/proto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func testHandler(t *testing.T) *handler.Tasks {
	// connect to the database
	db, err := gorm.Open(postgres.Open("postgresql://postgres@localhost:5433/tasks?sslmode=disable"), nil)
	if err != nil {
		t.Fatal(err)
	}

	// migrate the database
	if err := db.AutoMigrate(&model.Task{}); err != nil {
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

	return &handler.Tasks{DB: db.Debug()}
}

func TestCreate(t *testing.T) {
	t.Run("MissingType", func(t *testing.T) {
		h := testHandler(t)
		err := h.Create(context.TODO(), &pb.CreateRequest{
			SubjectIds: []string{"order-1"},
		}, &pb.CreateResponse{})
		assert.Error(t, err)
	})

	t.Run("MissingSubjectIDs", func(t *testing.T) {
		h := testHandler(t)
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick",
		}, &pb.CreateResponse{})
		assert.Error(t, err)
	})

	t.Run("NoGroupingID", func(t *testing.T) {
		h := testHandler(t)
		var rsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-1"},
			Tag:        "frozen",
		}, &rsp)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp.Task, "A task should be returned")
		assert.NotEmpty(t, rsp.Task.Id)
		assert.Contains(t, rsp.Task.SubjectIds, "order-1")
		assert.Equal(t, "pick", rsp.Task.Type)
		assert.Equal(t, "frozen", rsp.Task.Tag)
		assert.Empty(t, rsp.Task.AllocatedTo)
		assert.NotNil(t, rsp.Task.CreatedAt)
		assert.Nil(t, rsp.Task.CancelledAt)
		assert.Nil(t, rsp.Task.CompletedAt)
	})

	t.Run("WithNoMatchingGroupID", func(t *testing.T) {
		h := testHandler(t)
		var rsp1 pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-1"},
			GroupingId: "foo",
		}, &rsp1)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp1.Task, "A task should be returned")
		assert.Contains(t, rsp1.Task.SubjectIds, "order-1")
		assert.Len(t, rsp1.Task.SubjectIds, 1)

		var rsp2 pb.CreateResponse
		err = h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-2"},
			GroupingId: "bar",
		}, &rsp2)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp2.Task, "A task should be returned")
		assert.Contains(t, rsp2.Task.SubjectIds, "order-2")
		assert.Len(t, rsp2.Task.SubjectIds, 1)
	})

	t.Run("WithGroupIDButNoMatchingTag", func(t *testing.T) {
		h := testHandler(t)
		var rsp1 pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-1"},
			GroupingId: "foo",
		}, &rsp1)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp1.Task, "A task should be returned")
		assert.Contains(t, rsp1.Task.SubjectIds, "order-1")
		assert.Len(t, rsp1.Task.SubjectIds, 1)

		var rsp2 pb.CreateResponse
		err = h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-2"},
			GroupingId: "bar",
			Tag:        "frozen",
		}, &rsp2)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp2.Task, "A task should be returned")
		assert.Contains(t, rsp2.Task.SubjectIds, "order-2")
		assert.Len(t, rsp2.Task.SubjectIds, 1)
		assert.NotEqual(t, rsp2.Task.Id, rsp1.Task.Id)
	})

	t.Run("WithMatches", func(t *testing.T) {
		h := testHandler(t)

		var rsp1 pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-1"},
			GroupingId: "foo",
		}, &rsp1)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp1.Task, "A task should be returned")
		assert.Contains(t, rsp1.Task.SubjectIds, "order-1")
		assert.Len(t, rsp1.Task.SubjectIds, 1)

		var rsp2 pb.CreateResponse
		err = h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-2"},
			GroupingId: "foo",
		}, &rsp2)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp2.Task, "A task should be returned")
		assert.Contains(t, rsp2.Task.SubjectIds, "order-1")
		assert.Contains(t, rsp2.Task.SubjectIds, "order-2")
		assert.Len(t, rsp2.Task.SubjectIds, 2)
		assert.Equal(t, rsp1.Task.Id, rsp2.Task.Id)
	})
}

func TestCancel(t *testing.T) {
	t.Run("MissingID", func(t *testing.T) {
		h := testHandler(t)
		err := h.Cancel(context.TODO(), &pb.CancelRequest{}, &pb.CancelResponse{})
		assert.Error(t, err)
	})

	t.Run("NotFound", func(t *testing.T) {
		h := testHandler(t)
		err := h.Cancel(context.TODO(), &pb.CancelRequest{
			Id: uuid.New().String(),
		}, &pb.CancelResponse{})
		assert.Error(t, err)
	})

	t.Run("Completed", func(t *testing.T) {
		h := testHandler(t)

		var cRsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick", SubjectIds: []string{"order-1"},
		}, &cRsp)
		assert.NoError(t, err)

		err = h.Complete(context.TODO(), &pb.CompleteRequest{
			Id: cRsp.Task.Id,
		}, &pb.CompleteResponse{})
		assert.NoError(t, err)

		err = h.Cancel(context.TODO(), &pb.CancelRequest{
			Id: cRsp.Task.Id,
		}, &pb.CancelResponse{})
		assert.Error(t, err)
	})

	t.Run("Pending", func(t *testing.T) {
		h := testHandler(t)

		var cRsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick", SubjectIds: []string{"order-1"},
		}, &cRsp)
		assert.NoError(t, err)

		err = h.Cancel(context.TODO(), &pb.CancelRequest{
			Id: cRsp.Task.Id,
		}, &pb.CancelResponse{})
		assert.NoError(t, err)

		// should be retry safe
		err = h.Cancel(context.TODO(), &pb.CancelRequest{
			Id: cRsp.Task.Id,
		}, &pb.CancelResponse{})
		assert.NoError(t, err)
	})
}

func TestComplete(t *testing.T) {
	t.Run("MissingID", func(t *testing.T) {
		h := testHandler(t)
		err := h.Complete(context.TODO(), &pb.CompleteRequest{}, &pb.CompleteResponse{})
		assert.Error(t, err)
	})

	t.Run("NotFound", func(t *testing.T) {
		h := testHandler(t)
		err := h.Complete(context.TODO(), &pb.CompleteRequest{
			Id: uuid.New().String(),
		}, &pb.CompleteResponse{})
		assert.Error(t, err)
	})

	t.Run("Cancelled", func(t *testing.T) {
		h := testHandler(t)

		var cRsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick", SubjectIds: []string{"order-1"},
		}, &cRsp)
		assert.NoError(t, err)

		err = h.Cancel(context.TODO(), &pb.CancelRequest{
			Id: cRsp.Task.Id,
		}, &pb.CancelResponse{})
		assert.NoError(t, err)

		err = h.Complete(context.TODO(), &pb.CompleteRequest{
			Id: cRsp.Task.Id,
		}, &pb.CompleteResponse{})
		assert.Error(t, err)
	})

	t.Run("Pending", func(t *testing.T) {
		h := testHandler(t)

		var cRsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick", SubjectIds: []string{"order-1"},
		}, &cRsp)
		assert.NoError(t, err)

		err = h.Complete(context.TODO(), &pb.CompleteRequest{
			Id: cRsp.Task.Id,
		}, &pb.CompleteResponse{})
		assert.NoError(t, err)

		// should be retry safe
		err = h.Complete(context.TODO(), &pb.CompleteRequest{
			Id: cRsp.Task.Id,
		}, &pb.CompleteResponse{})
		assert.NoError(t, err)
	})
}

func TestDefer(t *testing.T) {
	t.Run("MissingID", func(t *testing.T) {
		h := testHandler(t)
		err := h.Defer(context.TODO(), &pb.DeferRequest{
			DeferredUntil: timestamppb.New(time.Now().Add(time.Hour)),
		}, &pb.DeferResponse{})
		assert.Error(t, err)
	})

	t.Run("MissingDeferredUntil", func(t *testing.T) {
		h := testHandler(t)
		err := h.Defer(context.TODO(), &pb.DeferRequest{
			Id: uuid.New().String(),
		}, &pb.DeferResponse{})
		assert.Error(t, err)
	})

	t.Run("NotFound", func(t *testing.T) {
		h := testHandler(t)
		err := h.Defer(context.TODO(), &pb.DeferRequest{
			Id:            uuid.New().String(),
			DeferredUntil: timestamppb.New(time.Now().Add(time.Hour)),
		}, &pb.DeferResponse{})
		assert.Error(t, err)
	})

	t.Run("Completed", func(t *testing.T) {
		h := testHandler(t)

		var cRsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick", SubjectIds: []string{"order-1"},
		}, &cRsp)
		assert.NoError(t, err)

		err = h.Complete(context.TODO(), &pb.CompleteRequest{
			Id: cRsp.Task.Id,
		}, &pb.CompleteResponse{})
		assert.NoError(t, err)

		err = h.Defer(context.TODO(), &pb.DeferRequest{
			Id:            cRsp.Task.Id,
			DeferredUntil: timestamppb.New(time.Now().Add(time.Hour)),
		}, &pb.DeferResponse{})
		assert.Error(t, err)
	})

	t.Run("Cancelled", func(t *testing.T) {
		h := testHandler(t)

		var cRsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick", SubjectIds: []string{"order-1"},
		}, &cRsp)
		assert.NoError(t, err)

		err = h.Cancel(context.TODO(), &pb.CancelRequest{
			Id: cRsp.Task.Id,
		}, &pb.CancelResponse{})
		assert.NoError(t, err)

		err = h.Defer(context.TODO(), &pb.DeferRequest{
			Id:            cRsp.Task.Id,
			DeferredUntil: timestamppb.New(time.Now().Add(time.Hour)),
		}, &pb.DeferResponse{})
		assert.Error(t, err)
	})

	t.Run("Valid", func(t *testing.T) {
		h := testHandler(t)

		var cRsp pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pick", SubjectIds: []string{"order-1"},
		}, &cRsp)
		assert.NoError(t, err)

		// the task should now appear in next
		var nRsp pb.NextResponse
		err = h.Next(context.TODO(), &pb.NextRequest{
			Type: "pick", UserId: uuid.New().String(),
		}, &nRsp)
		assert.NoError(t, err)
		assert.NotNilf(t, nRsp.Task, "Should return the task")
		assert.Equal(t, cRsp.Task.Id, nRsp.Task.Id)

		// defer the task
		err = h.Defer(context.TODO(), &pb.DeferRequest{
			Id:            cRsp.Task.Id,
			DeferredUntil: timestamppb.New(time.Now().Add(time.Hour)),
		}, &pb.DeferResponse{})
		assert.NoError(t, err)

		// the task sould no longer appear in the task queue
		err = h.Next(context.TODO(), &pb.NextRequest{
			Type: "pick", UserId: uuid.New().String(),
		}, &nRsp)
		assert.Error(t, err)

		// after the deferred until time has passed, the task should reappear
		model.Time = func() time.Time {
			return time.Now().Add(time.Hour)
		}
		nRsp = pb.NextResponse{}
		err = h.Next(context.TODO(), &pb.NextRequest{
			Type: "pick", UserId: uuid.New().String(),
		}, &nRsp)
		assert.NoError(t, err)
		assert.NotNilf(t, nRsp.Task, "Should return the task")
		assert.Equal(t, cRsp.Task.Id, nRsp.Task.Id)
		model.Time = time.Now
	})
}

func TestNext(t *testing.T) {
	t.Run("MissingUserID", func(t *testing.T) {
		h := testHandler(t)
		err := h.Next(context.TODO(), &pb.NextRequest{Type: "pick"}, &pb.NextResponse{})
		assert.Error(t, err)
	})

	t.Run("MissingType", func(t *testing.T) {
		h := testHandler(t)
		err := h.Next(context.TODO(), &pb.NextRequest{UserId: uuid.New().String()}, &pb.NextResponse{})
		assert.Error(t, err)
	})

	t.Run("NoneLeftWithType", func(t *testing.T) {
		h := testHandler(t)

		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type: "pickx", SubjectIds: []string{"order-1"},
		}, &pb.CreateResponse{})
		assert.NoError(t, err)

		err = h.Next(context.TODO(), &pb.NextRequest{
			Type:   "pick",
			UserId: uuid.New().String(),
		}, &pb.NextResponse{})
		assert.Error(t, err)
	})

	t.Run("NoneLeftWithTag", func(t *testing.T) {
		h := testHandler(t)

		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-1"},
			Tag:        "tag-1",
		}, &pb.CreateResponse{})
		assert.NoError(t, err)

		err = h.Next(context.TODO(), &pb.NextRequest{
			Type:   "pick",
			UserId: uuid.New().String(),
			Tags:   []string{"tag-2"},
		}, &pb.NextResponse{})
		assert.Error(t, err)
	})

	t.Run("MultipleTags", func(t *testing.T) {
		h := testHandler(t)

		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-1"},
			Tag:        "tag-1",
		}, &pb.CreateResponse{})
		assert.NoError(t, err)

		rsp := pb.NextResponse{}
		err = h.Next(context.TODO(), &pb.NextRequest{
			Type:   "pick",
			UserId: uuid.New().String(),
			Tags:   []string{"tag-1", "tag-2"},
		}, &rsp)
		assert.NoError(t, err)
		assert.NotNil(t, rsp.Task)
	})

	t.Run("SingleTag", func(t *testing.T) {
		h := testHandler(t)

		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-1"},
			Tag:        "tag-1",
		}, &pb.CreateResponse{})
		assert.NoError(t, err)

		rsp := pb.NextResponse{}
		err = h.Next(context.TODO(), &pb.NextRequest{
			Type:   "pick",
			UserId: uuid.New().String(),
			Tags:   []string{"tag-1"},
		}, &rsp)
		assert.NoError(t, err)
		assert.NotNil(t, rsp.Task)
	})

	t.Run("Valid", func(t *testing.T) {
		h := testHandler(t)

		// create three test orders
		var cRsp1 pb.CreateResponse
		err := h.Create(context.TODO(), &pb.CreateRequest{
			Type:          "pick",
			SubjectIds:    []string{"order-1"},
			Tag:           "tag-1",
			DeferredUntil: timestamppb.New(time.Now().Add(time.Hour)),
		}, &cRsp1)
		assert.NoError(t, err)

		var cRsp2 pb.CreateResponse
		err = h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-2"},
		}, &cRsp2)
		assert.NoError(t, err)

		var cRsp3 pb.CreateResponse
		err = h.Create(context.TODO(), &pb.CreateRequest{
			Type:       "pick",
			SubjectIds: []string{"order-3"},
		}, &cRsp3)
		assert.NoError(t, err)

		// check the second task is returned
		var rsp pb.NextResponse
		err = h.Next(context.TODO(), &pb.NextRequest{
			Type:   "pick",
			UserId: uuid.New().String(),
		}, &rsp)
		assert.NoError(t, err)
		assert.NotNilf(t, rsp.Task, "Should return a task")
		assert.Equal(t, cRsp2.Task.Id, rsp.Task.Id)
	})
}
