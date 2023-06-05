package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpNodeTree(diagramPackage *gongdoc_models.DiagramPackage) {

	//
	// NEW IMPLEMENTATION
	//

	// create a tree for classdiagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc_new"}).Stage(diagramPackage.Stage_)

	// add the root of class diagrams
	rootOfClassdiagrams := (&gongdoc_models.Node{Name: "class diagrams 2"}).Stage(diagramPackage.Stage_)
	rootOfClassdiagrams.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, rootOfClassdiagrams)

	// append an implementation
	rootOfClassdiagrams.Impl2 = NewNodeImplRootOfClassDiagrams()

	//
	// LEGACY
	//

	// a node tree is agnostic of the node types it manages
	// therefore, a callback functiion is necessary
	nodeCb := new(NodeCB)
	nodeCb.diagramPackage = diagramPackage

	nodeCb.FillUpDiagramNodeTree(diagramPackage)
	nodeCb.FillUpTreeOfGongObjects()
	nodeCb.computeNodesConfiguration(diagramPackage.Stage_)

	// set callbacks on node updates
	diagramPackage.Stage_.OnAfterNodeUpdateCallback = nodeCb
	diagramPackage.Stage_.OnAfterNodeCreateCallback = nodeCb
	diagramPackage.Stage_.OnAfterNodeDeleteCallback = nodeCb
}
