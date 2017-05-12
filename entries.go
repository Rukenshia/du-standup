package main

import (
	"encoding/json"
)

// Entries represent an Array of Entry
type Entries []Entry

// Entry consists of a Category and Title for the Standup
// you can also vote on Entries in the future. Hopefully.
type Entry struct {
	ID    int
	Title string
	Votes uint32
}

// NewEntry creates a new Entry
func NewEntry(id int, title string) Entry {
	return Entry{
		id,
		title,
		0,
	}
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
