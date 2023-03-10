package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// computeGongNodesConfiguration set up all gong nodes according to the classdiagram
func (nodeCb *NodeCB) computeGongNodesConfiguration(stage *gongdoc_models.StageStruct, classdiagram *gongdoc_models.Classdiagram) {

	//
	// compute maps of displayed gong objects
	//
	namesOfDisplayedGongstructs := make(map[string]bool)
	for _, gongStructShape := range classdiagram.GongStructShapes {
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		namesOfDisplayedGongstructs[gongStructName] = true
	}

	namesOfDisplayedGongenums := make(map[string]bool)
	for _, gongEnumShape := range classdiagram.GongEnumShapes {
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier)
		namesOfDisplayedGongenums[gongStructName] = true
	}

	namesOfDisplayedGongfields := make(map[string]bool)
	for _, gongStructShape := range classdiagram.GongStructShapes {
		for _, gongFieldShape := range gongStructShape.Links {
			gongFieldName := gongdoc_models.IdentifierToGongObjectName(gongFieldShape.Identifier)
			namesOfDisplayedGongfields[gongFieldName] = true
		}
	}

	namesOfDisplayedGongnotes := make(map[string]bool)
	for _, noteShape := range classdiagram.NoteShapes {
		gongNoteName := gongdoc_models.IdentifierToGongObjectName(noteShape.Identifier)
		namesOfDisplayedGongnotes[gongNoteName] = true
	}

	// disable nodes of fields objects when
	// - the DSL is not displayed
	// - the gongstruct type of the field is not displayed
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		isPresent := namesOfDisplayedGongstructs[gongStruct.Name]
		_ = isPresent

		for _, gongField := range gongStruct.Fields {
			switch fieldReal := gongField.(type) {
			case *gong_models.PointerToGongStructField:
				impl := GetNodeBackPointer(fieldReal)
				if !isPresent {
					impl.DisableNodeCheckbox()
				}
				if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
					impl.DisableNodeCheckbox()
				}
			case *gong_models.SliceOfPointerToGongStructField:
				impl := GetNodeBackPointer(fieldReal)
				if !isPresent {
					impl.DisableNodeCheckbox()
				}
				if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
					impl.DisableNodeCheckbox()
				}
			case *gong_models.GongBasicField:
				impl := GetNodeBackPointer(fieldReal)
				if !isPresent {
					impl.DisableNodeCheckbox()
				}
			case *gong_models.GongTimeField:
				impl := GetNodeBackPointer(fieldReal)
				if !isPresent {
					impl.DisableNodeCheckbox()
				}
			default:

			}
		}
	}

	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
		isPresent := namesOfDisplayedGongenums[gongEnum.Name]
		_ = isPresent

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			impl := GetNodeBackPointer(gongEnumValue)
			if !isPresent {
				impl.DisableNodeCheckbox()
			}
		}
	}

	// disable link of notes to gongstruct / gongfields that are not displayed
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {

		isPresent := namesOfDisplayedGongnotes[gongNote.Name]
		_ = isPresent

		for _, gongLink := range gongNote.Links {
			impl := GetNodeBackPointer(gongLink)
			if !isPresent {
				impl.DisableNodeCheckbox()
			}
			if gongLink.Recv == "" {
				gongStructShapeWithThisNameIsPresent := namesOfDisplayedGongstructs[gongLink.Name]
				gongEnumShapeWithThisNameIsPresent := namesOfDisplayedGongenums[gongLink.Name]

				if !gongStructShapeWithThisNameIsPresent && !gongEnumShapeWithThisNameIsPresent {

					// no corresponding gong struct shape, therefore, disable the node
					impl.DisableNodeCheckbox()
				}
			} else {
				fieldName := gongLink.Recv + "." + gongLink.Name
				if ok := namesOfDisplayedGongfields[fieldName]; !ok {
					impl.DisableNodeCheckbox()
				}
			}
		}
	}

	// alternative :
	// parses all DSL object node,
	// get their implementation
	// get the DSL object, computes if disabling he checkbox is necessay , performs DisableNodeCheckbox on the DSL object node

	// for _, _node := range nodeCb.treeOfGongObjects.RootNodes {
	// 	for _, _node := range _node.Children {
	// 		for _, _node := range _node.Children {
	// 			switch nodeImpl := _node.Impl.(type) {
	// 			case *FieldImpl:
	// 				gongField := nodeImpl.field
	// 				switch fieldReal := gongField.(type) {
	// 				case *gong_models.PointerToGongStructField:
	// 					if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
	// 						nodeImpl.DisableNodeCheckbox()
	// 					}
	// 				case *gong_models.SliceOfPointerToGongStructField:
	// 					if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
	// 						nodeImpl.DisableNodeCheckbox()
	// 					}
	// 				default:
	// 				}
	// 			case *GongEnumValueImpl:
	// 			case *GongLinkImpl:
	// 			default:
	// 				_ = nodeImpl
	// 				log.Panic("No known implementation")
	// 			}
	// 		}
	// 	}
	// }

	// parse the tree of diagram elements and
	// get the corresponding gong object
	// and from the gong object, go to its node implementation and
	// set up that it has a diagram element and that is has to be checked
	for _, gongStructShape := range classdiagram.GongStructShapes {

		// get the gong struct related to the gong struct, and set t
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		gongStruct := (*gong_models.GetGongstructInstancesMap[gong_models.GongStruct]())[gongStructName]
		gongStructImpl := GetNodeBackPointer(gongStruct)
		gongStructImpl.CheckNode()

		for _, field := range gongStructShape.Fields {
			_ = field
			fieldName := gongdoc_models.IdentifierToFieldName(field.Identifier)

			// range over gongStruct fields (to be redone)
			for _, gongField := range gongStruct.Fields {
				if gongField.GetName() == fieldName {
					switch fieldReal := gongField.(type) {
					case *gong_models.GongBasicField:
						GetNodeBackPointer(fieldReal).CheckNode()
					case *gong_models.GongTimeField:
						GetNodeBackPointer(fieldReal).CheckNode()
					default:
					}
					continue
				}
			}
		}
		for _, link := range gongStructShape.Links {
			linkName := gongdoc_models.IdentifierToFieldName(link.Identifier)

			// range over gongStruct fields (to be redone)
			for _, gongField := range gongStruct.Fields {
				if gongField.GetName() == linkName {
					switch fieldReal := gongField.(type) {
					case *gong_models.PointerToGongStructField:
						GetNodeBackPointer(fieldReal).CheckNode()
					case *gong_models.SliceOfPointerToGongStructField:
						GetNodeBackPointer(fieldReal).CheckNode()
					default:
					}
				}
				// compute if the node for pointer or slice of pointer has to be disabled
			}
		}
	}
	for _, gongEnumShape := range classdiagram.GongEnumShapes {

		// get the gong struct related to the gong struct, and set t
		gongEnumName := gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier)
		gongEnum := (*gong_models.GetGongstructInstancesMap[gong_models.GongEnum]())[gongEnumName]
		gongEnumImpl := GetNodeBackPointer(gongEnum)
		gongEnumImpl.CheckNode()

		for _, gongEnumValueEntry := range gongEnumShape.GongEnumValueEntrys {
			_ = gongEnumValueEntry
			gongEnumValueName := gongdoc_models.IdentifierToFieldName(gongEnumValueEntry.Identifier)

			// range over gongStruct fields (to be redone)
			for _, gongEnumValue := range gongEnum.GongEnumValues {
				if gongEnumValue.GetName() == gongEnumValueName {
					GetNodeBackPointer(gongEnumValue).CheckNode()
				}
				continue
			}
		}
	}
	for _, noteShape := range classdiagram.NoteShapes {

		// get the gong struct related to the gong struct, and set t
		noteShapeName := gongdoc_models.IdentifierToGongObjectName(noteShape.Identifier)
		gongNote := (*gong_models.GetGongstructInstancesMap[gong_models.GongNote]())[noteShapeName]
		GetNodeBackPointer(gongNote).CheckNode()

		for _, nodeShapeLink := range noteShape.NoteShapeLinks {

			switch nodeShapeLink.Type {
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
				targetShapeGongStructName := gongdoc_models.IdentifierToGongObjectName(nodeShapeLink.Identifier)
				// range over gongStruct fields (to be redone)
				for _, link := range gongNote.Links {
					if link.GetName() == targetShapeGongStructName {
						GetNodeBackPointer(link).CheckNode()
					}
					continue
				}
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD:
				receiverName, fieldName :=
					gongdoc_models.IdentifierToReceiverAndFieldName(nodeShapeLink.Identifier)
					// range over gongStruct fields (to be redone)
				for _, link := range gongNote.Links {
					if link.GetName() == fieldName && link.Recv == receiverName {
						GetNodeBackPointer(link).CheckNode()
					}
					continue
				}
			}
		}
	}
}
