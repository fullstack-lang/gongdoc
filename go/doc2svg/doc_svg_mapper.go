package doc2svg

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type DocSVGMapper struct {
	map_Rect_GongstructShape map[*gongsvg_models.Rect]*gongdoc_models.GongStructShape
	map_GongstructShape_Rect map[*gongdoc_models.GongStructShape]*gongsvg_models.Rect
	map_Structname_Rect      map[string]*gongsvg_models.Rect
}

func (docSVGMapper *DocSVGMapper) GenerateSvg(
	gongdocStage *gongdoc_models.StageStruct,
	gongsvgStage *gongsvg_models.StageStruct) {

	log.Println("DocSVGMapper.GenerateSvg")

	docSVGMapper.map_Rect_GongstructShape = make(map[*gongsvg_models.Rect]*gongdoc_models.GongStructShape)
	docSVGMapper.map_GongstructShape_Rect = make(map[*gongdoc_models.GongStructShape]*gongsvg_models.Rect)
	docSVGMapper.map_Structname_Rect = make(map[string]*gongsvg_models.Rect)

	gongsvgStage.Reset()
	gongsvgStage.Commit()

	var selectedDiagram *gongdoc_models.Classdiagram

	for diagramPackage := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.DiagramPackage](gongdocStage) {

		selectedDiagram = diagramPackage.SelectedClassdiagram
		if selectedDiagram == nil {
			return
		}
	}

	svg := new(gongsvg_models.SVG).Stage(gongsvgStage)

	for _, gongstructShape := range selectedDiagram.GongStructShapes {

		rectLayer := new(gongsvg_models.Layer).Stage(gongsvgStage)
		rectLayer.Name = "Layer" + gongstructShape.Identifier
		svg.Layers = append(svg.Layers, rectLayer)

		rect := new(gongsvg_models.Rect).Stage(gongsvgStage)
		rect.Name = gongstructShape.Identifier

		docSVGMapper.map_GongstructShape_Rect[gongstructShape] = rect
		docSVGMapper.map_Rect_GongstructShape[rect] = gongstructShape
		docSVGMapper.map_Structname_Rect[gongstructShape.Identifier] = rect

		rectLayer.Rects = append(rectLayer.Rects, rect)
		rect.X = gongstructShape.Position.X
		rect.Y = gongstructShape.Position.Y

		rect.Width = gongstructShape.Width
		rect.Height = gongstructShape.Heigth

		rect.Stroke = gongsvg_models.Lightsalmon.ToString()
		rect.StrokeWidth = 1
		rect.StrokeDashArrayWhenSelected = "5 5"

		rect.FillOpacity = 100
		rect.Color = gongsvg_models.Lightsalmon.ToString()

		// moveability
		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true

		// resizing
		rect.IsSelectable = true
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveTopHandle = true
		rect.CanHaveRightHandle = true

		//
		// Title
		//
		title := new(gongsvg_models.RectAnchoredText).Stage(gongsvgStage)
		title.Name = gongdoc_models.IdentifierToGongObjectName(gongstructShape.Identifier)
		title.Content = title.Name
		title.X_Offset = 0
		title.Y_Offset = 20
		title.RectAnchorType = gongsvg_models.RECT_ANCHOR_TOP
		title.TextAnchorType = gongsvg_models.TEXT_ANCHOR_CENTER
		title.FontWeight = "bold"
		title.Color = gongsvg_models.Black.ToString()
		title.FillOpacity = 1.0
		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)

		// additional box to hightlight the title
		titleBox := new(gongsvg_models.RectAnchoredRect).Stage(gongsvgStage)
		titleBox.Name = gongdoc_models.IdentifierToGongObjectName(gongstructShape.Identifier)
		titleBox.X_Offset = 0
		titleBox.Y_Offset = 0
		titleBox.Width = rect.Width
		titleBox.Height = 30
		titleBox.RectAnchorType = gongsvg_models.RECT_ANCHOR_TOP_LEFT
		titleBox.Color = "#ff8450"
		titleBox.WidthFollowRect = true
		titleBox.FillOpacity = 100

		rect.RectAnchoredRects = append(rect.RectAnchoredRects, titleBox)

		//
		// fields
		//
		for idx, field := range gongstructShape.Fields {
			fieldText := new(gongsvg_models.RectAnchoredText).Stage(gongsvgStage)
			fieldText.Name = field.Name + ":" + field.FieldTypeAsString
			fieldText.Content = field.Name

			// field position
			fieldText.X_Offset = 10
			fieldText.Y_Offset = 20 + 30 + float64(idx)*15
			fieldText.RectAnchorType = gongsvg_models.RECT_ANCHOR_TOP_LEFT
			fieldText.TextAnchorType = gongsvg_models.TEXT_ANCHOR_LEFT

			fieldText.Color = "black"
			fieldText.FillOpacity = 1.0
			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, fieldText)
		}

	}

	// display links
	for _, gongstructShape := range selectedDiagram.GongStructShapes {

		startRect := docSVGMapper.map_GongstructShape_Rect[gongstructShape]
		for _, docLink := range gongstructShape.Links {

			endRect := docSVGMapper.map_Structname_Rect[docLink.Fieldtypename]

			link := new(gongsvg_models.Link).Stage(gongsvgStage)
			link.Name = startRect.Name + " - to - " + endRect.Name
			linkLayer := new(gongsvg_models.Layer).Stage(gongsvgStage)

			linkLayer.Links = append(linkLayer.Links, link)
			svg.Layers = append(svg.Layers, linkLayer)

			// configuration
			link.Stroke = gongsvg_models.Black.ToString()
			link.StrokeWidth = 2
			link.HasEndArrow = true
			link.EndArrowSize = 8
			link.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL
			link.StartOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
			link.StartRatio = 0.5
			link.EndOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
			link.EndRatio = 0.5

			link.CornerOffsetRatio = (endRect.X - startRect.X - 80) / startRect.Width

			link.CornerRadius = 3

			link.Start = startRect
			link.End = endRect

			// add text to the arrow
			targetMulitplicity := new(gongsvg_models.AnchoredText).Stage(gongsvgStage)
			link.TextAtArrowEnd = append(link.TextAtArrowEnd, targetMulitplicity)
			targetMulitplicity.Name = docLink.TargetMultiplicity.ToString()
			targetMulitplicity.Content = targetMulitplicity.Name
			targetMulitplicity.Y_Offset = 16
			targetMulitplicity.X_Offset = -50
			targetMulitplicity.Stroke = gongsvg_models.Black.ToString()
			targetMulitplicity.StrokeWidth = 1

			fieldName := new(gongsvg_models.AnchoredText).Stage(gongsvgStage)
			link.TextAtArrowEnd = append(link.TextAtArrowEnd, fieldName)
			fieldName.Name = docLink.GetName()
			fieldName.Content = fieldName.Name
			fieldName.Y_Offset = -16
			fieldName.X_Offset = -50
			fieldName.Stroke = gongsvg_models.Black.ToString()
			fieldName.StrokeWidth = 1

			sourceMultiplicity := new(gongsvg_models.AnchoredText).Stage(gongsvgStage)
			link.TextAtArrowStart = append(link.TextAtArrowStart, sourceMultiplicity)
			sourceMultiplicity.Name = docLink.SourceMultiplicity.ToString()
			sourceMultiplicity.Content = sourceMultiplicity.Name
			sourceMultiplicity.Y_Offset = 10
			sourceMultiplicity.X_Offset = 10
			sourceMultiplicity.Stroke = gongsvg_models.Black.ToString()
			sourceMultiplicity.StrokeWidth = 1

		}
	}

	gongsvgStage.Commit()
}
