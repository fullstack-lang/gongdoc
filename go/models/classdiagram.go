package models

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"path/filepath"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
	"github.com/tdewolff/canvas/renderers/svg"

	"github.com/fullstack-lang/gongdoc/go/static"
)

// Classdiagram diagram struct store a class diagram
// temporary here
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// list of classshapes in the diagram
	Classshapes []*Classshape

	// list of notes in the diagram
	Notes []*NoteShape

	// IsInDrawMode indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	IsInDrawMode bool
}

const DiagramMarginX = 10.0
const DiagramMarginY = 10.0

// Extent compute max X and max Y
func (classdiagram *Classdiagram) Extent() (x, y, maxClassshapeHeigth float64) {

	for _, classshape := range classdiagram.Classshapes {

		if maxClassshapeHeigth < classshape.Heigth {
			maxClassshapeHeigth = classshape.Heigth
		}

		if classshape.Position.X+classshape.Width > x {
			x = classshape.Position.X + classshape.Width
		}
		if classshape.Position.Y+classshape.Heigth > y {
			y = classshape.Position.Y + classshape.Heigth
		}

		for _, link := range classshape.Links {
			if link.Middlevertice.X > x {
				x = link.Middlevertice.X
			}
			if link.Middlevertice.Y > y {
				y = link.Middlevertice.Y
			}
		}
	}
	// margin
	x += DiagramMarginX
	y += DiagramMarginY
	return
}

// ClassdiagramMap is a Map of all Classdiagram via their Name
type ClassdiagramMap map[string]*Classdiagram

// ClassdiagramStore is a handy ClassdiagramMap
var ClassdiagramStore ClassdiagramMap = make(map[string]*Classdiagram, 0)

// ModelSpaceToSVGPaperYCoord convert the Y coordinates from the model space
// to the SVG space
// in model space, Y coordinates are descending. In SVG, Y coordinates are
// ascending.
// You need to have the paper heigth mind
var SVGPaperHeight float64

// MaxClassshapeHeigth because you also need the max heigth of all classshapes
var MaxClassshapeHeigth float64

// ModelToSVGYCoord converts a Y coordinate in the model to a Y coordinate in the SVG canvas
func ModelToSVGYCoord(yModel float64) (ySVG float64) {
	return SVGPaperHeight - yModel
}

// ModelToSVGRectangleYOrigin converts a top left Y coordinate of a rectangle in the model to a
// bottom left Y coordinate in the SVG canvas
func ModelToSVGRectangleYOrigin(yModel, classshapeHeigth float64) (ySVG float64) {
	return ModelToSVGYCoord(yModel + classshapeHeigth)
}

func (classdiagram *Classdiagram) OutputSVG(path string) {

	var maxx, maxy float64
	maxx, maxy, MaxClassshapeHeigth = classdiagram.Extent()
	var width float64
	width, SVGPaperHeight = maxx, maxy

	c := canvas.New(width, SVGPaperHeight)
	ctx := canvas.NewContext(c)
	ctx.SetStrokeColor(color.Black)

	mapStringClassshape := make(map[string]*Classshape)
	for _, classshape := range classdiagram.Classshapes {
		mapStringClassshape[IdentifierToShape(classshape.Identifier)] = classshape
	}

	dejaVuSerif := canvas.NewFontFamily("dejavu-serif")
	staticFontFile := []byte(static.Files["../gongdoc/DejaVuSerif.ttf"])
	if err := dejaVuSerif.LoadFont(
		staticFontFile,
		0,
		canvas.FontRegular); err != nil { // TTF, OTF, WOFF, or WOFF2
		log.Panic("cannot load font", err.Error())
	}
	ff := dejaVuSerif.Face(24.0, color.Black, canvas.FontBlack, canvas.FontNormal)

	// First, draw all links with full transparency
	ctx.SetFillColor(color.Transparent)
	for _, classshape := range classdiagram.Classshapes {

		bottomLeftX := classshape.Position.X
		bottomLeftY := ModelToSVGRectangleYOrigin(classshape.Position.Y, classshape.Heigth)

		for _, link := range classshape.Links {
			// tech target point of link

			targetClassshape := mapStringClassshape[link.Fieldtypename]

			// sometimes, links to not refer to classhape that are existing on the diagram (for instance Interfaces)
			if targetClassshape == nil {
				continue
			}
			targetPosition := targetClassshape.Position

			// from center to center
			sourceClassshapeToMiddleVerticePath := &canvas.Path{}

			var verticeSVGX = link.Middlevertice.X
			var verticeSVGY = ModelToSVGYCoord(link.Middlevertice.Y)

			var centerOfClassshapeX = bottomLeftX + classshape.Width/2.0
			var centerOfClassshapeY = bottomLeftY + classshape.Heigth/2.0

			// draw with middle vertice
			linkToMiddleVerticePathInDiagram := Position{
				X: verticeSVGX - centerOfClassshapeX,
				Y: verticeSVGY - centerOfClassshapeY,
			}
			sourceClassshapeToMiddleVerticePath.LineTo(linkToMiddleVerticePathInDiagram.X, linkToMiddleVerticePathInDiagram.Y)

			ctx.DrawPath(centerOfClassshapeX, centerOfClassshapeY, sourceClassshapeToMiddleVerticePath)

			targetClassshapeToMiddleVerticePath := &canvas.Path{}

			targetBottomLeftX := targetPosition.X
			targetBottomLeftY := ModelToSVGRectangleYOrigin(targetPosition.Y, classshape.Heigth)

			centerOfClassshapeX = targetBottomLeftX + classshape.Width/2.0
			centerOfClassshapeY = targetBottomLeftY + classshape.Heigth/2.0

			// draw with middle vertice
			linkToMiddleVerticePathInDiagram = Position{
				X: verticeSVGX - centerOfClassshapeX,
				Y: verticeSVGY - centerOfClassshapeY,
			}
			targetClassshapeToMiddleVerticePath.LineTo(linkToMiddleVerticePathInDiagram.X, linkToMiddleVerticePathInDiagram.Y)

			ctx.DrawPath(centerOfClassshapeX, centerOfClassshapeY, targetClassshapeToMiddleVerticePath)

			// draw text between middle vertice & target position
			text := canvas.NewTextLine(ff, IdentifierToFieldName(link.Identifier), canvas.TextAlign(canvas.Left)) // simple text line
			linkFieldPosition := Position{
				X: (verticeSVGX + centerOfClassshapeX) / 2.0,
				Y: (verticeSVGY + centerOfClassshapeY) / 2.0,
			}
			ctx.DrawText(linkFieldPosition.X, linkFieldPosition.Y, text)
		}
	}

	// draw the classshape AFTER the links
	ctx.SetFillColor(color.RGBA{204, 224, 218, 255})

	for _, classshape := range classdiagram.Classshapes {

		heigthBetweenLines := 16.0

		bottomLeftX := classshape.Position.X
		bottomLeftY := ModelToSVGRectangleYOrigin(classshape.Position.Y, classshape.Heigth)

		text := canvas.NewTextLine(ff,
			IdentifierToShape(classshape.Identifier),
			canvas.TextAlign(canvas.Left)) // simple text line
		p := canvas.Rectangle(classshape.Width, classshape.Heigth)

		// draw the line below the name of the class
		p.MoveTo(0, classshape.Heigth-heigthBetweenLines)
		p.LineTo(classshape.Width, classshape.Heigth-heigthBetweenLines)
		ctx.DrawPath(bottomLeftX, bottomLeftY, p)
		// draw the name of the class
		ctx.DrawText(bottomLeftX+10, bottomLeftY+classshape.Heigth-heigthBetweenLines+6, text)

		// draw fields name
		for i, field := range classshape.Fields {
			text := canvas.NewTextLine(ff,
				IdentifierToFieldName(field.Identifier)+
					":"+field.Fieldtypename,
				canvas.TextAlign(canvas.Left)) // simple text line
			ctx.DrawText(bottomLeftX+10.0, bottomLeftY+classshape.Heigth-12.0-float64(i+1)*heigthBetweenLines, text)
		}
	}

	c.WriteFile(fmt.Sprintf(filepath.Join(path, "%s.png"), classdiagram.Name), rasterizer.PNGWriter(5.0))
	c.WriteFile(fmt.Sprintf(filepath.Join(path, "%s.svg"), classdiagram.Name), svg.Writer)
}

func (classdiagram *Classdiagram) RemoveClassshape(classshapeName string) {

	foundClassshape := false
	var classshape *Classshape
	for _, _classshape := range classdiagram.Classshapes {

		// strange behavior when the classshape is remove within the loop
		if IdentifierToShape(_classshape.Identifier) == classshapeName && !foundClassshape {
			classshape = _classshape
		}
	}

	classdiagram.Classshapes = remove(classdiagram.Classshapes, classshape)
	classshape.Position.Unstage()
	classshape.Unstage()

	// remove links that go from this classshape
	for _, link := range classshape.Links {
		link.Middlevertice.Unstage()
		link.Unstage()
	}
	classshape.Links = []*Link{}

	// remove links that go to this classshape
	for _, fromClassshape := range classdiagram.Classshapes {

		newSliceOfLinks := make([]*Link, 0)
		for _, link := range fromClassshape.Links {
			if link.Fieldtypename == IdentifierToShape(classshape.Identifier) {
				link.Middlevertice.Unstage()
				link.Unstage()
			} else {
				newSliceOfLinks = append(newSliceOfLinks, link)
			}
		}
		fromClassshape.Links = newSliceOfLinks
	}

	// remove fields of the classshape
	for _, field := range classshape.Fields {
		field.Unstage()
	}

	// log.Println("RemoveClassshape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	Stage.Commit()
	// log.Println("RemoveClassshape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
}

func (classdiagram *Classdiagram) AddClassshape(nodesCb *NodeCallbacksSingloton, classshapeName string, referenceType ReferenceType) {

	var classshape Classshape
	classshape.Name = classdiagram.Name + "-" + classshapeName
	classshape.Identifier = ShapenameToIdentifier(classshapeName)
	classshape.Width = 240
	classshape.Heigth = 63

	// attach GongStruct to classshape
	nbInstances, ok := Map_Identifier_NbInstances[classshape.Identifier]
	if ok {
		classshape.ShowNbInstances = true
		classshape.NbInstances = nbInstances
	}
	classshape.Stage()

	var position Position
	position.Name = "Pos-" + classshape.Name
	position.X = float64(int(rand.Float32()*100) + 10)
	position.Y = float64(int(rand.Float32()*100) + 10)
	classshape.Position = &position
	position.Stage()

	classdiagram.Classshapes = append(classdiagram.Classshapes, &classshape)

	// log.Println("AddClassshape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	Stage.Commit()
	// log.Println("AddClassshape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())

}
