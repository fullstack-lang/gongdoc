package models

type AnchoredText struct {
	Name    string
	Content string

	X_Offset float64
	Y_Offset float64

	Presentation
	Animates []*Animate
}
