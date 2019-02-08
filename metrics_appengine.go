// +build appengine

package goa

import (
	"time"
)

// IncrCounter; Not supported in Google App Engine
func IncrCounter(key []string, val float32) {
	// Do nothing
}

// MeasureSince; Not supported in Google App Engine
func MeasureSince(key []string, start time.Time) {
	// Do nothing
}
