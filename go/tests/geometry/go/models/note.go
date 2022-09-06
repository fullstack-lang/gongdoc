package models

// GONGNOTE(Note on the models): This is an example of a note that
// could be displayed on a diagram.
//
// It could explain one aspect of the model
// for intance, describing relations between structs
//
// The text of a UML note refers a comment with the GONGNOTE keyword (see example
// for details) in the go code of the models. This follows the go code convention
// https://pkg.go.dev/go/doc#Note
//
// A Note represents a marked comment starting with "MARKER(uid): note body".
// Any note with a marker of 2 or more upper case [A-Z] letters and a uid of at least one character is recognized.
// The ":" following the uid is optional. Notes are collected in the Package.Notes map indexed by the notes marker.
//
// In the UML diagram, the size of the note is automaticaly computed from the note
// number of lines (for the width) and the number of characters per line (for the height)
// in the go code
//

// GONGNOTE(Short note on the models): this is an of a short note
