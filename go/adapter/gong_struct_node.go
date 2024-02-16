package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongStructNode(
	stage *gong_models.StageStruct,
	name string) (gongStructNode *GongStructNode) {
	gongStructNode = new(GongStructNode)
	gongStructNode.stage = stage
	gongStructNode.Name = name
	return
}

type GongStructNode struct {
	stage *gong_models.StageStruct
	Name  string
}

// IsExpanded implements bridge.Node.
func (gongStructNode *GongStructNode) IsExpanded() bool {
	return true
}

// GetChildren implements bridge.Node.
func (gongStructNode *GongStructNode) GetChildren() (children []bridge.Node) {

	// for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](gongStructNode.stage) {

	// 	gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)
	// }

	return
}

// GetName implements bridge.Node.
func (gongStructNode *GongStructNode) GetName() string {
	return gongStructNode.Name
}

// IsNameEditable implements bridge.Node.
func (gongStructNode *GongStructNode) IsNameEditable() bool {
	return false
}
