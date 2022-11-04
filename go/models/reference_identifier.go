package models

// ReferenceIdentifier store usefull informations about the gongstruct or the gongenum
// refernced by a shape present in the diagrams
type ReferenceIdentifier struct {
	Name        string // name of the gong struct
	NbInstances int
}
