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

	leagacyTreeOfDiagramNodes := nodeCb.FillUpDiagramNodeTree(diagramPackage)
	nodeCb.FillUpTreeOfGongObjects()
	nodeCb.computeNodesConfiguration(diagramPackage.Stage_)

	//
	// NEW IMPLEMENTATION
	//

	// create a tree for classdiagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc_new"}).Stage(diagramPackage.Stage_)

	// add the root of class diagrams
	rootOfClassdiagramsNode := (&gongdoc_models.Node{Name: "class diagrams 2"}).Stage(diagramPackage.Stage_)
	rootOfClassdiagramsNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, rootOfClassdiagramsNode)

	// append an implementation
	rootOfClassdiagramsNode.Impl2 = NewNodeImplRootOfClassDiagrams()

	// add one node for each diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](diagramPackage.Stage_) {
		classdiagramNode := (&gongdoc_models.Node{Name: classdiagram.Name}).Stage(diagramPackage.Stage_)

		rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)

		classdiagramNode.HasCheckboxButton = true

		// fetch the root of diagram nodes
		var legacyRootOfClassdiagramsNode *gongdoc_models.Node
		for _, _node := range leagacyTreeOfDiagramNodes.RootNodes {
			legacyRootOfClassdiagramsNode = _node
		}

		classdiagramNode.Impl2 = NewNodeImplClasssiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			legacyRootOfClassdiagramsNode,
			nodeCb.treeOfGongObjects)
	}

	// set callbacks on node updates
	// diagramPackage.Stage_.OnAfterNodeUpdateCallback = nodeCb
	// diagramPackage.Stage_.OnAfterNodeCreateCallback = nodeCb
	// diagramPackage.Stage_.OnAfterNodeDeleteCallback = nodeCb
}
