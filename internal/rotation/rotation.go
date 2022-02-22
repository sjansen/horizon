package rotation

import (
	"encoding/json"

	"github.com/tidwall/btree"
)

type Rotation struct {
	rotation
}

type rotation struct {
	Schedules btree.Map[int64, *Schedule] `json:"schedules"`
}

func New() *Rotation {
	return &Rotation{}
}

func (r *Rotation) AddSchedule(s *Schedule) *Rotation {
	key := s.Dates.GetDTStart().Unix()
	r.Schedules.Set(key, s)
	return r
}

func (r *Rotation) All() []*Turn {
	turns := []*Turn{}

	r.Schedules.Scan(func(key int64, s *Schedule) bool {
		turns = append(turns, s.All()...)
		return true
	})

	return turns
}

func (r *Rotation) MarshalJSON() ([]byte, error) {
	schedules := make([]*Schedule, 0, r.Schedules.Len())
	r.Schedules.Scan(func(k int64, v *Schedule) bool {
		schedules = append(schedules, v)
		return true
	})
	return json.Marshal(schedules)
}

func (r *Rotation) UnmarshalJSON(b []byte) error {
	var schedules []*Schedule
	if err := json.Unmarshal(b, &schedules); err != nil {
		return err
	}
	for _, s := range schedules {
		r.AddSchedule(s)
	}
	return nil
}
