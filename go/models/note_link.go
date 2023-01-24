package models

// A NoteLink is a visual link from a shape to
// a Link or to a Classshape
//
// The end point of the link is computed from the Type
type NoteLink struct {
	Name string

	// Identifier of the target shape / link of the note link
	//gong:ident
	Identifier string

	// Type indicates wether this is a note link (a dotted line) to classshape
	Type ReferenceType

	Classshape *Classshape
	Link       *Link

	// Vertices at the middle
	Middlevertice *Vertice
}
