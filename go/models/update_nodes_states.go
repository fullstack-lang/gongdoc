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
		classshapeNode := nodesCb.map_Identifier_Node[gongStruct.Name]
		classshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	//
	// Enum
	//
	// 1. for each enum of the model, enable the check button if the class diagram is in draw mode
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
		classshapeNode := nodesCb.map_Identifier_Node[gongEnum.Name]
		classshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	// FIRST STAGE : for each identifiers node with a related visual element is in
	// the classdiagram: check the node and enable it if the diagram is in drawMode

	// 1. gongstructs / gongenum referenced by the classshape
	for _, classshape := range classdiagram.Classshapes {

		classshapeName := IdentifierToShapename(classshape.Identifier)
		var classshapeNode *Node
		var ok bool
		classshapeNode, ok = nodesCb.map_Identifier_Node[classshapeName]

		if !ok {
			log.Println("Unknown node, diagram not synchro with model ", classshapeName)
			continue
		}
		classshapeNode.IsChecked = true

		// disable checkbox of all children of the gongstruct
		for _, fieldOfClassshapeNode := range classshapeNode.Children {
			fieldOfClassshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
		}

		for _, field := range classshape.Fields {
			nodeId := classshapeName + "." + IdentifierToFieldName(field.Identifier)
			// classshapeFieldNode, ok := map_FieldId_Node[nodeId]
			classshapeFieldNode, ok := nodesCb.map_Identifier_Node[nodeId]

			if !ok {
				log.Fatalln(nodeId, "unknown node")
			}

			classshapeFieldNode.IsChecked = true
			classshapeFieldNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		}
		for _, link := range classshape.Links {
			nodeId := classshapeName + "." + IdentifierToFieldName(link.Identifier)
			linkNode, ok := nodesCb.map_Identifier_Node[nodeId]

			if !ok {
				log.Fatalln(nodeId, "unknown node")
			}

			linkNode.IsChecked = true
			linkNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
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
		classshapeNode := nodesCb.map_Identifier_Node[gongStruct.Name]

		for _, fieldOfClassshapeNode := range classshapeNode.Children {
			// then disable the checkbox
			if fieldOfClassshapeNode.Type == GONG_STRUCT_FIELD {
				switch fieldWithRef := fieldOfClassshapeNode.Gongfield.(type) {
				case *gong_models.PointerToGongStructField:
					if _, ok := set_of_classshape_names[fieldWithRef.GongStruct.Name]; !ok {
						fieldOfClassshapeNode.IsCheckboxDisabled = true
					}
				case *gong_models.SliceOfPointerToGongStructField:
					if _, ok := set_of_classshape_names[fieldWithRef.GongStruct.Name]; !ok {
						fieldOfClassshapeNode.IsCheckboxDisabled = true
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
		noteshapeNode := nodesCb.map_Identifier_Node[note.Name]
		noteshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		for _, noteLink := range note.Links {

			noteLinkName := note.Name + "." + noteLink.Name
			noteLinkNode := nodesCb.map_Identifier_Node[noteLinkName]

			// default choice
			noteLinkNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

			// disable check box of the note link if the note is not present
			if _, ok := set_of_noteshape_names[note.Name]; !ok {
				noteLinkNode.IsCheckboxDisabled = true
			}

			// disable check box of the note link if the target shape is not present
			if _, ok := set_of_classshape_names[noteLink.Name]; !ok {
				noteLinkNode.IsCheckboxDisabled = true
			}

		}
	}
	// 2. for each note of the diagram, set the check button to true
	for _, note := range classdiagram.NoteShapes {
		noteshapeNode := nodesCb.map_Identifier_Node[IdentifierToShapename(note.Identifier)]

		noteshapeNode.IsChecked = true
	}

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
