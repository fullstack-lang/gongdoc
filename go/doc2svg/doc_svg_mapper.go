package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type DocSVGMapper struct {
}

func (docSVGMapper *DocSVGMapper) GenerateSvg(
	gongdocStage *gongdoc_models.StageStruct,
	gongsvgStage *gongsvg_models.StageStruct) {

	log.Println("DocSVGMapper.GenerateSvg")

	gongsvgStage.Reset()
	gongsvgStage.Commit()

	for diagramPackage := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](gongdocStage) {

		selectedDiagram := diagramPackage.SelectedClassdiagram
		if selectedDiagram == nil {
			return
		}

		svg := new(gongsvg_models.SVG).Stage(gongsvgStage)

		for _, gongstructShape := range selectedDiagram.GongStructShapes {

			rectLayer := new(gongsvg_models.Layer).Stage(gongsvgStage)
			svg.Layers = append(svg.Layers, rectLayer)

			rect := new(gongsvg_models.Rect).Stage(gongsvgStage)
			rectLayer.Rects = append(rectLayer.Rects, rect)
			rect.X = gongstructShape.Position.X
			rect.Y = gongstructShape.Position.Y

			rect.Width = gongstructShape.Width
			rect.Height = gongstructShape.Heigth

			rect.Stroke = "black"
			rect.StrokeWidth = 1
			rect.FillOpacity = 100
			rect.Color = gongsvg_models.Lightsalmon.ToString()

			text := new(gongsvg_models.Text).Stage(gongsvgStage)
			text.Name = gongdoc_models.IdentifierToGongObjectName(gongstructShape.Identifier)
			text.Content = text.Name
			text.X = rect.X + 10
			text.Y = rect.Y + 10
			text.Color = "black"
			text.FillOpacity = 1.0
			rectLayer.Texts = append(rectLayer.Texts, text)
		}
	}

	gongsvgStage.Commit()
}
