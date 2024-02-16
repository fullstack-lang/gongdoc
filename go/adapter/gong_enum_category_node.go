package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongEnumCategoryNode(
	stage *gong_models.StageStruct,
	name string) (categoryNode *GongEnumCategoryNode) {
	categoryNode = new(GongEnumCategoryNode)
	categoryNode.stage = stage
	categoryNode.Name = name
	return
}

type GongEnumCategoryNode struct {
	stage *gong_models.StageStruct
	Name  string
}

func (gongEnumCategoryNode *GongEnumCategoryNode) IsExpanded() bool {
	return true
}

func (gongEnumCategoryNode *GongEnumCategoryNode) HasCheckboxButton() bool {
	return false
}

// GetChildren implements bridge.Node.
func (categoryNode *GongEnumCategoryNode) GetChildren() (children []bridge.Node) {

	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum](categoryNode.stage) {

		gongEnumNode := NewGongEnumNode(categoryNode.stage, gongEnum)
		children = append(children, gongEnumNode)
	}

	slices.SortFunc(children, func(a, b bridge.Node) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}

// GetName implements bridge.Node.
func (categoryNode *GongEnumCategoryNode) GetName() string {
	return categoryNode.Name
}

// IsNameEditable implements bridge.Node.
func (categoryNode *GongEnumCategoryNode) IsNameEditable() bool {
	return false
}
