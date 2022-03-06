package models

// swagger:enum Stage
type Stage string

// values for Action Type
const (
	BEFORE Stage = "BEFORE" // iota // Parse the spinosa model (temp)
	AFTER  Stage = "AFTER"
)
