package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Categories is an array of Category
type Categories []*Category

// Category represents a Category with Entries
type Category struct {
	idCounter int
	ID        int
	Type      string
	Name      string
	Entries   Entries
}

// NewCategory creates a new Category
func NewCategory(id int, categoryType, name string) *Category {
	return &Category{
		0,
		id,
		name,
		categoryType,
		make(Entries, 0),
	}
}

// ToJSON marshals the Categories into a JSON string
func (c *Categories) ToJSON() []byte {
	res, _ := json.Marshal(c)

	return res
}

// ToJSON marshals the Category into a JSON string
func (c *Category) ToJSON() []byte {
	res, _ := json.Marshal(c)

	return res
}

// NewListEntry creates a new ListEntry inside the Category
func (c *Category) NewListEntry(title string) Entry {
	id := c.idCounter
	c.idCounter++

	var e Entry = NewListEntry(id, title)

	c.Entries = append(c.Entries, e)

	return e
}

// NewEventEntry creates a new EventEntry inside the Category
func (c *Category) NewEventEntry(title, where string, start time.Time) Entry {
	id := c.idCounter
	c.idCounter++

	var e Entry = NewEventEntry(id, title, start, where)

	c.Entries = append(c.Entries, e)

	return e
}

// GetEntryByID returns the Entry by its ID
func (c *Category) GetEntryByID(id int) Entry {
	for _, e := range c.Entries {
		if e.GetID() == id {
			return e
		}
	}
	return nil
}

// GetEntryByTitle returns the Entry by its Title
func (c *Category) GetEntryByTitle(title string) Entry {
	for _, e := range c.Entries {
		if e.GetTitle() == title {
			return e
		}
	}
	return nil
}

// RemoveEntryByID removed the Entry by its ID
func (c *Category) RemoveEntryByID(id int) error {
	pos := -1

	for idx, e := range c.Entries {
		if e.GetID() == id {
			pos = idx
			break
		}
	}

	if pos == -1 {
		return errors.New("Not found")
	}

	c.Entries = append(c.Entries[:pos], c.Entries[pos+1:]...)

	return nil
}

func apiGetEntries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("category"))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	c := standup.GetCategoryByID(id)
	if c == nil {
		w.WriteHeader(404)
		w.Write([]byte("Category not found"))
		return
	}

	w.Write(c.Entries.ToJSON())
}

// Create a new entry, if one with the same name exists
// add a vote
func apiCreateEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
		w.Write([]byte("Category not found"))
		return
	}

	var entry Entry
	if c.Type == "list" {
		var le *ListEntry
		if err := getJSON(r, &le); err != nil {
			w.WriteHeader(400)
			return
		}

		entry = le
	} else if c.Type == "events" {
		var te *struct {
			Title string
			Start string
			Where string
		}
		if err := getJSON(r, &te); err != nil {
			w.WriteHeader(400)
			return
		}

		start, err := time.ParseInLocation("2006-01-02T15:04", te.Start, timezone)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Invalid date Format. Required is 2006-01-02T15:04"))
			return
		}

		var ee EventEntry
		ee.Title = te.Title
		ee.Where = te.Where
		ee.Start = start.In(time.UTC)

		entry = &ee
	}

	if entry == nil {
		w.WriteHeader(500)
		return
	}

	if len(strings.TrimSpace(entry.GetTitle())) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Title cannot be empty"))
		return
	}

	existing := c.GetEntryByTitle(entry.GetTitle())

	if existing != nil {
		w.WriteHeader(302)
		existing.AddVote()

		w.Write(existing.ToJSON())
		return
	}

	if c.Type == "list" {
		w.Write(c.NewListEntry(entry.GetTitle()).ToJSON())
	} else if c.Type == "events" {
		ee, ok := entry.(*EventEntry)
		if !ok {
			w.WriteHeader(500)
			return
		}

		w.Write(c.NewEventEntry(ee.Title, ee.Where, ee.Start).ToJSON())
	}
}

func apiGetEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	eID, err := strconv.Atoi(p.ByName("entry"))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	e := c.GetEntryByID(eID)

	w.Write(e.ToJSON())
}

// apiUpdateEntry updates an existing entry. If the Title is changed to an existing one,
// an error will be thrown
func apiUpdateEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	eID, err := strconv.Atoi(p.ByName("entry"))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	e := c.GetEntryByID(eID)

	if e == nil {
		w.WriteHeader(404)
		return
	}

	var entry Entry
	if c.Type == "list" {
		var le *ListEntry
		if err := getJSON(r, &le); err != nil {
			w.WriteHeader(400)
			return
		}

		entry = le
	} else if c.Type == "events" {
		var ee *EventEntry
		if err := getJSON(r, &ee); err != nil {
			w.WriteHeader(400)
			return
		}

		entry = ee
	}

	if entry == nil {
		w.WriteHeader(500)
		return
	}

	if len(strings.TrimSpace(entry.GetTitle())) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Title cannot be empty"))
		return
	}

	if e.GetTitle() != entry.GetTitle() {
		if ee := c.GetEntryByTitle(entry.GetTitle()); ee != nil {
			w.WriteHeader(409)
			w.Write([]byte("Entry with this Title exists"))
			return
		}
	}

	e.SetTitle(entry.GetTitle())
	e.SetVotes(entry.GetVotes())

	w.Write(e.ToJSON())
}

func apiVoteEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	eID, err := strconv.Atoi(p.ByName("entry"))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	e := c.GetEntryByID(eID)

	if e == nil {
		w.WriteHeader(404)
		return
	}

	e.AddVote()

	w.Write(e.ToJSON())
}

func apiDeleteEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	eID, err := strconv.Atoi(p.ByName("entry"))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	if c.RemoveEntryByID(eID) != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}
