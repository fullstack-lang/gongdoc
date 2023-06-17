package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func FillUpTreeOfDiagramNodes(
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
	treeOfGongObjects *gongdoc_models.Tree,
) (rootOfClassdiagramsNode *gongdoc_models.Node) {

	// create a tree for classdiagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc"}).Stage(gongdocStage)
	gongTreeTree := (&gongtree_models.Tree{Name: "gongdoc"}).Stage(gongtreeStage)

	// add the root of class diagrams
	rootOfClassdiagramsNode = (&gongdoc_models.Node{Name: "Class diagrams"}).Stage(gongdocStage)
	rootOfClassdiagramsNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, rootOfClassdiagramsNode)

	rootOfClassdiagramsTreeNode := (&gongtree_models.Node{Name: "Class diagrams"}).Stage(gongtreeStage)
	rootOfClassdiagramsTreeNode.IsExpanded = true
	gongTreeTree.RootNodes = append(gongTreeTree.RootNodes, rootOfClassdiagramsTreeNode)

	// add add button
	addDocButton := (&gongdoc_models.Button{
		Name: diagramPackage.Name + " " + string(BUTTON_add),
		Icon: string(BUTTON_add)}).Stage(gongdocStage)
	rootOfClassdiagramsNode.Buttons = append(rootOfClassdiagramsNode.Buttons, addDocButton)
	addDocButton.Impl = NewButtonImplRootOfClassdiagrams(
		diagramPackage,
		rootOfClassdiagramsNode,
		treeOfGongObjects,
		BUTTON_add,
	)

	// add one node for each diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
		classdiagramNode := NewClassdiagramNode(classdiagram, diagramPackage, rootOfClassdiagramsNode, treeOfGongObjects)
		rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)
	}

	return
}
