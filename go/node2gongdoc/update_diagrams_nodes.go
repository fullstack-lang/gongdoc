package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func updateDiagramsNodes(stage *gongdoc_models.StageStruct, nodesCb *NodeCB) {

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool
	for _, classdiagramNode := range nodesCb.diagramPackageNode.Children {
		if classdiagramNode.IsInDrawMode || classdiagramNode.IsInEditMode {
			inModificationMode = true
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range nodesCb.diagramPackageNode.Children {

		classdiagramNode.HasEditButton = false
		classdiagramNode.HasDeleteButton = false
		classdiagramNode.HasDrawButton = false
		classdiagramNode.HasDrawOffButton = false

		classdiagramNode.IsCheckboxDisabled = inModificationMode

		if !classdiagramNode.IsChecked {
			classdiagramNode.IsInEditMode = false
			classdiagramNode.IsInDrawMode = false
			continue
		}

		editable := nodesCb.diagramPackage.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode

		classdiagramNode.HasEditButton = editable
		classdiagramNode.HasDeleteButton = editable
		classdiagramNode.HasDrawButton = editable

	}
}
