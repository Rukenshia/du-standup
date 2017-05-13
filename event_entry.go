package main

import (
	"encoding/json"
	"time"
)

// EventEntry consists of a Category and Title for the Standup
// you can also vote on Entries in the future. Hopefully.
type EventEntry struct {
	ID    int
	Title string
	Start time.Time
	Where string
	Votes uint32
}

// NewEventEntry creates a new EventEntry
func NewEventEntry(id int, title string, start time.Time, where string) *EventEntry {
	return &EventEntry{
		id,
		title,
		start,
		where,
		0,
	}
}

// Implementations for Entry interface

// GetID returns the ID of the Entry
func (e *EventEntry) GetID() int {
	return e.ID
}

// GetTitle returns the Title of the Entry
func (e *EventEntry) GetTitle() string {
	return e.Title
}

// SetTitle sets the Title of the Entry
func (e *EventEntry) SetTitle(title string) {
	e.Title = title
}

// GetVotes returns the Votes of the Entry
func (e *EventEntry) GetVotes() uint32 {
	return e.Votes
}

// SetVotes sets the Votes of the Entry
func (e *EventEntry) SetVotes(votes uint32) {
	e.Votes = votes
}

// AddVote adds a Vote to the Entry
func (e *EventEntry) AddVote() {
	e.Votes++
}

// ToJSON marshals the EventEntry into a JSON string
func (e *EventEntry) ToJSON() []byte {
	res, _ := json.Marshal(e)

	return res
}
