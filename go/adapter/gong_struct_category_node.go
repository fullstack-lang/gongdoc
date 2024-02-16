package adapter

import (
	"cmp"
	"slices"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongStructCategoryNode(
	stage *gong_models.StageStruct,
	name string) (categoryNode *GongStructCategoryNode) {
	categoryNode = new(GongStructCategoryNode)
	categoryNode.stage = stage
	categoryNode.Name = name
	return
}

type GongStructCategoryNode struct {
	stage *gong_models.StageStruct
	Name  string
}

// GetChildren implements bridge.Node.
func (categoryNode *GongStructCategoryNode) GetChildren() (children []bridge.Node) {

	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](categoryNode.stage) {

		gongStructNode := NewGongStructNode(categoryNode.stage, gongStruct.Name)
		children = append(children, gongStructNode)
	}

	slices.SortFunc(children, func(a, b bridge.Node) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}

// GetName implements bridge.Node.
func (categoryNode *GongStructCategoryNode) GetName() string {
	return categoryNode.Name
}

// IsNameEditable implements bridge.Node.
func (categoryNode *GongStructCategoryNode) IsNameEditable() bool {
	return false
}
