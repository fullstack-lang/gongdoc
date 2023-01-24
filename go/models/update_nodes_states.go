package models

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// updateNodesStates updates the tree of symbols
// according to the selected diagram
//
// ## The algorithm is as follow
//
// ## For the diagram nodes
//
// For the identifiers nodes
func updateNodesStates(stage *StageStruct, nodesCb *NodeCB) {

	nodesCb.idTree.UncheckAndDisable()

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range nodesCb.ClassdiagramsRootNode.Children {

		classdiagramNode.HasEditButton = false
		classdiagramNode.HasDeleteButton = false
		classdiagramNode.HasDrawButton = false
		classdiagramNode.HasDrawOffButton = false

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

	// get the diagram
	classdiagram := nodesCb.selectedClassdiagram

	if classdiagram == nil {
		Stage.Commit()
		return
	}

	// 1. allow all gongstructs node to be checked/unckecked
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
		nodesCb.map_Identifier_Node[gongStruct.Name].IsCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	//
	// Enum
	//
	// 1. for each enum of the model, enable the check button if the class diagram is in draw mode
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
		nodesCb.map_Identifier_Node[gongEnum.Name].IsCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	// 2. get the all classshape and check the node of the
	// gongstructs / gongenum referenced by the classshape
	for _, classshape := range classdiagram.Classshapes {

		classshapeName := IdentifierToShapename(classshape.Identifier)
		var node *Node
		var ok bool
		node, ok = nodesCb.map_Identifier_Node[classshapeName]

		if !ok {
			log.Println("Unknown node, diagram not synchro with model ", classshapeName)
			continue
		}
		node.IsChecked = true

		// disable checkbox of all children of the gongstruct
		for _, gongfieldNode := range node.Children {
			gongfieldNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
		}

		for _, field := range classshape.Fields {
			nodeId := classshapeName + "." + IdentifierToFieldName(field.Identifier)
			// node, ok := map_FieldId_Node[nodeId]
			node, ok := nodesCb.map_Identifier_Node[nodeId]

			if !ok {
				log.Fatalln(nodeId, "unknown node")
			}

			node.IsChecked = true
			node.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		}
		for _, link := range classshape.Links {
			nodeId := classshapeName + "." + IdentifierToFieldName(link.Identifier)
			node, ok := nodesCb.map_Identifier_Node[nodeId]

			if !ok {
				log.Fatalln(nodeId, "unknown node")
			}

			node.IsChecked = true
			node.IsCheckboxDisabled = !classdiagram.IsInDrawMode
		}
	}

	// disable checkbox field node that reference a gongstruct whose classshape is
	// not present in the diagram
	// first, construct map of all gongstructs present in the diagram
	set_of_classshape_names := make(map[string]bool)
	for _, classshape := range classdiagram.Classshapes {
		set_of_classshape_names[IdentifierToShapename(classshape.Identifier)] = true
	}
	// then iterate over all fields of all gongstructs node
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
		gongstructNode := nodesCb.map_Identifier_Node[gongStruct.Name]

		for _, fieldNode := range gongstructNode.Children {
			// then disable the checkbox
			if fieldNode.Type == GONG_STRUCT_FIELD {
				switch fieldWithRef := fieldNode.Gongfield.(type) {
				case *gong_models.PointerToGongStructField:
					if _, ok := set_of_classshape_names[fieldWithRef.GongStruct.Name]; !ok {
						fieldNode.IsCheckboxDisabled = true
					}
				case *gong_models.SliceOfPointerToGongStructField:
					if _, ok := set_of_classshape_names[fieldWithRef.GongStruct.Name]; !ok {
						fieldNode.IsCheckboxDisabled = true
					}
				}
			}
		}
	}

	//
	// For notes
	//
	// 1. for each note of the model, enable the check button if the class diagram is in draw mode
	set_of_noteshape_names := make(map[string]bool)
	for _, noteshape := range classdiagram.NoteShapes {
		set_of_noteshape_names[IdentifierToShapename(noteshape.Identifier)] = true
	}
	for note := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {
		nodesCb.map_Identifier_Node[note.Name].IsCheckboxDisabled = !classdiagram.IsInDrawMode

		for _, noteLink := range note.Links {

			noteLinkName := note.Name + "." + noteLink.Name
			node := nodesCb.map_Identifier_Node[noteLinkName]

			// default choice
			node.IsCheckboxDisabled = !classdiagram.IsInDrawMode

			// disable check box of the note link if the note is not present
			if _, ok := set_of_noteshape_names[note.Name]; !ok {
				node.IsCheckboxDisabled = true
			}
		}
	}
	// 2. for each note of the diagram, set the check button to true
	for _, note := range classdiagram.NoteShapes {
		nodesCb.map_Identifier_Node[IdentifierToShapename(note.Identifier)].IsChecked = true
	}

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
