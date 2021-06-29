package main

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/teambition/rrule-go"
)

func main() {
	cal := ics.NewCalendar()
	cal.SetProductId("-//example.com//Horizon 0.1//EN")
	cal.SetMethod(ics.MethodRequest)

	event := cal.AddEvent("12345")
	event.SetCreatedTime(time.Now())
	event.SetDtStampTime(time.Now())
	event.SetModifiedAt(time.Now())
	event.SetStartAt(time.Now())
	event.SetEndAt(time.Now().Add(30 * time.Minute))
	r := &rrule.ROption{
		Freq:  rrule.WEEKLY,
		Count: 4,
	}
	event.AddRrule(r.String())
	event.SetSummary("Summary")
	event.SetLocation("Address")
	event.SetDescription("Description")
	event.SetURL("https://example.com/calendar.ics")
	event.SetOrganizer("sender@example.com", ics.WithCN("John Doe"))
	event.AddAttendee(
		"receiver@example.com",
		ics.CalendarUserTypeIndividual,
		ics.ParticipationStatusNeedsAction,
		ics.ParticipationRoleReqParticipant,
		ics.WithRSVP(true),
	)
	fmt.Println(cal.Serialize())
}
