package models

import "time"

// A Point is a 2 dimensional coordinate (X,Y) on a plane
type Point struct {
	X, Y      float64
	CreatedAt time.Time
}
