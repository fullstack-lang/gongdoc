package adapter

import (
	"cmp"
	"slices"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ClassDiagramCategoryNode struct {
	stage      *gongdoc_models.StageStruct
	Name       string
	diagrammer *diagrammer.Diagrammer
}

func NewClassDiagramCategoryNode(
	stage *gongdoc_models.StageStruct,
	name string,
	diagrammer *diagrammer.Diagrammer,
) (categoryNode *ClassDiagramCategoryNode) {
	categoryNode = new(ClassDiagramCategoryNode)

	categoryNode.stage = stage
	categoryNode.Name = name
	categoryNode.diagrammer = diagrammer

	return
}

func (ClassDiagramCategoryNode *ClassDiagramCategoryNode) IsExpanded() bool {
	return true
}

// GetChildren implements bridge.Node.
func (categoryNode *ClassDiagramCategoryNode) GetChildren() (children []diagrammer.PortfolioNode) {

	for classDiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](categoryNode.stage) {

		classDiagramNode := NewClassDiagramNode(categoryNode.stage, classDiagram, categoryNode.diagrammer)
		children = append(children, classDiagramNode)
	}

	slices.SortFunc(children, func(a, b diagrammer.PortfolioNode) int {
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

// HasAddButton implements diagrammer.PortfolioNode.
func (*ClassDiagramCategoryNode) HasAddButton() bool {
	return true
}
