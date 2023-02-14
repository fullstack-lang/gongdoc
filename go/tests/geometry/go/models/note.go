package models

// GONGNOTE(LongNodeOnModels): This is an example of a note that
// could be displayed on a diagram.
//
// It could explain one aspect of the model
// for intance, describing relations between structs
//
// The text of a UML note refers a comment with the GONGNOTE keyword which is
// a special case of go Note convention. See example
// for details in the go code of the models.
//
// This follows the go code convention described in https://pkg.go.dev/go/doc#Note
//
// "A Note represents a marked comment starting with "MARKER(uid): note body".
// Any note with a marker of 2 or more upper case [A-Z] letters and a uid of at least one character is recognized.
// The ":" following the uid is optional. Notes are collected in the Package.Notes map indexed by the notes marker."
//
// In the UML diagram, the size of the note is automaticaly computed from the note
// number of lines (for the width) and the number of characters per line (for the height)
// in the go code
const LongNodeOnModels = ""

// GONGNOTE(ShortNodeOnModels): this is an example of a short note
// It uses the DocLink convention for referencing Identifiers
// In this case [models.Line], [models.Point] and [models.Line.Start] [models.LineTypeString]
// are referenced in the go code
const ShortNodeOnModels = ""

// GONGNOTE(MarkdownNodeOnModels): this is an example of an Markdown note
// It uses the DocLink convention for referencing Identifiers
// In this case [models.Line], [models.Point] and [models.Line.Start] [models.LineTypeString]
// are referenced in the go code
//
// # This is a title
//
// ## This is a sub title
//
// - some comment
// Some comment
const MarkdownNodeOnModels = ""
