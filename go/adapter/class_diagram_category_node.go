package adapter

import (
	"cmp"
	"slices"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewClassDiagramCategoryNode(
	stage *gongdoc_models.StageStruct,
	name string) (categoryNode *ClassDiagramCategoryNode) {
	categoryNode = new(ClassDiagramCategoryNode)
	categoryNode.stage = stage
	categoryNode.Name = name
	return
}

type ClassDiagramCategoryNode struct {
	stage *gongdoc_models.StageStruct
	Name  string
}

func (ClassDiagramCategoryNode *ClassDiagramCategoryNode) IsExpanded() bool {
	return true
}

func (ClassDiagramCategoryNode *ClassDiagramCategoryNode) HasCheckboxButton() bool {
	return false
}

// GetChildren implements bridge.Node.
func (categoryNode *ClassDiagramCategoryNode) GetChildren() (children []bridge.PortfolioNode) {

	// for ClassDiagram := range *gong_models.GetClassDiagramInstancesSet[gong_models.ClassDiagram](categoryNode.stage) {

	// 	ClassDiagramNode := NewClassDiagramNode(categoryNode.stage, ClassDiagram)
	// 	children = append(children, ClassDiagramNode)
	// }

	slices.SortFunc(children, func(a, b bridge.PortfolioNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return
}

// GetName implements bridge.Node.
func (categoryNode *ClassDiagramCategoryNode) GetName() string {
	return categoryNode.Name
}

// IsNameEditable implements bridge.Node.
func (categoryNode *ClassDiagramCategoryNode) IsNameEditable() bool {
	return false
}
