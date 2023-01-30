package models

import gong_models "github.com/fullstack-lang/gong/go/models"

func FillUpTreeOfGongObjects(pkgelt *DiagramPackage, nodeCb *NodeCB) {

	// set up the gongTree to display elements
	gongTree := (&Tree{Name: "gong"}).Stage()
	nodeCb.treeOfGongObjects = gongTree

	nodeCb.map_Identifier_Node = make(map[string]*Node)
	nodeCb.map_gongObject_gongObjectImpl = make(map[gong_models.GongStructInterface]NodeImplInterface)

	gongstructRootNode := (&Node{Name: "gongstructs"}).Stage()
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		nodeGongstruct := (&Node{Name: gongStruct.Name}).Stage()
		nodeGongstruct.Type = GONG_STRUCT
		nodeGongstruct.HasCheckboxButton = true
		nodeGongstruct.IsExpanded = false

		// set up the back pointer from the shape to the node
		gongStructImpl := new(GongStructImpl)
		gongStructImpl.node = nodeGongstruct
		gongStructImpl.gongStruct = gongStruct
		gongStructImpl.nodeCb = nodeCb
		nodeGongstruct.impl = gongStructImpl

		// append to the tree
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)
		nodeCb.map_Identifier_Node[ShapenameToIdentifier(gongStruct.Name)] = nodeGongstruct
		nodeCb.map_gongObject_gongObjectImpl[gongStruct] = gongStructImpl

		for _, field := range gongStruct.Fields {
			nodeGongField := (&Node{Name: field.GetName()}).Stage()
			nodeGongField.Type = GONG_STRUCT_FIELD
			nodeGongField.HasCheckboxButton = true

			fieldImpl := new(FieldImpl)
			fieldImpl.node = nodeGongstruct
			fieldImpl.field = field
			fieldImpl.nodeCb = nodeCb
			nodeGongField.impl = fieldImpl

			// append to tree
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeGongField)
			nodeCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(gongStruct.Name, field.GetName())] = nodeGongField
			nodeCb.map_gongObject_gongObjectImpl[field] = fieldImpl
		}
	}

	gongenumRootNode := (&Node{Name: "gongenums"}).Stage()
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {

		node := (&Node{Name: gongEnum.Name}).Stage()
		node.HasCheckboxButton = true
		node.IsExpanded = false
		node.Type = GONG_ENUM

		gongEnumImpl := new(GongEnumImpl)
		gongEnumImpl.node = node
		gongEnumImpl.gongEnum = gongEnum
		gongEnumImpl.nodeCb = nodeCb
		node.impl = gongEnumImpl

		// append to tree
		gongenumRootNode.Children = append(gongenumRootNode.Children, node)
		nodeCb.map_Identifier_Node[ShapenameToIdentifier(gongEnum.Name)] = node
		nodeCb.map_gongObject_gongObjectImpl[gongEnum] = gongEnumImpl

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			nodeGongEnumValue := (&Node{Name: gongEnumValue.GetName()}).Stage()
			nodeGongEnumValue.Type = GONG_ENUM_VALUE
			nodeGongEnumValue.HasCheckboxButton = true

			gongEnumValueImpl := new(GongEnumValueImpl)
			gongEnumValueImpl.node = node
			gongEnumValueImpl.gongEnumValue = gongEnumValue
			gongEnumValueImpl.nodeCb = nodeCb
			nodeGongEnumValue.impl = gongEnumValueImpl

			// append to tree
			node.Children = append(node.Children, nodeGongEnumValue)
			nodeCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(gongEnum.Name, gongEnumValue.GetName())] = nodeGongEnumValue
			nodeCb.map_gongObject_gongObjectImpl[gongEnumValue] = gongEnumValueImpl
		}
	}

	gongNotesRootNode := (&Node{Name: "notes"}).Stage()
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {

		node := (&Node{Name: gongNote.Name}).Stage()
		node.HasCheckboxButton = true
		node.Type = GONG_NOTE
		node.IsExpanded = true

		gongNoteImpl := new(GongNoteImpl)
		gongNoteImpl.node = node
		gongNoteImpl.gongNote = gongNote
		gongNoteImpl.nodeCb = nodeCb
		node.impl = gongNoteImpl

		// append to tree
		gongNotesRootNode.Children = append(gongNotesRootNode.Children, node)
		nodeCb.map_Identifier_Node[ShapenameToIdentifier(gongNote.Name)] = node
		nodeCb.map_gongObject_gongObjectImpl[gongNote] = gongNoteImpl

		for _, gongLink := range gongNote.Links {
			nodeGongLink := (&Node{Name: gongLink.Name}).Stage()
			nodeGongLink.HasCheckboxButton = true
			nodeGongLink.Type = GONG_NOTE_LINK

			gongLinkImpl := new(GongLinkImpl)
			gongLinkImpl.node = node
			gongLinkImpl.gongLink = gongLink
			gongLinkImpl.nodeCb = nodeCb
			nodeGongLink.impl = gongLinkImpl

			// append to tree
			node.Children = append(node.Children, nodeGongLink)
			nodeCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(gongNote.Name, gongLink.Name)] = nodeGongLink
			nodeCb.map_gongObject_gongObjectImpl[gongLink] = gongLinkImpl
		}
	}

	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[Node]().Children[0].Name
	nodeCb.map_Children_Parent = GetSliceOfPointersReverseMap[Node, Node](fieldName)
}
