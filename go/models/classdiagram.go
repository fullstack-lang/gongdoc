package models

import (
	"math/rand"
)

// Classdiagram diagram struct store a class diagram
// temporary here
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// list of gongstructshapes in the diagram
	GongStructShapes []*GongStructShape

	GongEnumShapes []*GongEnumShape

	// list of notes in the diagram
	NoteShapes []*NoteShape

	// IsInDrawMode indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	IsInDrawMode bool
}

func (classdiagram *Classdiagram) RemoveGongStructShape(gongstructshapeName string) {

	foundGongStructShape := false
	var gongstructshape *GongStructShape
	for _, _gongstructshape := range classdiagram.GongStructShapes {

		// strange behavior when the gongstructshape is remove within the loop
		if IdentifierToGongStructName(_gongstructshape.Identifier) == gongstructshapeName && !foundGongStructShape {
			gongstructshape = _gongstructshape
		}
	}

	classdiagram.GongStructShapes = remove(classdiagram.GongStructShapes, gongstructshape)
	gongstructshape.Position.Unstage()
	gongstructshape.Unstage()

	// remove links that go from this gongstructshape
	for _, link := range gongstructshape.Links {
		link.Middlevertice.Unstage()
		link.Unstage()
	}
	gongstructshape.Links = []*Link{}

	// remove association links that go to this gongstructshape
	for _, fromGongStructShape := range classdiagram.GongStructShapes {

		newSliceOfLinks := make([]*Link, 0)
		for _, link := range fromGongStructShape.Links {
			if link.Fieldtypename == IdentifierToGongStructName(gongstructshape.Identifier) {
				link.Middlevertice.Unstage()
				link.Unstage()
			} else {
				newSliceOfLinks = append(newSliceOfLinks, link)
			}
		}
		fromGongStructShape.Links = newSliceOfLinks
	}

	// remove fields of the gongstructshape
	for _, field := range gongstructshape.Fields {
		field.Unstage()
	}

	//
	// remove documentation links that go this gongstructshape
	//
	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[NoteShape]().NoteShapeLinks[0].Name
	map_NoteShapeLink_NodeShape := GetSliceOfPointersReverseMap[NoteShape, NoteShapeLink](fieldName)
	for noteShapeLink := range *GetGongstructInstancesSet[NoteShapeLink]() {
		if noteShapeLink.Name == gongstructshapeName {

			// get the note shape
			noteShape := map_NoteShapeLink_NodeShape[noteShapeLink]

			// remove it from the slice of links
			noteShape.NoteShapeLinks = remove(noteShape.NoteShapeLinks, noteShapeLink)

			noteShapeLink.DeleteStageAndCommit()
		}
	}

	// log.Println("RemoveGongStructShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	Stage.Commit()
	// log.Println("RemoveGongStructShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
}

func (classdiagram *Classdiagram) AddGongStructShape(diagramPackage *DiagramPackage, gongstructshapeName string) {

	var gongstructshape GongStructShape
	gongstructshape.Name = classdiagram.Name + "-" + gongstructshapeName
	gongstructshape.Identifier = GongStructNameToIdentifier(gongstructshapeName)
	gongstructshape.Width = 240
	gongstructshape.Heigth = 63

	// attach GongStruct to gongstructshape
	nbInstances, ok := diagramPackage.Map_Identifier_NbInstances[gongstructshape.Identifier]
	if ok {
		gongstructshape.ShowNbInstances = true
		gongstructshape.NbInstances = nbInstances
	}
	gongstructshape.Stage()

	var position Position
	position.Name = "Pos-" + gongstructshape.Name
	position.X = float64(int(rand.Float32()*100) + 10)
	position.Y = float64(int(rand.Float32()*100) + 10)
	gongstructshape.Position = &position
	position.Stage()

	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, &gongstructshape)

	// log.Println("AddGongStructShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	Stage.Commit()
	// log.Println("AddGongStructShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())

}
