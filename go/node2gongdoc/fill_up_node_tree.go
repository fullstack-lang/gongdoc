package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpNodeTree(gongdocStage *gongdoc_models.StageStruct, diagramPackage *gongdoc_models.DiagramPackage) {

	//
	// LEGACY
	//

	// a node tree is agnostic of the node types it manages
	// therefore, a callback functiion is necessary
	nodeCb := new(NodeCB)
	nodeCb.diagramPackage = diagramPackage

	treeOfGongObjects := FillUpTreeOfGongObjects(gongdocStage)

	//
	// NEW IMPLEMENTATION
	//

	// create a tree for classdiagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc"}).Stage(gongdocStage)

	// add the root of class diagrams
	rootOfClassdiagramsNode := (&gongdoc_models.Node{Name: "Class diagrams"}).Stage(gongdocStage)
	rootOfClassdiagramsNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, rootOfClassdiagramsNode)

	// add add button
	addButton := (&gongdoc_models.Button{
		Name:      diagramPackage.Name + " " + string(BUTTON_add),
		Icon:      string(BUTTON_add),
		Displayed: true}).Stage(gongdocStage)
	_ = addButton
	rootOfClassdiagramsNode.Buttons = append(rootOfClassdiagramsNode.Buttons, addButton)
	addButton.Impl = NewButtonImplRootOfClassdiagrams(
		diagramPackage,
		rootOfClassdiagramsNode,
		treeOfGongObjects,
		BUTTON_add,
	)

	// add one node for each diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
		classdiagramNode := NewClassdiagramNode(classdiagram, diagramPackage, rootOfClassdiagramsNode, treeOfGongObjects)
		rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)
	}

	computeNodeConfs(gongdocStage,
		rootOfClassdiagramsNode,
		diagramPackage,
		treeOfGongObjects)
	// nodeCb.computeNodesConfiguration(gongdocStage)

	// set callbacks on node updates
	// gongdocStage.OnAfterNodeUpdateCallback = nodeCb
	// gongdocStage.OnAfterNodeCreateCallback = nodeCb
	// gongdocStage.OnAfterNodeDeleteCallback = nodeCb
}
