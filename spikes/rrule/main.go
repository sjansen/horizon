package main

import (
	"fmt"
	"time"

	"github.com/teambition/rrule-go"
)

func main() {
	fmt.Println("Daily, for 4 occurrences.")
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   4,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	fmt.Println(r)
	fmt.Println(r.Between(
		time.Date(1997, 9, 3, 0, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 5, 0, 0, 0, 0, time.UTC),
		true,
	))
	fmt.Println(r.All())

	fmt.Println("==========")

	fmt.Println("Every four years, the first Tuesday after a Monday in November, 3 occurrences (U.S. Presidential Election day).")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.YEARLY,
		Interval:   4,
		Count:      3,
		Bymonth:    []int{11},
		Byweekday:  []rrule.Weekday{rrule.TU},
		Bymonthday: []int{2, 3, 4, 5, 6, 7, 8},
		Dtstart:    time.Date(2020, 11, 3, 9, 0, 0, 0, time.UTC)})
	fmt.Println(r)
	fmt.Println(r.All())

	fmt.Println("==========")

	fmt.Println("Daily, for 7 days, jumping Saturday and Sunday occurrences.")
	set := &rrule.Set{}
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	fmt.Println(set)
	fmt.Println(set.All())

	fmt.Println("==========")

	fmt.Println("Weekly, for 4 weeks, plus one time on day 12, and not on day 16.")
	set = &rrule.Set{}
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.WEEKLY,
		Count:   4,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.RDate(time.Date(1997, 9, 12, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC))
	fmt.Println(set)
	fmt.Println(set.All())

	fmt.Println("==========")

	r, _ = rrule.StrToRRule("FREQ=DAILY;INTERVAL=10;COUNT=5")
	fmt.Println(r)
	fmt.Println(r.All())

	fmt.Println("==========")

	s, _ := rrule.StrToRRuleSet(
		`RRULE:FREQ=DAILY;INTERVAL=10;COUNT=5
RDATE:20060102T150405Z`,
	)
	fmt.Println(s)
	fmt.Println(s.All())
}
