package adapter

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/doc2svg"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ClassDiagramNode struct {
	portfolioAdapter    *PortfolioAdapter
	diagrammer          *diagrammer.Diagrammer
	classdiagramAdapter *ClassdiagramAdapter
}

func NewClassDiagramNode(
	portfolioAdapter *PortfolioAdapter,
	classDiagram *gongdoc_models.Classdiagram,
	diagrammer *diagrammer.Diagrammer,

) (classDiagramNode *ClassDiagramNode) {
	classDiagramNode = &ClassDiagramNode{
		portfolioAdapter: portfolioAdapter,
	}

	classDiagramNode.classdiagramAdapter = &ClassdiagramAdapter{
		classdiagram: classDiagram,
	}
	classDiagramNode.diagrammer = diagrammer

	return
}

// GetChildren implements bridge.Node.
func (classDiagramNode *ClassDiagramNode) GetChildren() (children []diagrammer.PortfolioNode) {
	return
}

// GetName implements bridge.Node.
func (classDiagramNode *ClassDiagramNode) GetName() string {
	return classDiagramNode.classdiagramAdapter.GetName()
}

// GetDiagram implements bridge.PortfolioNode.
func (classDiagramNode *ClassDiagramNode) GetDiagram() diagrammer.Diagram {
	return classDiagramNode.classdiagramAdapter
}

// IsExpanded implements bridge.PortfolioNode.
func (*ClassDiagramNode) IsExpanded() bool {
	return true
}

// IsNameEditable implements bridge.PortfolioNode.
func (*ClassDiagramNode) IsNameEditable() bool {
	return false
}

// DisplayDiagram implements diagrammer.Portfolio.
//
// It generated the SVG
func (classDiagramNode *ClassDiagramNode) DisplayDiagram() (
	setOfModelNode map[diagrammer.ModelNode]diagrammer.Shape) {

	gongdocStage := classDiagramNode.portfolioAdapter.gongdocStage
	gongsvgStage := classDiagramNode.portfolioAdapter.gongsvgStage
	gongStage := classDiagramNode.portfolioAdapter.gongStage

	var diagramPackage *gongdoc_models.DiagramPackage
	for diagramPackage_ := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](gongdocStage) {
		diagramPackage = diagramPackage_
	}

	diagramPackage.SelectedClassdiagram = classDiagramNode.classdiagramAdapter.classdiagram
	selectedClassdiagram := diagramPackage.SelectedClassdiagram

	docSVGMapper := doc2svg.NewDocSVGMapper(gongsvgStage)
	docSVGMapper.GenerateSvg(gongdocStage)

	// compute the set of Model Node
	setOfModelNode = make(map[diagrammer.ModelNode]diagrammer.Shape)

	// 1. Create the map of model element to model node
	map_ModelElement_ModelNode := make(map[any]diagrammer.ModelNode)
	for modelNode := range classDiagramNode.diagrammer.GetMap_elementNode_treeNode() {

		if elementNode, ok := modelNode.(diagrammer.ElementNode); ok {
			if modelElement := elementNode.GetElement(); modelElement != nil {
				map_ModelElement_ModelNode[modelElement] = modelNode
			}
		}
	}

	// 2. Parse the selected diagram, for every shape, get the element model of the shape
	// use the aforementioned map to populate the set
	gongStructSet := *gong_models.GetGongstructInstancesMap[gong_models.GongStruct](gongStage)
	for _, gongStructShape := range selectedClassdiagram.GongStructShapes {

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

	gongEnumSet := *gong_models.GetGongstructInstancesMap[gong_models.GongEnum](gongStage)
	for _, gongEnumShape := range selectedClassdiagram.GongEnumShapes {

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

	gongNoteSet := *gong_models.GetGongstructInstancesMap[gong_models.GongNote](gongStage)
	for _, gongNoteShape := range selectedClassdiagram.NoteShapes {

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
