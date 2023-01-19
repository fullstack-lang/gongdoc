package models

// UmlState mirrors joint.shapes.uml.UmlState
// swagger:model UmlState
type UmlState struct {
	Name string

	X float64
	Y float64
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (state *UmlState) SerializeToStage() {

	state.Stage()
}
