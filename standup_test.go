package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestGetNextDailyOnMonday checks that at 7am UTC, the next daily is still on the same day
func TestGetNextDailyOnMonday(t *testing.T) {
	now := time.Date(2017, 01, 02, 07, 00, 00, 00, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 2, next.Day())
	assert.Equal(t, 9, next.Hour())
	assert.Equal(t, 20, next.Minute())
}

// TestGetNextDailyOnLateMonday checks that at 2pm UTC, the next daily is on the next day
func TestGetNextDailyOnLateMonday(t *testing.T) {
	now := time.Date(2017, 01, 02, 14, 00, 00, 00, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 3, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnTuesday checks that at 7am UTC, the next daily is still on the same day
func TestGetNextDailyOnTuesday(t *testing.T) {
	now := time.Date(2017, 01, 03, 07, 00, 00, 00, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 3, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}

// TestGetNextDailyOnLateTuesday checks that at 2pm UTC, the next daily is on the next day
func TestGetNextDailyOnLateTuesday(t *testing.T) {
	now := time.Date(2017, 01, 03, 14, 00, 00, 00, time.UTC)

	next := getNextDaily(now)

	assert.Equal(t, 4, next.Day())
	assert.Equal(t, 8, next.Hour())
	assert.Equal(t, 10, next.Minute())
}
