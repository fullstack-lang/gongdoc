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
		classshapeNode := nodesCb.map_Identifier_Node[ShapenameToIdentifier(gongStruct.Name)]
		classshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	//
	// Enum
	//
	// 1. for each enum of the model, enable the check button if the class diagram is in draw mode
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
		classshapeNode := nodesCb.map_Identifier_Node[ShapenameToIdentifier(gongEnum.Name)]
		classshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	// FIRST STAGE : for each identifiers node with a related visual element is in
	// the classdiagram: check the node and enable it if the diagram is in drawMode

	// 1. gongstructs / gongenum referenced by the classshape
	for _, classshape := range classdiagram.Classshapes {

		var classshapeNode *Node
		var ok bool
		classshapeNode, ok = nodesCb.map_Identifier_Node[classshape.Identifier]

		if !ok {
			log.Println("Unknown node, diagram not synchro with model ", classshape.Identifier)
			continue
		}
		classshapeNode.IsChecked = true

		// disable checkbox of all children of the gongstruct
		for _, fieldOfClassshapeNode := range classshapeNode.Children {
			fieldOfClassshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
		}

		for _, field := range classshape.Fields {
			classshapeFieldNode, ok := nodesCb.map_Identifier_Node[field.Identifier]

			if !ok {
				log.Fatalln(field.Identifier, "unknown node")
			}

			classshapeFieldNode.IsChecked = true
			classshapeFieldNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		}
		for _, link := range classshape.Links {
			linkNode, ok := nodesCb.map_Identifier_Node[link.Identifier]

			if !ok {
				log.Fatalln(link.Identifier, "unknown node")
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
		classshapeNode := nodesCb.map_Identifier_Node[ShapenameToIdentifier(gongStruct.Name)]

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
		noteshapeNode := nodesCb.map_Identifier_Node[ShapenameToIdentifier(note.Name)]
		noteshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		for _, noteLink := range note.Links {

			noteLinkNode := nodesCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(note.Name, noteLink.Name)]

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

	// 2. for each noteShape of the diagram, set the check button of the related node to true
	for _, noteshape := range classdiagram.NoteShapes {
		noteshapeNode := nodesCb.map_Identifier_Node[noteshape.Identifier]

		noteshapeNode.IsChecked = true

		// 2.1 for each link of the note in the diagram, set the check button to true
		for _, noteShapeLink := range noteshape.NoteShapeLinks {
			noteShapeLinkNode := nodesCb.map_Identifier_Node[noteShapeLink.Identifier]
			noteShapeLinkNode.IsChecked = true
		}
	}

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
