package model

import (
	"testing"
	"time"
)

func TestUnmarshallShouldReturnCorrectTimeForExpectedDateFormat(t *testing.T) {
	d := "24 Mar 1972"
	want, _ := time.Parse(ReleaseDateLayout, d)

	rd := ReleaseDate{}

	err := rd.UnmarshalJSON([]byte(d))

	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if rd.Time != want {
		t.Errorf("expected time %s but got %s", want, rd.Time)
	}
}
