package models

// swagger:enum SimulationStage
type SimulationStage string

// values for Action Type
const (
	BEFORE SimulationStage = "BEFORE" // iota // Parse the spinosa model (temp)
	AFTER  SimulationStage = "AFTER"
)
