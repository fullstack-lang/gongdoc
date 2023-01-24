package models

import (
	"log"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

func (nodesCb *NodeCB) OnAfterUpdateStructField(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	classdiagram := nodesCb.GetSelectedClassdiagram()

	// find the parent node to find the gongstruct to find the classshape
	// the node is field, one needs to find the gongstruct that contains it
	// get the parent node
	parentNode := nodesCb.map_Children_Parent[stagedNode]
	gongStruct := parentNode.Gongstruct

	// find the classhape in the classdiagram
	foundClassshape := false
	var classshape *Classshape
	for _, _classshape := range classdiagram.Classshapes {
		// strange behavior when the classshape is remove within the loop
		if IdentifierToShapename(_classshape.Identifier) ==
			gongStruct.Name && !foundClassshape {
			classshape = _classshape
		}
	}

	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		{
			var field *Field

			for _, _field := range classshape.Fields {
				if IdentifierToFieldName(_field.Identifier) == stagedNode.Name {
					field = _field
				}
			}
			if field != nil {
				classshape.Fields = remove(classshape.Fields, field)
				classshape.Heigth = classshape.Heigth - 15
				field.Unstage()
			}
		}
		{
			var link *Link

			for _, _field := range classshape.Links {
				if IdentifierToFieldName(IdentifierToFieldName(_field.Identifier)) == stagedNode.Name {
					link = _field
				}
			}
			if link != nil {
				classshape.Links = remove(classshape.Links, link)
				link.Unstage()
			}
		}

		Stage.Commit()
	}
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		switch stagedNode.Gongfield.(type) {
		case *gong_models.GongBasicField, *gong_models.GongTimeField:

			var field Field
			field.Name = stagedNode.Name
			field.Identifier = ShapeAndFieldnameToFieldIdentifier(gongStruct.Name, stagedNode.Name)

			switch realField := stagedNode.Gongfield.(type) {
			case *gong_models.GongBasicField:

				// get the type after the "."
				names := strings.Split(realField.DeclaredType, ".")
				fieldTypeName := names[len(names)-1]

				field.Fieldtypename = fieldTypeName
			case *gong_models.GongTimeField:
				field.Fieldtypename = "Time"
			case *gong_models.PointerToGongStructField:
			case *gong_models.SliceOfPointerToGongStructField:
			}

			field.Structname = IdentifierToShapename(classshape.Identifier)
			field.Stage()

			classshape.Heigth = classshape.Heigth + 15

			// construct ordered slice of fields
			map_Field_Rank := make(map[gong_models.FieldInterface]int, 0)
			map_Name_Field := make(map[string]gong_models.FieldInterface, 0)

			// what is the index of the field to insert in the gong struct ?
			fieldRank := 0

			// let's compute it by parsing the field of the gongstruct
			gongStruct_ := (*gong_models.GetGongstructInstancesMap[gong_models.GongStruct]())[gongStruct.Name]
			for idx, gongField := range gongStruct_.Fields {

				map_Field_Rank[gongField] = idx
				map_Name_Field[gongField.GetName()] = gongField

				if gongField.GetName() == field.Name {
					fieldRank = idx
				}
			}

			// compute insertionIndex (index where to insert the field to display)
			insertionIndex := 0
			for idx, field := range classshape.Fields {
				gongField := map_Name_Field[IdentifierToFieldName(field.Identifier)]
				_fieldRank := map_Field_Rank[gongField]
				if fieldRank > _fieldRank {
					insertionIndex = idx + 1
				}
			}

			// append the filed to display in the right index
			if insertionIndex == len(classshape.Fields) {
				classshape.Fields = append(classshape.Fields, &field)
			} else {
				classshape.Fields = append(classshape.Fields[:insertionIndex+1],
					classshape.Fields[insertionIndex:]...)
				classshape.Fields[insertionIndex] = &field
			}

		case *gong_models.PointerToGongStructField, *gong_models.SliceOfPointerToGongStructField:

			var targetStructName string
			var sourceMultiplicity MultiplicityType
			var targetMultiplicity MultiplicityType

			switch realField := stagedNode.Gongfield.(type) {
			case *gong_models.PointerToGongStructField:
				targetStructName = realField.GongStruct.Name
				sourceMultiplicity = MANY
				targetMultiplicity = ZERO_ONE
			case *gong_models.SliceOfPointerToGongStructField:
				targetStructName = realField.GongStruct.Name
				sourceMultiplicity = ZERO_ONE
				targetMultiplicity = MANY
			}
			targetSourceClassshape := false
			var targetClassshape *Classshape
			for _, _classshape := range classdiagram.Classshapes {

				// strange behavior when the classshape is remove within the loop
				if IdentifierToShapename(_classshape.Identifier) == targetStructName && !targetSourceClassshape {
					targetSourceClassshape = true
					targetClassshape = _classshape
				}
			}
			if !targetSourceClassshape {
				log.Panicf("Classshape %s of field not present ", targetStructName)
			}
			_ = targetClassshape

			link := new(Link).Stage()
			link.Name = stagedNode.Name
			link.SourceMultiplicity = sourceMultiplicity
			link.TargetMultiplicity = targetMultiplicity
			link.Identifier =
				ShapeAndFieldnameToFieldIdentifier(gongStruct.Name, stagedNode.Name)

			link.Structname = gongStruct.Name
			link.Fieldtypename = targetStructName

			classshape.Links = append(classshape.Links, link)
			link.Middlevertice = new(Vertice).Stage()
			link.Middlevertice.Name = "Verticle in class diagram " + classdiagram.Name +
				" in middle between " + classshape.Name + " and " + targetClassshape.Name
			link.Middlevertice.X = (classshape.Position.X+targetClassshape.Position.X)/2.0 +
				classshape.Width*1.5
			link.Middlevertice.Y = (classshape.Position.Y+targetClassshape.Position.Y)/2.0 +
				classshape.Heigth/2.0
			Stage.Commit()
		}

		Stage.Commit()
	}
}
