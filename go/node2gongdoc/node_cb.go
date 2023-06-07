package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// NodeCB is the singloton callback implementation of CUD operations on node
type NodeCB struct {

	// diagramPackageNode is used for iterating on classdiagram nodes
	// and for getting the selected diagram
	diagramPackageNode *gongdoc_models.Node
	diagramPackage     *gongdoc_models.DiagramPackage

	// treeOfGongObjects is the root of all nodes related to gong objects
	// it is necessary for performing operation on all elements of the tree
	treeOfGongObjects *gongdoc_models.Tree

	// map_Children_Parent is a map to navigate from a children node to its parent node
	// it is set up once at the init phase
	map_Children_Parent map[*gongdoc_models.Node]*gongdoc_models.Node
}

// GetSelectedClassdiagram
func (nodesCb *NodeCB) GetSelectedClassdiagram() (classdiagram *gongdoc_models.Classdiagram) {

	classdiagram = nodesCb.diagramPackage.SelectedClassdiagram
	return
}

// OnAfterUpdate is called each time the end user interacts
// with any node. The front commit the state of the front node
// to the back ([frontNode] in the function).
// Therefore, the [stageNode] is now different from the [frontNode].
//
// This functiion and its siblings analyse which field of the
// node has changed and performs all necessary actions
func (nodeCb *NodeCB) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	if stagedNode.IsExpanded != frontNode.IsExpanded {
		// setting the value of the staged node	to the new value
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	if stagedNode.Impl != nil {
		stagedNode.Impl.OnAfterUpdate(stage, stagedNode, frontNode)
		nodeCb.computeNodesConfiguration(stage)
	}

}

// OnAfterCreate is another callback
func (nodeCb *NodeCB) OnAfterCreate(
	gongdocStage *gongdoc_models.StageStruct,
	node *gongdoc_models.Node) {

}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (nodeCb *NodeCB) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

}

// computeNodesConfiguration computes both trees
func (nodeCb *NodeCB) computeNodesConfiguration(gongdocStage *gongdoc_models.StageStruct) {

	computeNodeConfs(gongdocStage,
		nodeCb.diagramPackageNode,
		nodeCb.diagramPackage,
		nodeCb.treeOfGongObjects)

}

func (nodesCb *NodeCB) computeDiagramNodesConfigurations(stage *gongdoc_models.StageStruct) {

	computeClassdiagramNodesConfigurations(
		nodesCb.diagramPackageNode,
		nodesCb.diagramPackage,
		stage)
}
