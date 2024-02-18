package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongEnumCategoryNode(stage *gong_models.StageStruct, name string) *GongEnumCategoryNode {
	return &GongEnumCategoryNode{CategoryNodeBase: CategoryNodeBase{stage: stage, Name: name}}
}

type GongEnumCategoryNode struct {
	CategoryNodeBase
}

// GetChildren implements bridge.Node.
func (categoryNode *GongEnumCategoryNode) GetChildren() (children []bridge.ModelNode) {

	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum](categoryNode.stage) {

		gongEnumNode := NewGongEnumNode(categoryNode.stage, gongEnum)
		children = append(children, gongEnumNode)
	}

	slices.SortFunc(children, func(a, b bridge.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}
