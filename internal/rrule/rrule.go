package rrule

import (
	"github.com/teambition/rrule-go"
)

type RRule struct {
	*rrule.RRule
}

func New(o *Options) (*RRule, error) {
	ro := rrule.ROption{
		Dtstart:   toDate(o.Begin),
		Until:     toDate(o.Until),
		Freq:      rrule.Frequency(o.Frequency),
		Interval:  o.Interval,
		Wkst:      rrule.Weekday(o.WeekStart),
		Byweekday: make([]rrule.Weekday, len(o.ByWeekDay)),
	}
	for idx, wd := range o.ByWeekDay {
		ro.Byweekday[idx] = rrule.Weekday(wd)
	}
	copy(ro.Byweekno, o.ByWeekNo)
	copy(ro.Bymonth, o.ByMonth)
	copy(ro.Bymonthday, o.ByMonthDay)
	copy(ro.Byyearday, o.ByYearDay)

	r, err := rrule.NewRRule(ro)
	if err != nil {
		return nil, err
	}

	rrule := &RRule{
		RRule: r,
	}

	return rrule, nil
}
