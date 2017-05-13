package main

import "encoding/json"

// ListEntry consists of a Category and Title for the Standup
// you can also vote on Entries in the future. Hopefully.
type ListEntry struct {
	ID    int
	Title string
	Votes uint32
}

// NewListEntry creates a new ListEntry
func NewListEntry(id int, title string) *ListEntry {
	return &ListEntry{
		id,
		title,
		0,
	}
}

// Implementations for Entry interface

// GetID returns the ID of the Entry
func (e *ListEntry) GetID() int {
	return e.ID
}

// GetTitle returns the Title of the Entry
func (e *ListEntry) GetTitle() string {
	return e.Title
}

// SetTitle sets the Title of the Entry
func (e *ListEntry) SetTitle(title string) {
	e.Title = title
}

// GetVotes returns the Votes of the Entry
func (e *ListEntry) GetVotes() uint32 {
	return e.Votes
}

// SetVotes sets the Votes of the Entry
func (e *ListEntry) SetVotes(votes uint32) {
	e.Votes = votes
}

// AddVote adds a Vote to the Entry
func (e *ListEntry) AddVote() {
	e.Votes++
}

// ToJSON marshals the ListEntry into a JSON string
func (e *ListEntry) ToJSON() []byte {
	res, _ := json.Marshal(e)

	return res
}
