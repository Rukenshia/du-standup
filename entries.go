package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Entries represent an Array of Entry
type Entries []Entry

// Entry consists of a Category and Title for the Standup
// you can also vote on Entries in the future. Hopefully.
type Entry struct {
	ID       int
	Category string
	Title    string
	Votes    uint32
}

// ToJSON marshals the Entries into a JSON string
func (e *Entries) ToJSON() []byte {
	res, _ := json.Marshal(e)

	return res
}

// ToJSON marshals the Entry into a JSON string
func (e *Entry) ToJSON() []byte {
	res, _ := json.Marshal(e)

	return res
}

func apiGetEntries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write(standup.Entries.ToJSON())
}

func apiCreateEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var entry Entry

	if err := getJSON(r, &entry); err != nil {
		w.WriteHeader(400)
		return
	}

	entry = standup.NewEntry(entry.Category, entry.Title)

	w.Write(entry.ToJSON())
}

func apiUpdateEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))

	if err != nil {
		w.WriteHeader(400)
		return
	}

	pos := -1

	for idx, entry := range standup.Entries {
		if entry.ID == id {
			pos = idx
			break
		}
	}

	if pos == -1 {
		w.WriteHeader(404)
		return
	}

	var entry Entry

	if err := getJSON(r, &entry); err != nil {
		w.WriteHeader(400)
		return
	}

	existing := &standup.Entries[pos]

	existing.Category = entry.Category
	existing.Title = entry.Title
	existing.Votes = entry.Votes

	w.Write(existing.ToJSON())
}

func apiDeleteEntry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))

	if err != nil {
		w.WriteHeader(400)
		return
	}

	pos := -1

	for idx, entry := range standup.Entries {
		if entry.ID == id {
			pos = idx
			break
		}
	}

	if pos == -1 {
		w.WriteHeader(404)
		return
	}

	standup.Entries = append(standup.Entries[:pos], standup.Entries[pos+1:]...)

	w.WriteHeader(200)
}
