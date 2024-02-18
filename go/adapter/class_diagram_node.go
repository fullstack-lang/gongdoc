package adapter

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/bridge"
)

type ClassDiagramNode struct {
	stage        *gongdoc_models.StageStruct
	ClassDiagram *gongdoc_models.Classdiagram
}

func NewClassDiagramNode(
	stage *gongdoc_models.StageStruct,
	classDiagram *gongdoc_models.Classdiagram) (classDiagramNode *ClassDiagramNode) {
	classDiagramNode = &ClassDiagramNode{stage: stage}
	classDiagramNode.ClassDiagram = classDiagram
	return
}

// GetChildren implements bridge.Node.
func (classDiagramNode *ClassDiagramNode) GetChildren() (children []bridge.PortfolioNode) {
	return
}

// GetName implements bridge.Node.
func (classDiagramNode *ClassDiagramNode) GetName() string {
	return classDiagramNode.ClassDiagram.GetName()
}

// HasCheckboxButton implements bridge.PortfolioNode.
func (*ClassDiagramNode) HasCheckboxButton() bool {
	return true
}

// IsExpanded implements bridge.PortfolioNode.
func (*ClassDiagramNode) IsExpanded() bool {
	return true
}

// IsNameEditable implements bridge.PortfolioNode.
func (*ClassDiagramNode) IsNameEditable() bool {
	return false
}
