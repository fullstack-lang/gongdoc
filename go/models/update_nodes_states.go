package models

// updateNodesStates updates the tree of symbols
// according to the selected diagram
//
// ## The algorithm is as follow
//
// ## For the diagram nodes
//
// For the identifiers nodes
func updateNodesStates(stage *StageStruct, nodeCb *NodeCB) {

	nodeCb.treeOfGongObjects.UncheckAndDisable()

	updateDiagramsNodesStates(stage, nodeCb)

	// get the diagram
	classdiagram := nodeCb.diagramPackage.SelectedClassdiagram

	if classdiagram == nil {
		Stage.Commit()
		return
	}

	updateGongObjectsNodesStates(stage, nodeCb, classdiagram)

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
