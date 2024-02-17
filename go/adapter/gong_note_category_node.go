package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongNoteCategoryNode(
	stage *gong_models.StageStruct,
	name string) (categoryNode *GongNoteCategoryNode) {
	categoryNode = new(GongNoteCategoryNode)
	categoryNode.stage = stage
	categoryNode.Name = name
	return
}

type GongNoteCategoryNode struct {
	stage *gong_models.StageStruct
	Name  string
}

func (gongNoteCategoryNode *GongNoteCategoryNode) IsExpanded() bool {
	return true
}

func (gongNoteCategoryNode *GongNoteCategoryNode) HasCheckboxButton() bool {
	return false
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

// GetName implements bridge.Node.
func (categoryNode *GongNoteCategoryNode) GetName() string {
	return categoryNode.Name
}

// IsNameEditable implements bridge.Node.
func (categoryNode *GongNoteCategoryNode) IsNameEditable() bool {
	return false
}
