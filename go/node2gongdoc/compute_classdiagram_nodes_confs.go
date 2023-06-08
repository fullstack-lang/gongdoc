package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func computeClassdiagramNodesConfigurations(
	diagramPackageNode *gongdoc_models.Node,
	diagramPackage *gongdoc_models.DiagramPackage,
	gongdocStage *gongdoc_models.StageStruct) {

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool
	for _, classdiagramNode := range diagramPackageNode.Children {

		nodeImplClasssiagram, ok := classdiagramNode.Impl.(*NodeImplClasssiagram)
		if !ok {
			log.Fatalln("not a good interface")
		}
		if nodeImplClasssiagram.IsInDrawMode {
			inModificationMode = true
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range diagramPackageNode.Children {

		SetButtonDiaplayState(classdiagramNode, BUTTON_draw, false)
		SetButtonDiaplayState(classdiagramNode, BUTTON_edit_off, false)
		SetButtonDiaplayState(classdiagramNode, BUTTON_save, false)
		SetButtonDiaplayState(classdiagramNode, BUTTON_delete, false)
		SetButtonDiaplayState(classdiagramNode, BUTTON_file_copy, false)
		SetButtonDiaplayState(classdiagramNode, BUTTON_edit, false)

		nodeImplClasssiagram, ok := classdiagramNode.Impl.(*NodeImplClasssiagram)
		if !ok {
			log.Fatalln("not a good interface")
		}

		classdiagramNode.IsCheckboxDisabled = inModificationMode

		if !classdiagramNode.IsChecked {
			continue
		}

		selectedForEdit := diagramPackage.IsEditable && !nodeImplClasssiagram.IsInDrawMode && !classdiagramNode.IsInEditMode
		inDrawingMode := diagramPackage.IsEditable && nodeImplClasssiagram.IsInDrawMode
		inEditMode := diagramPackage.IsEditable && classdiagramNode.IsInEditMode

		SetButtonDiaplayState(classdiagramNode, BUTTON_draw, selectedForEdit)
		SetButtonDiaplayState(classdiagramNode, BUTTON_delete, selectedForEdit)
		SetButtonDiaplayState(classdiagramNode, BUTTON_file_copy, selectedForEdit)
		SetButtonDiaplayState(classdiagramNode, BUTTON_edit, selectedForEdit)

		SetButtonDiaplayState(classdiagramNode, BUTTON_edit_off, inDrawingMode || inEditMode)
		SetButtonDiaplayState(classdiagramNode, BUTTON_save, inDrawingMode)

	}
}

// SetButtonDiaplayState set the display attribute of the button designed by buttonId of the node
func SetButtonDiaplayState(node *gongdoc_models.Node, icon ButtonType, displayed bool) {

	var found bool
	for _, _button := range node.Buttons {
		if _button.Icon == string(icon) {
			_button.Displayed = displayed
			found = true
		}
	}
	if !found {
		log.Fatal("No such button", icon)
	}
}

func GetButtonDiaplayState(node *gongdoc_models.Node, icon ButtonType) (displayed bool) {

	var found bool
	for _, _button := range node.Buttons {
		if _button.Icon == string(icon) {
			displayed = _button.Displayed
			found = true
		}
	}
	if !found {
		log.Fatal("No such button", icon)
	}
	return
}
