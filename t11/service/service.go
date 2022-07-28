package service

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
)

func (s *Store) CreateEvent(date time.Time, msg string) int {
	event := EventCalendar{date, msg}
	s.Lock()
	defer s.Unlock()

	rand.Seed(time.Now().Unix())
	id := rand.Int()

	s.store[id] = event
	return id
}

func (s *Store) UpdateEvent(id int, date time.Time, msg string) (EventCalendar, error) {
	s.Lock()
	defer s.Unlock()
	if reflect.DeepEqual(s.store[id], EventCalendar{}) {
		return EventCalendar{}, errors.New("503: invalid element")
	}

	event := EventCalendar{date, msg}
	s.store[id] = event

	return event, nil
}

func (s *Store) DeleteEvent(id int) error {
	s.Lock()
	defer s.Unlock()

	if reflect.DeepEqual(s.store[id], EventCalendar{}) {
		return errors.New("503: invalid element")
	}

	delete(s.store, id)

	return nil
}

func (s *Store) EventsFor(date time.Time, duration int) ([]EventCalendar, error) {
	var res []EventCalendar

	for _, value := range s.store {
		if value.Date.Sub(date) >= time.Duration(duration*time.Now().Day()) {
			res = append(res, value)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("503: invalid event")
	}

	return res, nil
}
