package handler

import (
	"fmt"
	"github.com/cruthm/t11/middleware"
	"github.com/cruthm/t11/service"
	"net/http"
)

type Handler struct {
	store *service.Store
}

func NewHandler() *Handler {
	return &Handler{store: service.NewStore()}
}

func (s *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	_, date, msg, err := middleware.ParseDataFromBody(r)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id := s.store.CreateEvent(date, msg)
	middleware.CreateJson(w, id)
}
func (s *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, date, msg, err := middleware.ParseDataFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	event, err := s.store.UpdateEvent(id, date, msg)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	middleware.CreateJson(w, event)
}
func (s *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _, _, err := middleware.ParseDataFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = s.store.DeleteEvent(id)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	middleware.CreateJson(w, fmt.Sprint("element ", id, "was deleted successfully"))
}

func (s *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	date := middleware.ParseQueryParams(r)
	events, err := s.store.EventsFor(date, 0)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	middleware.CreateJson(w, events)

}
func (s *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := middleware.ParseQueryParams(r)
	events, err := s.store.EventsFor(date, 7)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	middleware.CreateJson(w, events)

}
func (s *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := middleware.ParseQueryParams(r)
	events, err := s.store.EventsFor(date, 30)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	middleware.CreateJson(w, events)

}
