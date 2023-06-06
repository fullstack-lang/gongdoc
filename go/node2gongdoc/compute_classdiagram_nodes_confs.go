package node2gongdoc

import (
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
		classdiagramNode.HasDuplicateButton = editable
	}
}
