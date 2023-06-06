package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type NodeImplRootOfClassDiagrams struct {
}

func NewNodeImplRootOfClassDiagrams() (nodeImplRootOfClassDiagrams *NodeImplRootOfClassDiagrams) {

	nodeImplRootOfClassDiagrams = new(NodeImplRootOfClassDiagrams)

	return
}

func (nodeImplRootOfClassDiagrams *NodeImplRootOfClassDiagrams) NodeUpdated(
	stage *gongdoc_models.StageStruct,
	stagedNode,
	updatedNode *gongdoc_models.Node) {

}
