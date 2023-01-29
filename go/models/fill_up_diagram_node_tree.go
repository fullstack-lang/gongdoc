package models

func FillUpDiagramNodeTree(diagramPackage *DiagramPackage, nodeCb *NodeCB) {

	// generate tree of diagrams
	gongdocTree := (&Tree{Name: "gongdoc"}).Stage()

	// add the root of class diagrams
	diagramPackageNode := (&Node{Name: "class diagrams", Type: ROOT_OF_CLASS_DIAGRAMS}).Stage()
	diagramPackageNode.IsExpanded = true
	diagramPackageNode.HasAddChildButton = diagramPackage.IsEditable
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, diagramPackageNode)

	// add one node per class diagram
	for classdiagram := range *GetGongstructInstancesSet[Classdiagram]() {
		node := (&Node{Name: classdiagram.Name}).Stage()
		node.Classdiagram = classdiagram
		node.Type = CLASS_DIAGRAM
		node.HasCheckboxButton = true
		node.HasDeleteButton = true
		node.HasEditButton = true
		diagramPackageNode.Children = append(diagramPackageNode.Children, node)

		// set up the back pointer from the shape to the node
		classdiagramImpl := new(ClassdiagramImpl)
		classdiagramImpl.node = node
		classdiagramImpl.classdiagram = classdiagram
		node.impl = classdiagramImpl
	}

	// set callbacks on node updates
	nodeCb.diagramPackageNode = diagramPackageNode
}
