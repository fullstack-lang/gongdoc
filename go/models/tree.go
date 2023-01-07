package models

// Tree is a tree of nodes for selecting items to display
type Tree struct {
	Name string

	Type TreeType

	RootNodes []*Node

	// map to access each according to its id
	map_Identifier_Node map[string]*Node
}
