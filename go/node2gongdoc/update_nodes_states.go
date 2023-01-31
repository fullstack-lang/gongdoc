package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// updateNodesStates updates the tree of symbols
// according to the selected diagram
//
// ## The algorithm is as follow
//
// ## For the diagram nodes
//
// For the identifiers nodes
func updateNodesStates(stage *gongdoc_models.StageStruct, nodeCb *NodeCB) {

	nodeCb.treeOfGongObjects.UncheckAndDisable()

	updateDiagramsNodes(stage, nodeCb)

	// get the diagram
	classdiagram := nodeCb.diagramPackage.SelectedClassdiagram

	if classdiagram == nil {
		gongdoc_models.Stage.Commit()
		return
	}

	updateGongObjectsNodes(stage, nodeCb, classdiagram)

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
