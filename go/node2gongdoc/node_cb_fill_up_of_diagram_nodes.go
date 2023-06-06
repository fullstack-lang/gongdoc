package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func (nodeCb *NodeCB) FillUpDiagramNodeTree(
	diagramPackage *gongdoc_models.DiagramPackage,
) (gongdocTree *gongdoc_models.Tree) {

	// generate tree of diagrams
	gongdocTree = (&gongdoc_models.Tree{Name: "gongdoc"}).Stage(nodeCb.diagramPackage.Stage_)

	// add the root of class diagrams
	diagramPackageNode := (&gongdoc_models.Node{Name: "class diagrams"}).Stage(nodeCb.diagramPackage.Stage_)
	diagramPackageNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, diagramPackageNode)

	// add one node per class diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](nodeCb.diagramPackage.Stage_) {
		node := (&gongdoc_models.Node{Name: classdiagram.Name}).Stage(nodeCb.diagramPackage.Stage_)

		node.HasCheckboxButton = true
		node.HasDeleteButton = true
		node.HasEditButton = true
		node.HasDuplicateButton = true
		diagramPackageNode.Children = append(diagramPackageNode.Children, node)

		// set up the back pointer from the shape to the node
		classdiagramImpl := new(ClassdiagramImpl)
		classdiagramImpl.node = node
		classdiagramImpl.classdiagram = classdiagram
		classdiagramImpl.nodeCb = nodeCb
		node.Impl = classdiagramImpl
	}

	// set callbacks on node updates
	nodeCb.diagramPackageNode = diagramPackageNode

	return
}
