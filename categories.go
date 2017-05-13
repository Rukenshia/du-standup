package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Categories is an array of Category
type Categories []*Category

// Category represents a Category with Entries
type Category struct {
	idCounter int
	ID        int
	Name      string
	Entries   Entries
}

// NewCategory creates a new Category
func NewCategory(id int, name string) *Category {
	return &Category{
		0,
		id,
		name,
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

// NewEntry creates a new Entry inside the Category
func (c *Category) NewEntry(title string) *Entry {
	id := c.idCounter
	c.idCounter++

	e := NewEntry(id, title)

	c.Entries = append(c.Entries, e)

	return e
}

// GetEntryByID returns the Entry by its ID
func (c *Category) GetEntryByID(id int) *Entry {
	for _, e := range c.Entries {
		if e.ID == id {
			return e
		}
	}
	return nil
}

// GetEntryByID returns the Entry by its Title
func (c *Category) GetEntryByTitle(title string) *Entry {
	for _, e := range c.Entries {
		if e.Title == title {
			return e
		}
	}
	return nil
}

// RemoveEntryByID removed the Entry by its ID
func (c *Category) RemoveEntryByID(id int) error {
	pos := -1

	for idx, e := range c.Entries {
		if e.ID == id {
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

	var entry *Entry

	if err := getJSON(r, &entry); err != nil {
		w.WriteHeader(400)
		return
	}

	if len(strings.TrimSpace(entry.Title)) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Title cannot be empty"))
		return
	}

	existing := c.GetEntryByTitle(entry.Title)

	if existing != nil {
		w.WriteHeader(302)
		existing.Votes++

		w.Write(existing.ToJSON())
		return
	}

	entry = c.NewEntry(entry.Title)

	w.Write(entry.ToJSON())
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

	if err := getJSON(r, &entry); err != nil {
		w.WriteHeader(400)
		return
	}

	if len(strings.TrimSpace(entry.Title)) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Title cannot be empty"))
		return
	}

	if e.Title != entry.Title {
		if ee := c.GetEntryByTitle(entry.Title); ee != nil {
			w.WriteHeader(409)
			w.Write([]byte("Entry with this Title exists"))
			return
		}
	}

	e.Title = entry.Title
	e.Votes = entry.Votes

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

	e.Votes++

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
