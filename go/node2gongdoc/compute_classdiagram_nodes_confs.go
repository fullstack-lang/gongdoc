package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func computeDiagramNodesConfigurations(
	diagramPackageNode *gongdoc_models.Node,
	diagramPackage *gongdoc_models.DiagramPackage,
	gongdocStage *gongdoc_models.StageStruct) {

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool
	for _, classdiagramNode := range diagramPackageNode.Children {
		if classdiagramNode.IsInDrawMode || classdiagramNode.IsInEditMode {
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
		classdiagramNode.HasDrawOffButton = false

		classdiagramNode.IsCheckboxDisabled = inModificationMode

		if !classdiagramNode.IsChecked {
			classdiagramNode.IsInEditMode = false
			classdiagramNode.IsInDrawMode = false
			continue
		}

		// the classdiagram has been checked
		editable := diagramPackage.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode

		classdiagramNode.HasEditButton = editable
		classdiagramNode.HasDeleteButton = editable
		classdiagramNode.HasDrawButton = editable
		SetButtonDiaplayState(classdiagramNode, BUTTON_draw, editable)
		classdiagramNode.HasDuplicateButton = editable
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
