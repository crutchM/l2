package main

import (
	"github.com/cruthm/t11/handler"
	"github.com/cruthm/t11/middleware"
	"log"
	"net/http"
)

func initRoutes(mux *http.ServeMux, handler *handler.Handler) {
	mux.HandleFunc("/create-event", handler.CreateEvent)
	mux.HandleFunc("/update-event", handler.UpdateEvent)
	mux.HandleFunc("/delete-event", handler.DeleteEvent)
	mux.HandleFunc("/events-for-day", handler.EventsForDay)
	mux.HandleFunc("/events-for-week", handler.EventsForWeek)
	mux.HandleFunc("/event-for-month", handler.EventsForMonth)
}

func main() {
	server := http.NewServeMux()
	store := handler.NewHandler()
	initRoutes(server, store)
	hand := middleware.Log(server)

	log.Fatal(http.ListenAndServe("localhost:8080", hand))
}
