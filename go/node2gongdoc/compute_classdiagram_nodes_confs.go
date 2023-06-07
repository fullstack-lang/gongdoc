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

		nodeImplClasssiagram, ok := classdiagramNode.Impl2.(*NodeImplClasssiagram)
		if !ok {
			log.Fatalln("not a good interface")
		}
		if nodeImplClasssiagram.IsInDrawMode {
			inModificationMode = true
		}
	}

	diagramPackageNode.HasAddChildButton = !inModificationMode && diagramPackage.IsEditable

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range diagramPackageNode.Children {

		// reset the state of the classdiagram node
		classdiagramNode.HasEditButton = false
		classdiagramNode.HasDeleteButton = false
		classdiagramNode.HasDuplicateButton = false
		classdiagramNode.HasDrawButton = false

		SetButtonDiaplayState(classdiagramNode, BUTTON_draw, false)
		SetButtonDiaplayState(classdiagramNode, BUTTON_edit_off, false)
		SetButtonDiaplayState(classdiagramNode, BUTTON_save, false)

		nodeImplClasssiagram, ok := classdiagramNode.Impl2.(*NodeImplClasssiagram)
		if !ok {
			log.Fatalln("not a good interface")
		}

		classdiagramNode.IsCheckboxDisabled = inModificationMode

		if !classdiagramNode.IsChecked {
			continue
		}

		displayDrawButton := diagramPackage.IsEditable && !nodeImplClasssiagram.IsInDrawMode
		displayEditOffButton := diagramPackage.IsEditable && nodeImplClasssiagram.IsInDrawMode

		SetButtonDiaplayState(classdiagramNode, BUTTON_draw, displayDrawButton)
		SetButtonDiaplayState(classdiagramNode, BUTTON_edit_off, displayEditOffButton)
		SetButtonDiaplayState(classdiagramNode, BUTTON_save, displayEditOffButton)

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
