package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// computeGongNodesConfiguration set up all gong nodes according to the classdiagram
func (nodeCb *NodeCB) computeGongNodesConfiguration(
	stage *gongdoc_models.StageStruct,
	classdiagram *gongdoc_models.Classdiagram) {

	computeGongNodesConfigurations(stage, classdiagram, nodeCb.treeOfGongObjects)
}
