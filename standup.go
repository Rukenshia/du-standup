package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Standup is the current Standup Meeting
// It is denoted with an Expiry date
type Standup struct {
	idCounter  int
	Expires    time.Time
	Categories Categories
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
		Categories{
			NewCategory(0, "interests"),
			NewCategory(1, "needs"),
			NewCategory(2, "events"),
		},
	}
}

// ToJSON marshals the Standup into a JSON string
func (s *Standup) ToJSON() []byte {
	res, _ := json.Marshal(s)

	return res
}

// GetCategoryByID gets a Category by its id
func (s *Standup) GetCategoryByID(id int) *Category {
	for _, c := range s.Categories {
		if c.ID == id {
			return c
		}
	}
	return nil
}

// GetCategoryByName gets a Category by its name
func (s *Standup) GetCategoryByName(name string) *Category {
	for _, c := range s.Categories {
		if c.Name == name {
			return c
		}
	}
	return nil
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

func apiGetCategories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write(standup.Categories.ToJSON())
}

func apiGetCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var c *Category
	id, err := strconv.Atoi(p.ByName("category"))
	if err == nil {
		c = standup.GetCategoryByID(id)

		// maybe the category name is a number
		if c == nil {
			c = standup.GetCategoryByName(p.ByName("category"))

		}
	} else {
		// try to find it by name
		c = standup.GetCategoryByName(p.ByName("category"))
	}

	if c == nil {
		w.WriteHeader(404)
		return
	}

	w.Write(c.ToJSON())
}
