package main

import (
	"encoding/json"
)

// Entries represent an Array of ListEntry
type Entries []Entry

// Entry provides an Entry Interface
type Entry interface {
	// Have to use Get... here so we can still easily JSON marshal the ID prop
	// sadly not "Effective Go"

	GetID() int
	GetTitle() string
	SetTitle(string)

	GetVotes() uint32
	SetVotes(uint32)
	AddVote()

	ToJSON() []byte
}

// ToJSON marshals the Entries into a JSON string
func (e *Entries) ToJSON() []byte {
	res, _ := json.Marshal(e)

	return res
}
