package models

import gong_models "github.com/fullstack-lang/gong/go/models"

func FillUpCodeNodeTree(pkgelt *Pkgelt) {
	// set up the gongTree to display elements
	gongTree := (&Tree{Name: "gong", Type: TREE_OF_SYMBOLS}).Stage()
	gongstructRootNode := (&Node{Name: "gongstructs"}).Stage()
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		node := (&Node{Name: gongStruct.Name}).Stage()
		node.HasCheckboxButton = true
		gongstructRootNode.Children = append(gongstructRootNode.Children, node)

		for _, field := range gongStruct.Fields {
			node2 := (&Node{Name: field.GetName()}).Stage()
			node2.HasCheckboxButton = true
			node.Children = append(node.Children, node2)
		}
	}

	gongenumRootNode := (&Node{Name: "gongenums"}).Stage()
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {

		node := (&Node{Name: gongEnum.Name}).Stage()
		node.HasCheckboxButton = true
		gongenumRootNode.Children = append(gongenumRootNode.Children, node)

		for _, value := range gongEnum.GongEnumValues {
			node2 := (&Node{Name: value.GetName()}).Stage()
			node2.HasCheckboxButton = true
			node.Children = append(node.Children, node2)
		}
	}
}
