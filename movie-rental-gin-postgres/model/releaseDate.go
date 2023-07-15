package model

import (
	"strings"
	"time"
)

const ReleaseDateLayout = "02 Jan 2006"

type ReleaseDate struct {
	time.Time
}

func (rd *ReleaseDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		rd.Time = time.Time{}
		return
	}
	rd.Time, err = time.Parse(ReleaseDateLayout, s)
	return
}
