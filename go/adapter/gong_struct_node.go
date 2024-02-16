package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongStructNode(
	stage *gong_models.StageStruct,
	name string) (categoryNode *GongStructNode) {
	categoryNode = new(GongStructNode)
	categoryNode.stage = stage
	categoryNode.Name = name
	return
}

type GongStructNode struct {
	stage *gong_models.StageStruct
	Name  string
}

// GetChildren implements bridge.Node.
func (categoryNode *GongStructNode) GetChildren() (children []bridge.Node) {

	// for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](categoryNode.stage) {

	// 	gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)
	// }

	return
}

// GetName implements bridge.Node.
func (categoryNode *GongStructNode) GetName() string {
	return categoryNode.Name
}

// IsNameEditable implements bridge.Node.
func (categoryNode *GongStructNode) IsNameEditable() bool {
	return false
}
