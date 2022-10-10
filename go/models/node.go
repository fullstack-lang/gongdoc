package models

// Node is a node in the tree for selecting items to display
type Node struct {
	Name string

	IsExpanded bool

	// fields related to the selection of the node by a check box
	HasCheckboxButton bool
	IsChecked         bool

	Children []*Node
}
