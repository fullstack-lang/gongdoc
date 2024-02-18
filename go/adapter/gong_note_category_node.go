package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

type GongNoteCategoryNode struct {
	CategoryNodeBase
}

func NewGongNoteCategoryNode(stage *gong_models.StageStruct, name string) *GongNoteCategoryNode {
	return &GongNoteCategoryNode{CategoryNodeBase: CategoryNodeBase{stage: stage, Name: name}}
}

// GetChildren implements bridge.Node.
func (categoryNode *GongNoteCategoryNode) GetChildren() (children []bridge.ModelNode) {

	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote](categoryNode.stage) {

		gongNoteNode := NewGongNoteNode(categoryNode.stage, gongNote)
		children = append(children, gongNoteNode)
	}

	slices.SortFunc(children, func(a, b bridge.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}
