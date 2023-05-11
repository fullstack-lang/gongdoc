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

			rect.Stroke = "transparant"
			rect.StrokeWidth = 0

			rect.FillOpacity = 100
			rect.Color = gongsvg_models.Lightsalmon.ToString()

			// moveability
			rect.CanMoveHorizontaly = true
			rect.CanMoveVerticaly = true

			rect.CanHaveBottomHandle = true
			rect.CanHaveLeftHandle = true
			rect.CanHaveTopHandle = true
			rect.CanHaveRightHandle = true

			title := new(gongsvg_models.RectAnchoredText).Stage(gongsvgStage)
			title.Name = gongdoc_models.IdentifierToGongObjectName(gongstructShape.Identifier)
			title.Content = title.Name

			// title position
			title.X_Offset = 0
			title.Y_Offset = 20
			title.RectAnchorType = gongsvg_models.RECT_ANCHOR_TOP
			title.TextAnchorType = gongsvg_models.TEXT_ANCHOR_CENTER
			title.FontWeight = "bold"

			//

			title.Color = "black"
			title.FillOpacity = 1.0
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
		}
	}

	gongsvgStage.Commit()
}
