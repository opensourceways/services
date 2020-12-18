package handler

import (
	"context"

	"tasks/model"
	pb "tasks/proto"

	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Tasks struct {
	DB *gorm.DB
}

func (t *Tasks) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
	// validate the request
	if len(req.Type) == 0 {
		return errors.BadRequest("tasks.MissingType", "Missing type")
	} else if len(req.SubjectIds) == 0 {
		return errors.BadRequest("tasks.MissingSubjectIDs", "Missing subject ids")
	}

	// construct the task
	task := &model.Task{
		ID:         uuid.New().String(),
		Type:       req.Type,
		SubjectIDs: req.SubjectIds,
	}
	if len(req.Tag) > 0 {
		task.Tag = &req.Tag
	}
	if len(req.GroupingId) > 0 {
		task.GroupingID = &req.GroupingId
	}
	if req.OrderingTime != nil {
		ot := req.OrderingTime.AsTime()
		task.OrderingTime = &ot
	}
	if req.DeferredUntil != nil {
		du := req.DeferredUntil.AsTime()
		task.DeferredUntil = &du
	}

	// create the task
	task, err := model.CreateTask(t.DB, task)
	if err != nil {
		return err
	}

	// serialize the response
	rsp.Task = serializeTask(task)
	return nil
}

func (t *Tasks) Cancel(ctx context.Context, req *pb.CancelRequest, rsp *pb.CancelResponse) error {
	// validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("tasks.MissingID", "Missing ID")
	}

	return model.CancelTask(t.DB, req.Id)
}

func (t *Tasks) Complete(ctx context.Context, req *pb.CompleteRequest, rsp *pb.CompleteResponse) error {
	// validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("tasks.MissingID", "Missing ID")
	}

	// complete the task, passing user id if provided
	var uid *string
	if len(req.UserId) > 0 {
		uid = &req.UserId
	}
	return model.CompleteTask(t.DB, req.Id, uid)

}

func (t *Tasks) Defer(ctx context.Context, req *pb.DeferRequest, rsp *pb.DeferResponse) error {
	// validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("tasks.MissingID", "Missing ID")
	} else if req.DeferredUntil == nil {
		return errors.BadRequest("tasks.MissingDeferredUntil", "Missing DeferredUntil")
	}

	return model.DeferTask(t.DB, req.Id, req.DeferredUntil.AsTime())
}

func (t *Tasks) Next(ctx context.Context, req *pb.NextRequest, rsp *pb.NextResponse) error {
	// validate the request
	if len(req.Type) == 0 {
		return errors.BadRequest("tasks.MissingType", "Missing Type")
	} else if len(req.UserId) == 0 {
		return errors.BadRequest("tasks.MissingUserID", "Missing UserID")
	}

	// get the next task
	task, err := model.NextTask(t.DB, req.UserId, req.Type, req.Tags)
	if err != nil {
		return err
	}

	// serialize the response
	rsp.Task = serializeTask(task)
	return nil
}

func (t *Tasks) Unassign(ctx context.Context, req *pb.UnassignRequest, rsp *pb.UnassignResponse) error {
	// validate the request
	if len(req.UserId) == 0 {
		return errors.BadRequest("tasks.MissingUserID", "Missing UserID")
	}

	return model.UnassignUser(t.DB, req.UserId)
}

func (t *Tasks) Remove(ctx context.Context, req *pb.RemoveRequest, rsp *pb.RemoveResponse) error {
	// validate the request
	if len(req.SubjectId) == 0 {
		return errors.BadRequest("tasks.MissingSubjectID", "Missing SubjectID")
	}

	return model.RemoveSubject(t.DB, req.SubjectId)
}

func serializeTask(t *model.Task) *pb.Task {
	result := &pb.Task{
		Id:         t.ID,
		Type:       t.Type,
		SubjectIds: t.SubjectIDs,
		CreatedAt:  timestamppb.New(t.CreatedAt),
	}
	if t.Tag != nil {
		result.Tag = *t.Tag
	}
	if t.GroupingID != nil {
		result.GroupingId = *t.GroupingID
	}
	if t.AllocatedTo != nil {
		result.AllocatedTo = *t.AllocatedTo
	}
	if t.OrderingTime != nil {
		result.OrderingTime = timestamppb.New(*t.OrderingTime)
	}
	if t.DeferredUntil != nil {
		result.DeferredUntil = timestamppb.New(*t.DeferredUntil)
	}
	if t.CompletedAt != nil {
		result.CompletedAt = timestamppb.New(*t.CompletedAt)
	}
	if t.CancelledAt != nil {
		result.CancelledAt = timestamppb.New(*t.CancelledAt)
	}

	return result
}
