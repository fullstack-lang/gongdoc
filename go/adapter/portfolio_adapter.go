package adapter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	gong_models "github.com/fullstack-lang/gong/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"

	"github.com/fullstack-lang/gongdoc/go/doc2svg"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type PortfolioAdapter struct {
	gongStage    *gong_models.StageStruct
	gongdocStage *gongdoc_models.StageStruct
	svgStage     *gongsvg_models.StageStruct
	diagrammer   *diagrammer.Diagrammer
}

func NewPortfolioAdapter(
	gongStage *gong_models.StageStruct,
	stage *gongdoc_models.StageStruct,
	svgStage *gongsvg_models.StageStruct,
) (adapter *PortfolioAdapter) {
	adapter = new(PortfolioAdapter)

	adapter.gongStage = gongStage
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

// GenerateDiagram implements diagrammer.Portfolio.
//
// It generated the SVG
func (portfolioAdapter *PortfolioAdapter) GenerateDiagram(diagramNode diagrammer.DiagramNode) (
	setOfModelNode map[diagrammer.ModelNode]diagrammer.Shape) {

	selectedDiagramNode, ok := diagramNode.(*ClassDiagramNode)
	if !ok {
		log.Fatalln("Not a classdiagram node")
	}

	selectedDiagram := selectedDiagramNode.classdiagramAdapter.classdiagram

	var diagramPackage *gongdoc_models.DiagramPackage
	for diagramPackage_ := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](selectedDiagramNode.stage) {
		diagramPackage = diagramPackage_
	}

	diagramPackage.SelectedClassdiagram = selectedDiagram

	docSVGMapper := doc2svg.NewDocSVGMapper(portfolioAdapter.svgStage)
	docSVGMapper.GenerateSvg(portfolioAdapter.gongdocStage)

	// compute the set of Model Node
	setOfModelNode = make(map[diagrammer.ModelNode]diagrammer.Shape)

	// 1. Create the map of model element to model node
	map_ModelElement_ModelNode := make(map[any]diagrammer.ModelNode)
	for modelNode := range portfolioAdapter.diagrammer.GetMap_modelNode_treeNode() {
		if modelElement := modelNode.GetElement(); modelElement != nil {
			map_ModelElement_ModelNode[modelElement] = modelNode
		}
	}

	// 2. Parse the selected diagram, for every shape, get the element model of the shape
	// use the aforementioned map to populate the set
	gongStructSet := *gong_models.GetGongstructInstancesMap[gong_models.GongStruct](portfolioAdapter.gongStage)
	for _, gongStructShape := range selectedDiagram.GongStructShapes {

		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		gongStruct, ok := gongStructSet[gongStructName]

		gongStructNode, ok := map_ModelElement_ModelNode[gongStruct]
		if !ok {
			log.Fatalln("unkown element", gongStructShape.Identifier)
		}

		setOfModelNode[gongStructNode] = &GongStructShapeAdapter{
			gongStructShape: gongStructShape,
			element:         gongStruct,
		}

		for _, fieldShape := range gongStructShape.Fields {
			fieldShapeName := gongdoc_models.IdentifierToFieldName(fieldShape.Identifier)

			// parse fields of gongstruct to match the field shape
			for _, field := range gongStruct.Fields {
				if field.GetName() == fieldShapeName {

					gongStructFieldNode, ok := map_ModelElement_ModelNode[field]
					if !ok {
						log.Fatalln("unkown element", fieldShape.Identifier)
					}

					setOfModelNode[gongStructFieldNode] = &GongStructFieldShapeAdapter{
						fieldShape: fieldShape,
						element:    field,
					}
				}
			}
		}

		for _, linkShape := range gongStructShape.Links {
			linkShapeName := gongdoc_models.IdentifierToFieldName(linkShape.Identifier)

			for _, link := range gongStruct.SliceOfPointerToGongStructFields {
				if link.GetName() == linkShapeName {
					linkNode, ok := map_ModelElement_ModelNode[link]
					if !ok {
						log.Fatalln("unkown element", linkShape.Identifier)
					}

					setOfModelNode[linkNode] = &LinkShapeAdapter{
						linkShape: linkShape,
						element:   link,
					}
				}
			}

			for _, link := range gongStruct.PointerToGongStructFields {
				if link.GetName() == linkShapeName {
					linkNode, ok := map_ModelElement_ModelNode[link]
					if !ok {
						log.Fatalln("unkown element", linkShape.Identifier)
					}

					setOfModelNode[linkNode] = &LinkShapeAdapter{
						linkShape: linkShape,
						element:   link,
					}
				}
			}
		}
	}

	gongEnumSet := *gong_models.GetGongstructInstancesMap[gong_models.GongEnum](portfolioAdapter.gongStage)
	for _, gongEnumShape := range selectedDiagram.GongEnumShapes {

		gongEnumName := gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier)
		gongEnum, ok := gongEnumSet[gongEnumName]

		gongEnumNode, ok := map_ModelElement_ModelNode[gongEnum]
		if !ok {
			log.Fatalln("unkown element", gongEnumShape.Identifier)
		}

		setOfModelNode[gongEnumNode] = &GongEnumShapeAdapter{
			gongEnumShape: gongEnumShape,
			element:       gongEnum,
		}

		for _, valueShape := range gongEnumShape.GongEnumValueEntrys {
			valueShapeName := gongdoc_models.IdentifierToFieldName(valueShape.Identifier)

			// parse values of gongstruct to match the value shape
			for _, value := range gongEnum.GongEnumValues {
				if value.GetName() == valueShapeName {

					gongEnumValueNode, ok := map_ModelElement_ModelNode[value]
					if !ok {
						log.Fatalln("unkown element", valueShape.Identifier)
					}

					setOfModelNode[gongEnumValueNode] = &GongEnumValueShapeAdapter{
						valueShape: valueShape,
						element:    value,
					}
				}
			}
		}
	}

	gongNoteSet := *gong_models.GetGongstructInstancesMap[gong_models.GongNote](portfolioAdapter.gongStage)
	for _, gongNoteShape := range selectedDiagram.NoteShapes {

		gongNoteName := gongdoc_models.IdentifierToGongObjectName(gongNoteShape.Identifier)
		gongNote, ok := gongNoteSet[gongNoteName]

		gongNoteNode, ok := map_ModelElement_ModelNode[gongNote]
		if !ok {
			log.Fatalln("unkown element", gongNoteShape.Identifier)
		}

		setOfModelNode[gongNoteNode] = &GongNoteShapeAdapter{
			gongNoteShape: gongNoteShape,
			element:       gongNote,
		}

		for _, noteLinkShape := range gongNoteShape.NoteShapeLinks {

			switch noteLinkShape.Type {
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
				nodeLinkShapeTarget := gongdoc_models.IdentifierToGongObjectName(noteLinkShape.Identifier)
				for _, link := range gongNote.Links {
					if link.GetName() == nodeLinkShapeTarget {
						linkNode, ok := map_ModelElement_ModelNode[link]
						if !ok {
							log.Fatalln("unkown element", noteLinkShape.Identifier)
						}

						setOfModelNode[linkNode] = &NoteLinkShapeAdapter{
							nodeShapeLink: noteLinkShape,
							element:       link,
						}
					}
				}
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD:
				receiver, fieldName := gongdoc_models.IdentifierToReceiverAndFieldName(noteLinkShape.Identifier)
				for _, link := range gongNote.Links {
					if link.GetName() == fieldName && link.Recv == receiver {
						linkNode, ok := map_ModelElement_ModelNode[link]
						if !ok {
							log.Fatalln("unkown element", noteLinkShape.Identifier)
						}

						setOfModelNode[linkNode] = &NoteLinkShapeAdapter{
							nodeShapeLink: noteLinkShape,
							element:       link,
						}
					}
				}

			}
		}
	}

	return
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
