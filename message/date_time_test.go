package message

import (
	"testing"
	"time"
)

func TestDateTime(t *testing.T) {
	// Define a test time
	testTime := time.Date(1997, time.November, 21, 9, 55, 6, 0, time.FixedZone("CST", -6*60*60))

	// Create a new date_time instance
	dt := NewDateTime(testTime)

	// Check if the time method returns the correct time
	if dt.time() != testTime {
		t.Errorf("expected time %v, got %v", testTime, dt.time())
	}

	// Check if the String method returns the correct formatted string
	expectedString := "Fri, 21 Nov 1997 09:55:06 -0600"
	if dt.String() != expectedString {
		t.Errorf("expected string %q, got %q", expectedString, dt.String())
	}
}
