package models

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// updateNodesStates updates the tree of symbols
// according to the selected diagram
func updateNodesStates(stage *StageStruct, callbacksSingloton *NodeCallbacksSingloton) {

	// get the editable state of the package
	var pkglet *DiagramPackage
	for _pkgelt := range *GetGongstructInstancesSet[DiagramPackage]() {
		pkglet = _pkgelt
	}

	// map of gognstruct nodes according to their name
	map_StructId_Node := make(map[string]*Node)

	// map of gong fields according to their name
	map_FieldId_Node := make(map[string]*Node)

	// map of note links according to their name
	map_NoteLinkId_Node := make(map[string]*Node)

	// unckeck nodes and construct the map
	for _, gognstructNode := range callbacksSingloton.GongstructsRootNode.Children {

		gognstructNode.IsCheckboxDisabled = true

		gognstructNode.IsChecked = false
		map_StructId_Node[gognstructNode.Name] = gognstructNode

		for _, gongfieldNode := range gognstructNode.Children {
			gongfieldNode.IsChecked = false
			gongfieldNode.IsCheckboxDisabled = true
			map_FieldId_Node[gognstructNode.Name+"."+gongfieldNode.Name] = gongfieldNode
		}
	}
	for _, gognenumNode := range callbacksSingloton.GongenumsRootNode.Children {

		gognenumNode.IsCheckboxDisabled = true
		gognenumNode.IsChecked = false
		map_StructId_Node[gognenumNode.Name] = gognenumNode

		for _, gongValueNode := range gognenumNode.Children {
			gongValueNode.IsChecked = false
			gongValueNode.IsCheckboxDisabled = true
			map_FieldId_Node[gognenumNode.Name+"."+gongValueNode.Name] = gongValueNode
		}
	}
	for _, gongnoteNode := range callbacksSingloton.GongnotesRootNode.Children {
		gongnoteNode.IsCheckboxDisabled = true
		gongnoteNode.IsChecked = false
		map_StructId_Node[gongnoteNode.Name] = gongnoteNode

		for _, noteLink := range gongnoteNode.Children {
			map_NoteLinkId_Node[gongnoteNode.Name+"."+noteLink.Name] = noteLink
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range callbacksSingloton.ClassdiagramsRootNode.Children {

		classdiagramNode.HasEditButton = false
		classdiagramNode.HasDeleteButton = false
		classdiagramNode.HasDrawButton = false
		classdiagramNode.HasDrawOffButton = false

		if !classdiagramNode.IsChecked {
			classdiagramNode.IsInEditMode = false
			classdiagramNode.IsInDrawMode = false
			continue
		}

		classdiagramNode.HasEditButton =
			pkglet.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode
		classdiagramNode.HasDeleteButton =
			pkglet.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode
		classdiagramNode.HasDrawButton =
			pkglet.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode

		// get the diagram
		classDiagram := classdiagramNode.Classdiagram

		// 1. allow all gongstructs node to be checked/unckecked
		for _, gognstructNode := range callbacksSingloton.GongstructsRootNode.Children {
			gognstructNode.IsCheckboxDisabled = !classDiagram.IsInDrawMode
		}

		// 2. get the all classshape and check the node of the
		// referenced gongstructs if it is referenced by the classshape
		for _, classshape := range classDiagram.Classshapes {
			ref_GongStruct := classshape.Reference

			node, ok := map_StructId_Node[ref_GongStruct.Name]

			if !ok {
				log.Println("Unknown node ", ref_GongStruct.Name)
				continue
			}
			node.IsChecked = true

			// disable checkbox of all children of the gongstruct
			for _, gongfieldNode := range map_StructId_Node[ref_GongStruct.Name].Children {
				gongfieldNode.IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}

			for _, field := range classshape.Fields {
				nodeId := ref_GongStruct.Name + "." + field.Fieldname
				node, ok := map_FieldId_Node[nodeId]
				if ok {
					node.IsChecked = true
					node.IsCheckboxDisabled = !classDiagram.IsInDrawMode
				}
			}
			for _, link := range classshape.Links {
				map_FieldId_Node[ref_GongStruct.Name+"."+link.Name].IsChecked = true
				map_FieldId_Node[ref_GongStruct.Name+"."+link.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}
		}

		// disable checkbox field node that reference a gongstruct whose classshape is
		// not present in the diagram
		// first, construct map of all gongstructs present in the diagram
		map_OfGongstruct := make(map[string]bool)
		for _, classshape := range classDiagram.Classshapes {
			map_OfGongstruct[classshape.ReferenceName] = true
		}
		// then iterate over all fields of all gongstructs node
		for _, gognstructNode := range callbacksSingloton.GongstructsRootNode.Children {
			for _, fieldNode := range gognstructNode.Children {
				// then disable the checkbox
				if fieldNode.Type == GONG_STRUCT_FIELD {
					switch fieldWithRef := fieldNode.Gongfield.(type) {
					case *gong_models.PointerToGongStructField:
						if _, ok := map_OfGongstruct[fieldWithRef.GongStruct.Name]; !ok {
							fieldNode.IsCheckboxDisabled = true
						}
					case *gong_models.SliceOfPointerToGongStructField:
						if _, ok := map_OfGongstruct[fieldWithRef.GongStruct.Name]; !ok {
							fieldNode.IsCheckboxDisabled = true
						}
					}
				}
			}
		}

		//
		// For notes
		//
		// 1. for each note of the model, enable the check button is the class diagram is in draw mode
		// 2. for each note of the diagram, set the check button to true
		for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {
			map_StructId_Node[gongNote.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
			for _, docLink := range gongNote.Links {
				map_NoteLinkId_Node[gongNote.Name+"."+docLink.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}
		}

		for _, note := range classDiagram.Notes {
			map_StructId_Node[note.Name].IsChecked = true

			// the following line seems redondant with the loop of all gong notes
			// map_StructId_Node[note.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
		}

		for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
			map_StructId_Node[gongEnum.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
		}
	}

	for _, stateDiagramNode := range callbacksSingloton.StateDiagramsRootNode.Children {
		stateDiagramNode.HasEditButton = pkglet.IsEditable && stateDiagramNode.IsChecked && !stateDiagramNode.IsInEditMode
		stateDiagramNode.HasDeleteButton = pkglet.IsEditable && stateDiagramNode.IsChecked && !stateDiagramNode.IsInEditMode
		stateDiagramNode.HasDrawButton = pkglet.IsEditable && stateDiagramNode.IsChecked && !stateDiagramNode.IsInEditMode
	}

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
