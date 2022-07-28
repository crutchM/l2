package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		next.ServeHTTP(writer, request)
		log.Printf("%s %s %s", request.Method, request.RequestURI, time.Since(start))
	})
}

func ParseDataFromBody(r *http.Request) (int, time.Time, string, error) {
	var id int
	var date time.Time
	var mes string

	idString := r.FormValue("id")
	if idString != "" {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid int")
		}

		id = idInt
	}

	dateString := r.FormValue("date")
	if dateString != "" {
		dateString += "T00:00:00Z"
		dateTime, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: invalid date")
		}

		date = dateTime
	}

	mes = r.FormValue("mes")

	return id, date, mes, nil
}

func CreateJson(w http.ResponseWriter, obj any) {
	result := struct {
		Result any
	}{Result: obj}

	parsed, err := json.Marshal(&result)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsed)
}

func ParseQueryParams(r *http.Request) time.Time {
	dateF := r.FormValue("date") + "T00:00:00Z"
	date, err := time.Parse(time.RFC3339, dateF)
	if err != nil {
		fmt.Println(err)
	}

	return date
}
