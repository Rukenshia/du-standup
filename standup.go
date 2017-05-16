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

// getNextDaily returns the expiry date of the next daily.
// Monday: 9:20 UTC
// Tue-Fri: 08:10 UTC
// getNextDaily returns the next Monday if it is Friday after the current daily
// we don't work on saturdays after all :)
func getNextDaily(now time.Time) time.Time {

	hour := 9
	minute := 10

	if (now.Weekday() == time.Monday && (now.Hour() < 10 || now.Hour() == 9 && now.Minute() < 20)) ||
		(now.Weekday() == time.Friday && (now.Hour() > hour || now.Hour() == hour && now.Minute() >= minute)) ||
		now.Weekday() == time.Saturday ||
		now.Weekday() == time.Sunday {

		hour = 10
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

	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, timezone)
}

// NewStandup creates a new Standup with the next available Date
func NewStandup() Standup {
	return Standup{
		0,
		getNextDaily(time.Now().In(timezone)),
		Categories{
			NewCategory(0, "visitors", "list"),
			NewCategory(1, "interests", "list"),
			NewCategory(2, "needs", "list"),
			NewCategory(3, "pains", "list"),
			NewCategory(4, "events", "events"),
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

// IsExpired returns whether the daily is already expired
func (s *Standup) IsExpired(now time.Time) bool {
	return s.Expires.Format("2006-01-02") != getNextDaily(now.In(timezone)).Format("2006-01-02")
}

// StandupMiddleware to automatically regenerate the next Standup
func StandupMiddleware(handler http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		if standup.IsExpired(time.Now()) {
			// save events that are not expiring on this day
			var events []Entry

			if c := standup.GetCategoryByName("events"); c != nil {
				for _, entry := range c.Entries {
					event, ok := entry.(*EventEntry)
					if !ok {
						continue
					}

					if event.Start.Day() != standup.Expires.Day() {
						events = append(events, entry)
					}
				}
			}

			standup = NewStandup()

			if c := standup.GetCategoryByName("events"); c != nil {
				// and append them again. yay.
				c.Entries = append(c.Entries, events...)
			}

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
