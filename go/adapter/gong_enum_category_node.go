package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongEnumCategoryNode struct {
	ModelCategoryNodeBase
}

func NewGongEnumCategoryNode(stage *gong_models.StageStruct, name string) *GongEnumCategoryNode {
	return &GongEnumCategoryNode{ModelCategoryNodeBase: ModelCategoryNodeBase{stage: stage, Name: name}}
}

// GenerateChildren implements diagrammer.Node.
func (categoryNode *GongEnumCategoryNode) GenerateChildren() (children []diagrammer.ModelNode) {

	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum](categoryNode.stage) {

		gongEnumNode := NewGongEnumNode(categoryNode.stage, gongEnum)
		children = append(children, gongEnumNode)
	}

	slices.SortFunc(children, func(a, b diagrammer.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}
