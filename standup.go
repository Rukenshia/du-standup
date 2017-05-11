package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Standup is the current Standup Meeting
// It is denoted with an Expiry date
type Standup struct {
	idCounter int
	Expires   time.Time
	Entries   Entries
}

func getNextDaily() time.Time {
	now := time.Now()

	hour := 8
	minute := 10

	if (now.Weekday() == time.Friday && (now.Hour() > hour || now.Hour() == hour && now.Minute() >= minute)) || now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		hour = 9
		minute = 20

	}

	if now.Hour() > hour || now.Hour() == hour && now.Minute() >= minute {
		if now.Weekday() == time.Friday {
			now = now.Add(time.Hour * 24 * 3)
		} else if now.Weekday() == time.Saturday {
			now = now.Add(time.Hour * 24 * 2)
		} else {
			now = now.Add(time.Hour * 24)
		}
	}

	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.UTC)
}

// NewStandup creates a new Standup with the next available Date
func NewStandup() Standup {
	return Standup{
		0,
		getNextDaily(),
		make(Entries, 0),
	}
}

// NewEntry creates a new Entry for this Standup
func (s *Standup) NewEntry(category, title string) Entry {
	s.idCounter++

	entry := Entry{
		s.idCounter,
		category,
		title,
		0,
	}

	s.Entries = append(s.Entries, entry)
	return entry
}

// ToJSON marshals the Standup into a JSON string
func (s *Standup) ToJSON() []byte {
	res, _ := json.Marshal(s)

	return res
}

// StandupMiddleware to automatically regenerate the next Standup
func StandupMiddleware(handler http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		if standup.Expires.Format("2006-01-02") != getNextDaily().Format("2006-01-02") {
			standup = NewStandup()

			log.Printf("Generated new Standup for %s", standup.Expires.Format("2006-01-02"))
		}

		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(mw)
}

func apiGetStandup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write(standup.ToJSON())
}
