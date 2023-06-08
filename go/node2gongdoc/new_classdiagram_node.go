package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func NewClassdiagramNode(
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackage *gongdoc_models.DiagramPackage,
	rootOfClassdiagramsNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
) (classdiagramNode *gongdoc_models.Node) {
	classdiagramNode = (&gongdoc_models.Node{Name: classdiagram.Name}).Stage(diagramPackage.Stage_)

	rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)

	classdiagramNode.HasCheckboxButton = true

	nodeImplClassdiagram := NewNodeImplClasssiagram(
		diagramPackage,
		classdiagram,
		rootOfClassdiagramsNode,
		treeOfGongObjects,
	)
	classdiagramNode.Impl = nodeImplClassdiagram

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

	// add edit button
	editButton := (&gongdoc_models.Button{
		Name:      classdiagram.Name + " " + string(BUTTON_edit),
		Icon:      string(BUTTON_edit),
		Displayed: true}).Stage(diagramPackage.Stage_)
	_ = editButton
	classdiagramNode.Buttons = append(classdiagramNode.Buttons, editButton)
	editButton.Impl = NewButtonImplClassdiagram(
		diagramPackage,
		classdiagram,
		rootOfClassdiagramsNode,
		treeOfGongObjects,
		classdiagramNode,
		nodeImplClassdiagram,
		BUTTON_edit,
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

	return
}
