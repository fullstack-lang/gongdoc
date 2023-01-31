package models

import (
	"log"
)

func FillUpNodeTree(diagramPackage *DiagramPackage) {

	// a node tree is agnostic of the node types it manages
	// therefore, a callback functiion is necessary
	nodeCb := new(NodeCB)
	nodeCb.diagramPackage = diagramPackage

	nodeCb.FillUpDiagramNodeTree(diagramPackage)
	nodeCb.FillUpTreeOfGongObjects(diagramPackage)

	updateNodesStates(&Stage, nodeCb)

	// set callbacks on node updates
	Stage.OnAfterNodeUpdateCallback = nodeCb
	Stage.OnAfterNodeCreateCallback = nodeCb
	Stage.OnAfterNodeDeleteCallback = nodeCb

	log.Printf("Parse found %d diagrams\n", len(diagramPackage.Classdiagrams))
}
