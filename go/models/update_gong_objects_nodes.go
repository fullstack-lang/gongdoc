package models

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

func updateGongObjectsNodes(stage *StageStruct, nodeCb *NodeCB, classdiagram *Classdiagram) {

	// parse gong object nodes and disable the check box
	for _, rootNode := range nodeCb.treeOfGongObjects.RootNodes {
		for _, node := range rootNode.Children {
			node.IsCheckboxDisabled = !classdiagram.IsInDrawMode
			for _, node := range node.Children {
				node.IsCheckboxDisabled = !classdiagram.IsInDrawMode
			}
		}
	}

	// FIRST STAGE : for each identifiers node with a related visual element is in
	// the classdiagram: check the node and enable it if the diagram is in drawMode

	// 1. gongstructs / gongenum referenced by the classshape
	for _, classshape := range classdiagram.Classshapes {

		var classshapeNode *Node
		var ok bool
		classshapeNode, ok = nodeCb.map_Identifier_Node[classshape.Identifier]

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
			classshapeFieldNode, ok := nodeCb.map_Identifier_Node[field.Identifier]

			if !ok {
				log.Fatalln(field.Identifier, "unknown node")
			}

			classshapeFieldNode.IsChecked = true
			classshapeFieldNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		}
		for _, link := range classshape.Links {
			linkNode, ok := nodeCb.map_Identifier_Node[link.Identifier]

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
		classshapeNode := nodeCb.map_Identifier_Node[ShapenameToIdentifier(gongStruct.Name)]

		for _, fieldOfClassshapeNode := range classshapeNode.Children {
			// then disable the checkbox
			if fieldOfClassshapeNode.Type == GONG_STRUCT_FIELD {

				fieldImpl := fieldOfClassshapeNode.impl.(*FieldImpl)

				switch fieldWithRef := fieldImpl.field.(type) {
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
		noteshapeNode := nodeCb.map_Identifier_Node[ShapenameToIdentifier(note.Name)]
		noteshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		for _, noteLink := range note.Links {

			noteLinkNode := nodeCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(note.Name, noteLink.Name)]

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
		noteshapeNode, ok := nodeCb.map_Identifier_Node[noteshape.Identifier]

		if !ok {
			log.Println("Unknown identifier", noteshape.Identifier)
		}

		noteshapeNode.IsChecked = true

		// 2.1 for each link of the note in the diagram, set the check button to true
		for _, noteShapeLink := range noteshape.NoteShapeLinks {
			noteShapeLinkNode := nodeCb.map_Identifier_Node[noteShapeLink.Identifier]
			noteShapeLinkNode.IsChecked = true
		}
	}
}
