package adapter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	"github.com/fullstack-lang/gongdoc/go/doc2svg"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type PortfolioAdapter struct {
	gongdocStage *gongdoc_models.StageStruct
	svgStage     *gongsvg_models.StageStruct
	diagrammer   *diagrammer.Diagrammer
}

func NewPortfolioAdapter(
	stage *gongdoc_models.StageStruct,
	svgStage *gongsvg_models.StageStruct,
) (adapter *PortfolioAdapter) {
	adapter = new(PortfolioAdapter)

	adapter.gongdocStage = stage
	adapter.svgStage = svgStage

	return
}

func (portfolioAdapter *PortfolioAdapter) SetDiagrammer(diagrammer *diagrammer.Diagrammer) {
	portfolioAdapter.diagrammer = diagrammer
}

// IsInSelectionMode implements diagrammer.Portfolio.
func (*PortfolioAdapter) IsInSelectionMode() bool {
	return true // temp, will be false when a diagram is selected
}

// GetSelectedDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetSelectedDiagram() (diagram diagrammer.PortfolioNode) {
	return
}

// GetRootNodes implements bridge.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetChildren() (rootNodes []diagrammer.PortfolioNode) {
	classDiagramCategoryNode := NewClassDiagramCategoryNode(portfolioAdapter.gongdocStage, "class diagrams", portfolioAdapter.diagrammer)
	rootNodes = append(rootNodes, classDiagramCategoryNode)

	return
}

// GenerateSVG implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GenerateSVG() {
	docSVGMapper := doc2svg.NewDocSVGMapper(portfolioAdapter.svgStage)
	docSVGMapper.GenerateSvg(portfolioAdapter.gongdocStage)
}

// AddDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) AddDiagram(parentPortfolioNode diagrammer.PortfolioNode) diagrammer.PortfolioNode {

	gongdocStage := portfolioAdapter.gongdocStage
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

	classDiagramNode := NewClassDiagramNode(gongdocStage, classdiagram, portfolioAdapter.diagrammer)

	return classDiagramNode
}
