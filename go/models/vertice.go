package models

// Vertice mirrors joint.dia.Point
// swagger:model Vertice
type Vertice struct {
	X    float64
	Y    float64
	Name string // temporary
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (vertice *Vertice) SerializeToStage() {

	vertice.Stage()
}
