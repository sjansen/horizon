package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sjansen/horizon/internal/rotation"
	"github.com/sjansen/horizon/internal/rrule"
)

func main() {
	location, err := time.LoadLocation("America/Denver")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	r := rotation.New()

	s, err := createSchedule1(location)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	r.AddSchedule(s)

	s, err = createSchedule2(location)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	r.AddSchedule(s)

	b, err := json.Marshal(r)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	printJSON(b)

	rotation := &rotation.Rotation{}
	json.Unmarshal(b, &rotation)
	printRotation(rotation)
}

func createSchedule1(location *time.Location) (*rotation.Schedule, error) {
	rrule, err := rrule.New(&rrule.Options{
		Begin:     time.Date(2021, 11, 2, 0, 0, 0, 0, location),
		Until:     time.Date(2022, 4, 1, 0, 0, 0, 0, location),
		Frequency: rrule.WEEKLY,
		ByWeekDay: []rrule.WeekDay{rrule.TU},
		Interval:  2,
	})
	if err != nil {
		return nil, err
	}

	schedule := rotation.NewSchedule(
		3, []string{"a", "b", "c", "d", "e"}, rrule,
	).ExDate(
		time.Date(2021, 12, 28, 0, 0, 0, 0, location),
	)

	return schedule, nil
}

func createSchedule2(location *time.Location) (*rotation.Schedule, error) {
	rrule, err := rrule.New(&rrule.Options{
		Begin:     time.Date(2022, 4, 1, 0, 0, 0, 0, location),
		Until:     time.Date(2022, 7, 5, 0, 0, 0, 0, location),
		Frequency: rrule.WEEKLY,
		ByWeekDay: []rrule.WeekDay{rrule.TU},
		Interval:  2,
	})
	if err != nil {
		return nil, err
	}

	schedule := rotation.NewSchedule(
		4, []string{"a", "b", "c", "w", "x", "y", "z"}, rrule,
	).ExDate(
		time.Date(2022, 5, 10, 0, 0, 0, 0, location),
	).RDate(
		time.Date(2022, 5, 11, 0, 0, 0, 0, location),
	)

	return schedule, nil
}

func printJSON(b []byte) {
	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	out.WriteTo(os.Stdout)
	fmt.Println()
}

func printRotation(r *rotation.Rotation) {
	turns := r.All()
	for _, turn := range turns {
		fmt.Print(turn.Time, " ")
		for _, item := range turn.List {
			fmt.Print(" ", item)
		}
		fmt.Println()
	}
}
