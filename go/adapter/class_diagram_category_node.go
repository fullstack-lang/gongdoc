package adapter

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ClassDiagramCategoryNode struct {
	gongdocStage *gongdoc_models.StageStruct
	Name         string
	diagrammer   *diagrammer.Diagrammer
}

func NewClassDiagramCategoryNode(
	stage *gongdoc_models.StageStruct,
	name string,
	diagrammer *diagrammer.Diagrammer,
) (categoryNode *ClassDiagramCategoryNode) {
	categoryNode = new(ClassDiagramCategoryNode)

	categoryNode.gongdocStage = stage
	categoryNode.Name = name
	categoryNode.diagrammer = diagrammer

	return
}

func (ClassDiagramCategoryNode *ClassDiagramCategoryNode) IsExpanded() bool {
	return true
}

// GetChildren implements bridge.Node.
func (categoryNode *ClassDiagramCategoryNode) GetChildren() (children []diagrammer.PortfolioNode) {

	for classDiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](categoryNode.gongdocStage) {

		classDiagramNode := NewClassDiagramNode(categoryNode.gongdocStage, classDiagram, categoryNode.diagrammer)
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

// HasAddDiagramButton
func (*ClassDiagramCategoryNode) HasAddDiagramButton() bool {
	return true
}

// AddDiagram implements diagrammer.Portfolio.
func (classDiagramCategoryNode *ClassDiagramCategoryNode) AddDiagram() diagrammer.PortfolioNode {

	gongdocStage := classDiagramCategoryNode.gongdocStage
	var diagramPackage *gongdoc_models.DiagramPackage
	for diagramPackage_ := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](gongdocStage) {
		diagramPackage = diagramPackage_
	}

	// check unicity of name, otherwise, add an index
	var hasNameCollision bool
	initialName := "Default"
	name := initialName
	index := 0
	// loop until the name of the new diagram is not in collision with an existing
	// diagram name
	for index == 0 || hasNameCollision {
		index++
		hasNameCollision = false
		for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
			if classdiagram.Name == name {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			name = initialName + fmt.Sprintf("_%d", index)
		}
	}

	classdiagram := (&gongdoc_models.Classdiagram{Name: name}).Stage(gongdocStage)
	diagramPackage.Classdiagrams =
		append(diagramPackage.Classdiagrams, classdiagram)

	filepath := filepath.Join(
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage,
			"../diagrams"),
		classdiagram.Name) + ".go"
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Cannot open diagram file" + err.Error())
	}
	defer file.Close()

	gongdocStage.Commit()

	// now save the diagram
	gongdocStage.Checkout()
	gongdocStage.Unstage()
	gongdoc_models.StageBranch(gongdocStage, classdiagram)

	gongdoc_models.SetupMapDocLinkRenaming(diagramPackage.ModelPkg.Stage_, gongdocStage)

	gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

	// restore the original stage
	gongdocStage.Unstage()
	gongdocStage.Checkout()

	classDiagramNode := NewClassDiagramNode(gongdocStage, classdiagram, classDiagramCategoryNode.diagrammer)

	return classDiagramNode
}
