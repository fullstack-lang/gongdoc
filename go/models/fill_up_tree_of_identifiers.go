package models

import gong_models "github.com/fullstack-lang/gong/go/models"

func FillUpTreeOfIdentifiers(pkgelt *DiagramPackage, nodeCb *NodeCB) {

	// set up the gongTree to display elements
	gongTree := (&Tree{Name: "gong"}).Stage()
	nodeCb.treeOfGongObjects = gongTree

	nodeCb.map_Identifier_Node = make(map[string]*Node)

	gongstructRootNode := (&Node{Name: "gongstructs"}).Stage()
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		nodeClassshape := (&Node{Name: gongStruct.Name}).Stage()
		nodeClassshape.Type = GONG_STRUCT
		nodeClassshape.HasCheckboxButton = true
		nodeClassshape.IsExpanded = false

		// set up the back pointer from the shape to the node
		gongStructImpl := new(GongStructImpl)
		gongStructImpl.node = nodeClassshape
		gongStructImpl.gongStruct = gongStruct
		gongStructImpl.nodeCb = nodeCb
		nodeClassshape.impl = gongStructImpl

		// append to the tree
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeClassshape)
		nodeCb.map_Identifier_Node[ShapenameToIdentifier(gongStruct.Name)] = nodeClassshape

		for _, field := range gongStruct.Fields {
			node2 := (&Node{Name: field.GetName()}).Stage()
			node2.Type = GONG_STRUCT_FIELD
			node2.HasCheckboxButton = true

			fieldImpl := new(FieldImpl)
			fieldImpl.node = nodeClassshape
			fieldImpl.field = field
			fieldImpl.nodeCb = nodeCb
			node2.impl = fieldImpl

			// append to tree
			nodeClassshape.Children = append(nodeClassshape.Children, node2)
			nodeCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(gongStruct.Name, field.GetName())] = node2
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

		for _, value := range gongEnum.GongEnumValues {
			node2 := (&Node{Name: value.GetName()}).Stage()
			node2.Type = GONG_ENUM_VALUE
			node2.HasCheckboxButton = true

			// append to tree
			node.Children = append(node.Children, node2)
			nodeCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(gongEnum.Name, value.GetName())] = node2
		}
	}

	gongNotesRootNode := (&Node{Name: "notes"}).Stage()
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {

		node := (&Node{Name: gongNote.Name}).Stage()
		node.HasCheckboxButton = true
		node.GongNote = gongNote
		node.Type = GONG_NOTE
		node.IsExpanded = true

		// append to tree
		gongNotesRootNode.Children = append(gongNotesRootNode.Children, node)
		nodeCb.map_Identifier_Node[ShapenameToIdentifier(gongNote.Name)] = node

		for _, gongLink := range gongNote.Links {
			node2 := (&Node{Name: gongLink.Name}).Stage()
			node2.HasCheckboxButton = true
			node2.Type = GONG_NOTE_LINK

			// append to tree
			node.Children = append(node.Children, node2)
			nodeCb.map_Identifier_Node[ShapeAndFieldnameToFieldIdentifier(gongNote.Name, gongLink.Name)] = node2
		}
	}

	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[Node]().Children[0].Name
	nodeCb.map_Children_Parent = GetSliceOfPointersReverseMap[Node, Node](fieldName)
}
