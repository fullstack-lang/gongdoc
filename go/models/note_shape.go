package models

// NoteShape is a UML note in a class diagram
type NoteShape struct {
	Name string

	//gong:ident
	Identifier string

	Body          string
	X, Y          float64
	Width, Heigth float64
	Matched       bool // if a note with the same name has been found

	NoteLinks []*NoteLink
}
