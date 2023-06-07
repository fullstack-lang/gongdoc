package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpNodeTree(diagramPackage *gongdoc_models.DiagramPackage) {

	//
	// LEGACY
	//

	// a node tree is agnostic of the node types it manages
	// therefore, a callback functiion is necessary
	nodeCb := new(NodeCB)
	nodeCb.diagramPackage = diagramPackage

	treeOfGongObjects := nodeCb.FillUpTreeOfGongObjects()

	//
	// NEW IMPLEMENTATION
	//

	// create a tree for classdiagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc"}).Stage(diagramPackage.Stage_)

	// add the root of class diagrams
	rootOfClassdiagramsNode := (&gongdoc_models.Node{Name: "Class diagrams"}).Stage(diagramPackage.Stage_)
	rootOfClassdiagramsNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, rootOfClassdiagramsNode)

	// append an implementation
	rootOfClassdiagramsNode.Impl2 = NewNodeImplRootOfClassDiagrams()

	// add one node for each diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](diagramPackage.Stage_) {
		classdiagramNode := NewClassdiagramNode(classdiagram, diagramPackage, rootOfClassdiagramsNode, treeOfGongObjects)
		rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)
	}

	computeNodeConfs(diagramPackage.Stage_,
		rootOfClassdiagramsNode,
		diagramPackage,
		treeOfGongObjects)
	// nodeCb.computeNodesConfiguration(diagramPackage.Stage_)

	// set callbacks on node updates
	// diagramPackage.Stage_.OnAfterNodeUpdateCallback = nodeCb
	// diagramPackage.Stage_.OnAfterNodeCreateCallback = nodeCb
	// diagramPackage.Stage_.OnAfterNodeDeleteCallback = nodeCb
}
