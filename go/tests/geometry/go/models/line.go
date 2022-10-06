package models

import "time"

// Line is a the line between Start and End
type Line struct {
	Name string

	Start *Point
	End   *Point

	CreationDate time.Time

	JourneyTime time.Duration

	LineType LineType

	VeryLongLongLongLongLongLongField string
}
