package rrule

import (
	"time"

	"github.com/teambition/rrule-go"
)

type Frequency int

const (
	YEARLY  = Frequency(rrule.YEARLY)
	MONTHLY = Frequency(rrule.MONTHLY)
	WEEKLY  = Frequency(rrule.WEEKLY)
	DAILY   = Frequency(rrule.DAILY)
)

type WeekDay rrule.Weekday

var (
	MO = WeekDay(rrule.MO)
	TU = WeekDay(rrule.TU)
	WE = WeekDay(rrule.WE)
	TH = WeekDay(rrule.TH)
	FR = WeekDay(rrule.FR)
	SA = WeekDay(rrule.SA)
	SU = WeekDay(rrule.SU)
)

type Options struct {
	Begin      time.Time
	Until      time.Time
	Frequency  Frequency
	Interval   int
	WeekStart  WeekDay
	ByWeekDay  []WeekDay
	ByWeekNo   []int
	ByMonth    []int
	ByMonthDay []int
	ByYearDay  []int
}

func toDate(dt time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
}
