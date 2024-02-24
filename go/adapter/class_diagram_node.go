package adapter

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ClassDiagramNode struct {
	stage        *gongdoc_models.StageStruct
	classDiagram *gongdoc_models.Classdiagram
	diagrammer   *diagrammer.Diagrammer
}

// HasHadButton implements diagrammer.PortfolioNode.
func (*ClassDiagramNode) HasHadButton() bool {
	return false
}

func NewClassDiagramNode(
	stage *gongdoc_models.StageStruct,
	classDiagram *gongdoc_models.Classdiagram,
	diagrammer *diagrammer.Diagrammer,

) (classDiagramNode *ClassDiagramNode) {
	classDiagramNode = &ClassDiagramNode{stage: stage}

	classDiagramNode.classDiagram = classDiagram
	classDiagramNode.diagrammer = diagrammer

	return
}

// GetChildren implements bridge.Node.
func (classDiagramNode *ClassDiagramNode) GetChildren() (children []diagrammer.PortfolioNode) {
	return
}

// GetName implements bridge.Node.
func (classDiagramNode *ClassDiagramNode) GetName() string {
	return classDiagramNode.classDiagram.GetName()
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

// OnCheckboxButtonCheck implements diagrammer.PortfolioNode.
func (classDiagramNode *ClassDiagramNode) OnCheckboxButtonCheck() {
	var diagramPackage *gongdoc_models.DiagramPackage
	for diagramPackage_ := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](classDiagramNode.stage) {
		diagramPackage = diagramPackage_
	}
	diagramPackage.SelectedClassdiagram = classDiagramNode.classDiagram
	classDiagramNode.diagrammer.GenerateSVG()
	// panic("unimplemented")
}
