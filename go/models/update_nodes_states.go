package models

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// updateNodesStates updates the tree of symbols
// according to the selected diagram
func updateNodesStates(stage *StageStruct, callbacksSingloton *NodeCallbacksSingloton) {

	// get the diagram package of interest
	var pkglet *DiagramPackage
	for _pkgelt := range *GetGongstructInstancesSet[DiagramPackage]() {
		pkglet = _pkgelt
	}

	// unckeck nodes and construct the map
	for _, gognstructNode := range callbacksSingloton.GongstructsRootNode.Children {

		gognstructNode.IsCheckboxDisabled = true
		gognstructNode.IsChecked = false

		for _, gongfieldNode := range gognstructNode.Children {
			gongfieldNode.IsChecked = false
			gongfieldNode.IsCheckboxDisabled = true
		}
	}
	for _, gognenumNode := range callbacksSingloton.GongenumsRootNode.Children {

		gognenumNode.IsCheckboxDisabled = true
		gognenumNode.IsChecked = false

		for _, gongValueNode := range gognenumNode.Children {
			gongValueNode.IsChecked = false
			gongValueNode.IsCheckboxDisabled = true
		}
	}
	for _, node := range callbacksSingloton.GongnotesRootNode.Children {
		node.IsCheckboxDisabled = true
		node.IsChecked = false

		for _, node2 := range node.Children {
			_ = node2
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

			node, ok := callbacksSingloton.tree.nodeMap[ref_GongStruct.Name]

			if !ok {
				log.Fatalln("Unknown node ", ref_GongStruct.Name)
			}
			node.IsChecked = true

			// disable checkbox of all children of the gongstruct
			for _, gongfieldNode := range node.Children {
				gongfieldNode.IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}

			for _, field := range classshape.Fields {
				nodeId := ref_GongStruct.Name + "." + field.Fieldname
				// node, ok := map_FieldId_Node[nodeId]
				node, _ := callbacksSingloton.tree.nodeMap[nodeId]

				node.IsChecked = true
				node.IsCheckboxDisabled = !classDiagram.IsInDrawMode

			}
			for _, link := range classshape.Links {
				nodeId := ref_GongStruct.Name + "." + link.Fieldname
				node, _ := callbacksSingloton.tree.nodeMap[nodeId]

				node.IsChecked = true
				node.IsCheckboxDisabled = !classDiagram.IsInDrawMode
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
			callbacksSingloton.tree.nodeMap[gongNote.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode

			for _, docLink := range gongNote.Links {
				callbacksSingloton.tree.nodeMap[gongNote.Name+"."+docLink.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
			}
		}

		for _, note := range classDiagram.Notes {
			callbacksSingloton.tree.nodeMap[note.Name].IsChecked = true
		}

		for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
			callbacksSingloton.tree.nodeMap[gongEnum.Name].IsCheckboxDisabled = !classDiagram.IsInDrawMode
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
