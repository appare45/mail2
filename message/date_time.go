package message

import (
	"time"
)

type date_time time.Time

func NewDateTime(t time.Time) date_time {
	return date_time(t)
}

func (d *date_time) time() time.Time {
	return time.Time(*d)
}

func (d date_time) String() string {
	return d.time().Format("Mon, 02 Jan 2006 15:04:05 -0700")
}
