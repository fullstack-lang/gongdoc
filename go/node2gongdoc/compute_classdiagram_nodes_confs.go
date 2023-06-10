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

		// remove all buttons
		for _, _button := range classdiagramNode.Buttons {
			_button.Unstage(gongdocStage)
		}
		classdiagramNode.Buttons = make([]*gongdoc_models.Button, 0)

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

		SetButton(classdiagramNode, diagramPackageNode, BUTTON_draw, selectedForEdit)
		SetButton(classdiagramNode, diagramPackageNode, BUTTON_delete, selectedForEdit)
		SetButton(classdiagramNode, diagramPackageNode, BUTTON_file_copy, selectedForEdit)
		SetButton(classdiagramNode, diagramPackageNode, BUTTON_edit, selectedForEdit)

		SetButton(classdiagramNode, diagramPackageNode, BUTTON_edit_off, inDrawingMode || inEditMode)
		SetButton(classdiagramNode, diagramPackageNode, BUTTON_save, inDrawingMode)

	}
}

// SetButton set the display attribute of the button designed by buttonId of the node
func SetButton(
	classdiagramNode *gongdoc_models.Node,
	diagramPackageNode *gongdoc_models.Node,
	icon ButtonType, displayed bool) {

	nodeImplClassdiagram, ok := classdiagramNode.Impl.(*NodeImplClasssiagram)
	if !ok {
		log.Fatalln("youre kidding me ?")
	}

	if displayed == true {
		drawButton := (&gongdoc_models.Button{
			Name: nodeImplClassdiagram.classdiagram.Name + " " + string(icon),
			Icon: string(icon)}).Stage(nodeImplClassdiagram.diagramPackage.Stage_)
		_ = drawButton
		classdiagramNode.Buttons = append(classdiagramNode.Buttons, drawButton)
		drawButton.Impl = NewButtonImplClassdiagram(
			nodeImplClassdiagram.diagramPackage,
			nodeImplClassdiagram.classdiagram,
			diagramPackageNode,
			nodeImplClassdiagram.treeOfGongObjects,
			classdiagramNode,
			nodeImplClassdiagram,
			icon,
		)
	}

	if displayed == false {
		return
		// for _, _button := range classdiagramNode.Buttons {
		// 	if _button.Icon == string(icon) {
		// 		remove[gongdoc_models.Button](classdiagramNode.Buttons, _button)
		// 		_button.Unstage(nodeImplClassdiagram.diagramPackage.Stage_)
		// 		continue
		// 	}
		// }
	}
}
