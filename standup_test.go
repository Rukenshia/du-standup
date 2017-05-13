package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestGetNextDailyOnMonday checks that at 7am UTC, the next daily is still on the same day
func TestGetNextDailyOnMonday(t *testing.T) {
	now := time.Date(2017, 1, 2, 7, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 2, next.Day())
	assert.Equal(t, 9, next.Hour())
	assert.Equal(t, 20, next.Minute())
}

// TestGetNextDailyOnLateMonday checks that at 2pm UTC, the next daily is on the next day
func TestGetNextDailyOnLateMonday(t *testing.T) {
	now := time.Date(2017, 1, 2, 14, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 3, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnTuesday checks that at 7am UTC, the next daily is still on the same day
func TestGetNextDailyOnTuesday(t *testing.T) {
	now := time.Date(2017, 1, 3, 7, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 3, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnLateTuesday checks that at 2pm UTC, the next daily is on the next day
func TestGetNextDailyOnLateTuesday(t *testing.T) {
	now := time.Date(2017, 1, 3, 14, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 4, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnWednesday checks that at 7am UTC, the next daily is still on the same day
func TestGetNextDailyOnWednesday(t *testing.T) {
	now := time.Date(2017, 1, 4, 7, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 4, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnLateWednesday checks that at 2pm UTC, the next daily is on the next day
func TestGetNextDailyOnLateWednesday(t *testing.T) {
	now := time.Date(2017, 1, 4, 14, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 5, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnThursday checks that at 7am UTC, the next daily is still on the same day
func TestGetNextDailyOnThursday(t *testing.T) {
	now := time.Date(2017, 1, 5, 7, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 5, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnLateThursday checks that at 2pm UTC, the next daily is on the next day
func TestGetNextDailyOnLateThursday(t *testing.T) {
	now := time.Date(2017, 1, 5, 14, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 6, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnFriday checks that at 7am UTC, the next daily is still on the same day
func TestGetNextDailyOnFriday(t *testing.T) {
	now := time.Date(2017, 1, 6, 7, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 6, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnLateFriday checks that at 2pm UTC, the next daily is on monday
func TestGetNextDailyOnLateFriday(t *testing.T) {
	now := time.Date(2017, 1, 6, 14, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 9, next.Day())
	assert.Equal(t, 9, next.Hour())
	assert.Equal(t, 20, next.Minute())
}

// TestGetNextDailyOnSaturday checks that at 2pm UTC, the next daily is on monday
func TestGetNextDailyOnSaturday(t *testing.T) {
	now := time.Date(2017, 1, 7, 14, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 9, next.Day())
	assert.Equal(t, 9, next.Hour())
	assert.Equal(t, 20, next.Minute())
}

// TestGetNextDailyOnSunday checks that at 2pm UTC, the next daily is on monday
func TestGetNextDailyOnSunday(t *testing.T) {
	now := time.Date(2017, 1, 8, 14, 0, 0, 0, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 9, next.Day())
	assert.Equal(t, 9, next.Hour())
	assert.Equal(t, 20, next.Minute())
}

// TestIsStandupExpired checks whether it returns true when 'now' is past the Expiry date and
// there is a new date available from getNextDaily
func TestIsStandupExpired(t *testing.T) {
	s := NewStandup()

	now := s.Expires.Add(time.Hour)

	assert.Equal(t, true, s.IsExpired(now))
}
