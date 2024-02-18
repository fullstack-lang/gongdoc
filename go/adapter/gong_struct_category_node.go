package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

type GongStructCategoryNode struct {
	CategoryNodeBase
}

func NewGongStructCategoryNode(stage *gong_models.StageStruct, name string) *GongStructCategoryNode {
	return &GongStructCategoryNode{CategoryNodeBase: CategoryNodeBase{stage: stage, Name: name}}
}

// GetChildren implements bridge.Node.
func (categoryNode *GongStructCategoryNode) GetChildren() (children []bridge.ModelNode) {

	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](categoryNode.stage) {

		gongStructNode := NewGongStructNode(categoryNode.stage, gongStruct)
		children = append(children, gongStructNode)
	}

	slices.SortFunc(children, func(a, b bridge.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}
