package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongEnumCategoryNode struct {
	CategoryNodeBase
}

func NewGongEnumCategoryNode(stage *gong_models.StageStruct, name string) *GongEnumCategoryNode {
	return &GongEnumCategoryNode{CategoryNodeBase: CategoryNodeBase{stage: stage, Name: name}}
}

// GetChildren implements bridge.Node.
func (categoryNode *GongEnumCategoryNode) GetChildren() (children []diagrammer.ModelNode) {

	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum](categoryNode.stage) {

		gongEnumNode := NewGongEnumNode(categoryNode.stage, gongEnum)
		children = append(children, gongEnumNode)
	}

	slices.SortFunc(children, func(a, b diagrammer.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}
