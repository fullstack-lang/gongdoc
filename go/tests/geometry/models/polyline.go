package models

type Polyline struct {
	Name              string
	Points            []*Point
	Points_additional []*Point
	Start_Point       *Point
	Follower          *Polyline
	Followers         []*Polyline
}
