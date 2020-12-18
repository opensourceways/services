package model

import (
	"time"

	"github.com/lib/pq"
	"github.com/micro/go-micro/errors"
	"gorm.io/gorm"
)

var Time func() time.Time = time.Now

type Task struct {
	ID            string
	Type          string
	Tag           *string
	GroupingID    *string
	SubjectIDs    pq.StringArray `gorm:"type:text[]"`
	CreatedAt     time.Time
	OrderingTime  *time.Time
	DeferredUntil *time.Time
	CompletedAt   *time.Time
	CancelledAt   *time.Time
	AllocatedTo   *string
}

func CreateTask(db *gorm.DB, t *Task) (*Task, error) {
	// no need to lookup groupable tasks if no grouping id is provided
	if t.GroupingID == nil {
		return t, db.Create(t).Error
	}

	var result *Task
	return result, db.Transaction(func(tx *gorm.DB) error {
		// check for an existing task with this grouping id which has not yet been completed or allocated
		var existing Task
		err := db.Set("gorm:query_option", "FOR UPDATE").
			Where(&Task{Type: t.Type, GroupingID: t.GroupingID, Tag: t.Tag}).
			Where("completed_at IS NULL AND cancelled_at IS NULL").
			Order("COALESCE(ordering_time, created_at)").
			First(&existing).
			Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		} else if err != gorm.ErrRecordNotFound {
			// the task exists, append the subject id to the existing task and return
			existing.SubjectIDs = uniq(append(existing.SubjectIDs, t.SubjectIDs...))
			result = &existing
			return tx.Save(existing).Error
		}

		// create the task
		result = t
		return tx.Create(t).Error
	})
}

func CompleteTask(db *gorm.DB, id string, userID *string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var t Task
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where(&Task{ID: id}).First(&t).Error; err != nil {
			return err
		}
		if userID != nil && *userID != *t.AllocatedTo {
			return errors.BadRequest("tasks.NotAllocated", "Task is not allocated to the user")
		}
		if t.CancelledAt != nil {
			return errors.BadRequest("tasks.Cancelled", "Task has already been cancelled")
		}
		if t.CompletedAt != nil {
			return nil
		}

		tnow := time.Now()
		t.CompletedAt = &tnow
		return tx.Save(t).Error
	})
}

func DeferTask(db *gorm.DB, id string, deferredUntil time.Time) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var t Task
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where(&Task{ID: id}).First(&t).Error; err != nil {
			return err
		}
		if t.CancelledAt != nil {
			return errors.BadRequest("tasks.Cancelled", "Task has already been cancelled")
		}
		if t.CompletedAt != nil {
			return errors.BadRequest("tasks.Cancelled", "Task has already been completed")
		}

		t.DeferredUntil = &deferredUntil
		t.AllocatedTo = nil
		return tx.Save(t).Error
	})
}

func CancelTask(db *gorm.DB, id string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var t Task
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where(&Task{ID: id}).First(&t).Error; err != nil {
			return err
		}
		if t.CompletedAt != nil {
			return errors.BadRequest("tasks.Cancelled", "Task has already been completed")
		}
		if t.CancelledAt != nil {
			return nil
		}

		tnow := time.Now()
		t.CancelledAt = &tnow
		return tx.Save(t).Error
	})
}

func NextTask(db *gorm.DB, userID string, taskType string, tags []string) (*Task, error) {
	var result *Task
	return result, db.Transaction(func(tx *gorm.DB) error {
		// check the user isn't allocated to an existing task
		var existing Task
		q := tx.Model(&Task{}).
			Set("gorm:query_option", "FOR UPDATE").
			Where(&Task{AllocatedTo: &userID}).Where("created_at IS NULL AND cancelled_at IS NULL")
		if err := q.First(&existing).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		} else if err == nil {
			// check to see if the existing task matches the search criteria, if it does we can return this
			// task as the result, otherwise we need to unallocate the user from the task
			if taskType == existing.Type && includesOrNil(tags, existing.Tag) {
				result = &existing
				return nil
			}
			existing.AllocatedTo = nil
			if err := tx.Save(&existing).Error; err != nil {
				return err
			}
		}

		// lookup the next task which matches the criteria
		var next Task
		q = tx.Set("gorm:query_option", "FOR UPDATE").
			Where(&Task{Type: taskType}).
			Where("completed_at IS NULL AND cancelled_at IS NULL").
			Where("deferred_until IS NULL OR deferred_until <= ?", Time()).
			Order("COALESCE(ordering_time, created_at)")
		if len(tags) > 0 {
			q = q.Where("tag IN ?", tags)
		}
		if err := q.First(&next).Error; err == gorm.ErrRecordNotFound {
			return errors.NotFound("tasks.NoneLeft", "No tasks left matching this criteria")
		} else if err != nil {
			return err
		}

		// allocate the task to the user
		next.AllocatedTo = &userID
		if err := tx.Save(next).Error; err != nil {
			return err
		}

		// return the task
		result = &next
		return nil
	})
}

func UnassignUser(db *gorm.DB, userID string) error {
	return db.Model(&Task{}).
		Where(&Task{AllocatedTo: &userID}).
		Where("completed_at IS NULL AND cancelled_at IS NULL").
		Update("allocated_to", "NULL").
		Error
}

func RemoveSubject(db *gorm.DB, subjectID string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var tasks []*Task
		err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("completed_at IS NULL AND cancelled_at IS NULL AND allocated_to IS NULL AND subject_ids @> ARRAY[?]", subjectID).
			Find(&tasks).
			Error
		if err != nil {
			return err
		}

		tnow := time.Now()
		for _, t := range tasks {
			t.SubjectIDs = without(t.SubjectIDs, subjectID)

			// cancel the task if it has no subjects remaining
			if len(t.SubjectIDs) == 0 {
				t.CancelledAt = &tnow
			}

			if err := tx.Save(t).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func without(arr []string, val string) []string {
	var result []string
	for _, v := range arr {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

func includesOrNil(arr []string, val *string) bool {
	if arr == nil {
		return true
	}
	if val == nil {
		return false
	}
	for _, v := range arr {
		if v == *val {
			return true
		}
	}
	return false
}

func uniq(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]struct{})
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			u = append(u, val)
		}
	}
	return u
}
