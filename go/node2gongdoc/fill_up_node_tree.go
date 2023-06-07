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

		nodeImplClassdiagram := NewNodeImplClasssiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			treeOfGongObjects,
		)
		classdiagramNode.Impl2 = nodeImplClassdiagram

		//
		// Buttons that are displayed if the node is selected
		//

		// add draw button
		drawButton := (&gongdoc_models.Button{
			Name:      classdiagram.Name + " " + string(BUTTON_draw),
			Icon:      string(BUTTON_draw),
			Displayed: true}).Stage(diagramPackage.Stage_)
		_ = drawButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, drawButton)
		drawButton.Impl = NewButtonImplClassdiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
			BUTTON_draw,
		)

		// delete button
		deleteButton := (&gongdoc_models.Button{
			Name:      classdiagram.Name + " " + string(BUTTON_delete),
			Icon:      string(BUTTON_delete),
			Displayed: true}).Stage(diagramPackage.Stage_)
		_ = deleteButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, deleteButton)
		deleteButton.Impl = NewButtonImplClassdiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
			BUTTON_delete,
		)

		// file_copy
		// delete button
		file_copyButton := (&gongdoc_models.Button{
			Name:      classdiagram.Name + " " + string(BUTTON_file_copy),
			Icon:      string(BUTTON_file_copy),
			Displayed: true}).Stage(diagramPackage.Stage_)
		_ = file_copyButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, file_copyButton)
		file_copyButton.Impl = NewButtonImplClassdiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
			BUTTON_file_copy,
		)
		//
		// Buttons that are displayed if the node is drawned
		//

		// add editoff button
		editoffButton := (&gongdoc_models.Button{
			Name:      classdiagram.Name + " " + string(BUTTON_edit_off),
			Icon:      string(BUTTON_edit_off),
			Displayed: true}).Stage(diagramPackage.Stage_)
		_ = editoffButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, editoffButton)
		editoffButton.Impl = NewButtonImplClassdiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
			BUTTON_edit_off,
		)

		// add save button
		saveButton := (&gongdoc_models.Button{
			Name:      classdiagram.Name + " " + string(BUTTON_save),
			Icon:      string(BUTTON_save),
			Displayed: true}).Stage(diagramPackage.Stage_)
		_ = saveButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, saveButton)
		saveButton.Impl = NewButtonImplClassdiagram(
			diagramPackage,
			classdiagram,
			rootOfClassdiagramsNode,
			treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
			BUTTON_save,
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
