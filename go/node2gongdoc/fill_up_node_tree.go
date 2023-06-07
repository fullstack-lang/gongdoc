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
	treeOfGongObjects := nodeCb.FillUpTreeOfGongObjects()

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

		nodeImplClassdiagram := NewNodeImplClasssiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			legacyRootOfClassdiagramsNode,
			nodeCb.treeOfGongObjects,
		)
		classdiagramNode.Impl2 = nodeImplClassdiagram

		// add draw button
		drawButton := (&gongdoc_models.Button{
			Name:      classdiagram.Name + " " + string(BUTTON_draw),
			Icon:      string(BUTTON_draw),
			Displayed: true}).Stage(diagramPackage.Stage_)
		_ = drawButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, drawButton)
		drawButton.Impl = NewButtonImplClassdiagramDraw(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			legacyRootOfClassdiagramsNode,
			nodeCb.treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
		)

		// add editoff button
		editoffButton := (&gongdoc_models.Button{
			Name:      classdiagram.Name + " " + string(BUTTON_edit_off),
			Icon:      string(BUTTON_edit_off),
			Displayed: true}).Stage(diagramPackage.Stage_)
		_ = editoffButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, editoffButton)
		editoffButton.Impl = NewButtonImplClassdiagramEditOff(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			legacyRootOfClassdiagramsNode,
			nodeCb.treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
		)
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
