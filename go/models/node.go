package models

// Node is a node in the tree for selecting items to display
type Node struct {
	Name string

	Type GongdocNodeType

	// impl is the pointer to the implementation of the node in the models of interest
	impl NodeImplInterface

	IsExpanded bool

	// fields related to the selection of the node by a check box
	HasCheckboxButton  bool
	IsChecked          bool
	IsCheckboxDisabled bool

	HasAddChildButton bool

	// to edit the name
	HasEditButton bool
	IsInEditMode  bool

	// to edit the diagram
	HasDrawButton    bool
	HasDrawOffButton bool
	IsInDrawMode     bool
	IsSaved          bool

	HasDeleteButton bool

	Children []*Node
}

func (node *Node) UncheckAndDisableBranch() {

	node.UncheckAndDisable()

	for _, _node := range node.Children {
		_node.UncheckAndDisableBranch()
	}
}

func (node *Node) UncheckAndDisable() {

	node.IsCheckboxDisabled = true
	node.IsChecked = false
}
