package node2gongdoc

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func updateGongObjectsNodes(stage *gongdoc_models.StageStruct, nodeCb *NodeCB, classdiagram *gongdoc_models.Classdiagram) {

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
	for _, gongStructShape := range classdiagram.GongStructShapes {

		var gongStructShapeNode *gongdoc_models.Node
		var ok bool
		gongStructShapeNode, ok = nodeCb.map_Identifier_Node[gongStructShape.Identifier]

		if !ok {
			log.Println("Unknown node, diagram not synchro with model ", gongStructShape.Identifier)
			continue
		}

		// get the gong object from the identifier
		gongStructName := gongdoc_models.IdentifierToShapename(gongStructShape.Identifier)
		gongStruct, ok :=
			(*gong_models.GetGongstructInstancesMap[gong_models.GongStruct]())[gongStructName]

		if !ok {
			log.Println("updateGongObjectsNodes: Unknown gong struct related to identifier:", gongStructShape.Identifier)
			continue
		}
		gongStructImplIF, ok := nodeCb.map_gongObject_gongObjectImpl[gongStruct]
		if !ok {
			log.Println("updateGongObjectsNodes: Unknown gong struct impl related to gong struct:", gongStruct.Name)
			continue
		}
		// set up the pointer to the shape
		gongStructImpl, ok := gongStructImplIF.(*GongStructImpl)

		gongStructImpl2 := GetNodeBackPointer(gongStruct)
		_ = gongStructImpl2

		gongStructImpl.HasDiagramElt = true

		gongStructShapeNode.IsChecked = true

		// disable checkbox of all children of the gongstruct
		for _, fieldOfClassshapeNode := range gongStructShapeNode.Children {
			fieldOfClassshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
		}

		for _, field := range gongStructShape.Fields {
			classshapeFieldNode, ok := nodeCb.map_Identifier_Node[field.Identifier]

			if !ok {
				log.Fatalln(field.Identifier, "unknown node")
			}

			classshapeFieldNode.IsChecked = true
			classshapeFieldNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		}
		for _, link := range gongStructShape.Links {
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
	for _, classshape := range classdiagram.GongStructShapes {
		set_of_classshape_names[gongdoc_models.IdentifierToShapename(classshape.Identifier)] = true
	}

	// then iterate over all fields of all gongstructs node
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
		classshapeNode := nodeCb.map_Identifier_Node[gongdoc_models.ShapenameToIdentifier(gongStruct.Name)]

		for _, fieldOfClassshapeNode := range classshapeNode.Children {
			// then disable the checkbox
			if fieldOfClassshapeNode.Type == gongdoc_models.GONG_STRUCT_FIELD {

				fieldImpl := fieldOfClassshapeNode.Impl.(*FieldImpl)

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
		set_of_noteshape_names[gongdoc_models.IdentifierToShapename(noteshape.Identifier)] = true
	}
	for note := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {
		noteshapeNode := nodeCb.map_Identifier_Node[gongdoc_models.ShapenameToIdentifier(note.Name)]
		noteshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		for _, noteLink := range note.Links {

			noteLinkNode := nodeCb.map_Identifier_Node[gongdoc_models.ShapeAndFieldnameToFieldIdentifier(note.Name, noteLink.Name)]

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
