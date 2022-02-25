package rrule

import (
	"encoding/json"
	"time"

	"github.com/teambition/rrule-go"
)

type Set struct {
	*rrule.Set
}

func NewSet(r *RRule) *Set {
	s := &Set{
		Set: &rrule.Set{},
	}
	s.Set.RRule(r.RRule)
	return s
}

func (s *Set) ExDate(dt time.Time) *Set {
	s.Set.ExDate(dt)
	return s
}

func (s *Set) RDate(dt time.Time) *Set {
	s.Set.RDate(dt)
	return s
}

func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Set.String())
}

func (s *Set) UnmarshalJSON(b []byte) error {
	var tmp string
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	parsed, err := rrule.StrToRRuleSet(tmp)
	if err != nil {
		return err
	}

	s.Set = parsed
	return nil
}
