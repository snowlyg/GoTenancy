package publish2

import (
	"time"

	"GoTenancy/libs/validations"
	"github.com/jinzhu/gorm"
)

type Schedule struct {
	ScheduledStartAt *time.Time `gorm:"index"`
	ScheduledEndAt   *time.Time `gorm:"index"`
	ScheduledEventID *uint
}

func (schedule *Schedule) GetScheduledStartAt() *time.Time {
	return schedule.ScheduledStartAt
}

func (schedule *Schedule) SetScheduledStartAt(t *time.Time) {
	schedule.ScheduledStartAt = t
}

func (schedule *Schedule) GetScheduledEndAt() *time.Time {
	return schedule.ScheduledEndAt
}

func (schedule *Schedule) SetScheduledEndAt(t *time.Time) {
	schedule.ScheduledEndAt = t
}

func (schedule *Schedule) GetScheduledEventID() *uint {
	return schedule.ScheduledEventID
}

type ScheduledInterface interface {
	GetScheduledStartAt() *time.Time
	SetScheduledStartAt(*time.Time)
	GetScheduledEndAt() *time.Time
	SetScheduledEndAt(*time.Time)
	GetScheduledEventID() *uint
}

type ScheduledEvent struct {
	gorm.Model
	Name             string
	ScheduledStartAt *time.Time
	ScheduledEndAt   *time.Time
}

func (scheduledEvent ScheduledEvent) ToParam() string {
	return "scheduled_events"
}

func (scheduledEvent ScheduledEvent) BeforeSave(tx *gorm.DB) {
	if scheduledEvent.Name == "" {
		tx.AddError(validations.NewError(scheduledEvent, "Name", "Name can not be empty"))
	}
}
