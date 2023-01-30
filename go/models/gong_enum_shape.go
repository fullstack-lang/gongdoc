package models

const GongEnumShapeDefaultWidth = 240.0
const GongEnumShapeDefaultHeigth = 48.0

// GongEnumShape mirrors joint.shapes.uml.Class
// swagger:model GongEnumShape
type GongEnumShape struct {
	Name string

	Position *Position

	// Identifier is the identifier of the struct referenced by the shape in the modeled package
	//gong:ident
	Identifier string

	// gongEnumImpl, at run time, points to the gong object represented by the shape
	gongEnumImpl *GongEnumImpl

	// models of the composition of Field
	Fields []*Field

	// with and height of the shape when they are rendered on SVG or with jointjs
	// They are optional fields. they can be computed when empty
	Width, Heigth float64
}
