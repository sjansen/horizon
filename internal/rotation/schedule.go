package rotation

import (
	"encoding/json"
	"time"

	"github.com/sjansen/horizon/internal/rrule"
)

type Schedule struct {
	schedule
}

type schedule struct {
	Dates *rrule.Set `json:"dates"`
	List  []string   `json:"list"`
	Size  int        `json:"size"`
}

func NewSchedule(size int, list []string, r *rrule.RRule) *Schedule {
	s := &Schedule{}
	s.Dates = rrule.NewSet(r)
	s.List = list
	s.Size = size
	return s
}

func (s *Schedule) All() []*Turn {
	turns := []*Turn{}

	length := len(s.List)
	offset := 0
	for _, begin := range s.Dates.All() {
		turn := &Turn{
			Time: begin,
			List: make([]string, 0, s.Size),
		}
		for i := 0; i < s.Size; i++ {
			idx := (offset + i) % length
			turn.List = append(turn.List, s.List[idx])
		}
		offset += s.Size
		turns = append(turns, turn)
	}

	return turns
}

func (s *Schedule) ExDate(dt time.Time) *Schedule {
	s.Dates.ExDate(dt)
	return s
}

func (s *Schedule) RDate(dt time.Time) *Schedule {
	s.Dates.RDate(dt)
	return s
}

func (s *Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.schedule)
}

func (s *Schedule) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.schedule)
}
