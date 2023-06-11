package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func (nodeCb *NodeCB) FillUpTreeOfGongObjectsLegacy() (gongTree *gongdoc_models.Tree) {

	// set up the gongTree to display elements
	gongTree = (&gongdoc_models.Tree{Name: "gong"}).Stage(nodeCb.diagramPackage.Stage_)
	nodeCb.treeOfGongObjects = gongTree

	gongstructRootNode := (&gongdoc_models.Node{Name: "gongstructs"}).Stage(nodeCb.diagramPackage.Stage_)
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)

	gongenumRootNode := (&gongdoc_models.Node{Name: "gongenums"}).Stage(nodeCb.diagramPackage.Stage_)
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum](nodeCb.diagramPackage.ModelPkg.GetStage()) {

		nodeGongEnum := (&gongdoc_models.Node{Name: gongEnum.Name}).Stage(nodeCb.diagramPackage.Stage_)
		nodeGongEnum.HasCheckboxButton = true
		nodeGongEnum.IsExpanded = false

		gongEnumImpl := new(GongEnumImpl)
		gongEnumImpl.node = nodeGongEnum
		gongEnumImpl.gongEnum = gongEnum
		gongEnumImpl.nodeCb = nodeCb
		nodeGongEnum.Impl = gongEnumImpl

		// append to tree
		gongenumRootNode.Children = append(gongenumRootNode.Children, nodeGongEnum)

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			nodeGongEnumValue := (&gongdoc_models.Node{Name: gongEnumValue.GetName()}).Stage(nodeCb.diagramPackage.Stage_)

			nodeGongEnumValue.HasCheckboxButton = true

			gongEnumValueImpl := new(GongEnumValueImpl)
			gongEnumValueImpl.node = nodeGongEnumValue
			gongEnumValueImpl.gongEnumValue = gongEnumValue
			gongEnumValueImpl.nodeCb = nodeCb
			nodeGongEnumValue.Impl = gongEnumValueImpl

			// append to tree
			nodeGongEnum.Children = append(nodeGongEnum.Children, nodeGongEnumValue)

		}
	}

	gongNotesRootNode := (&gongdoc_models.Node{Name: "notes"}).Stage(nodeCb.diagramPackage.Stage_)
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)

	// generate the map to navigate from children to parents
	fieldName := gongdoc_models.GetAssociationName[gongdoc_models.Node]().Children[0].Name
	nodeCb.map_Children_Parent =
		gongdoc_models.GetSliceOfPointersReverseMap[gongdoc_models.Node, gongdoc_models.Node](fieldName, nodeCb.diagramPackage.Stage_)

	return

}
