package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func FillUpTreeOfGongObjects(gongdocStage *gongdoc_models.StageStruct) (gongTree *gongdoc_models.Tree) {

	// set up the gongTree to display elements
	gongTree = (&gongdoc_models.Tree{Name: "gong"}).Stage(gongdocStage)

	gongstructRootNode := (&gongdoc_models.Node{Name: "gongstructs"}).Stage(gongdocStage)
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)

	gongenumRootNode := (&gongdoc_models.Node{Name: "gongenums"}).Stage(gongdocStage)
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)

	gongNotesRootNode := (&gongdoc_models.Node{Name: "notes"}).Stage(gongdocStage)
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)

	return
}
