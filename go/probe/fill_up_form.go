// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("GongStructShapes", instanceWithInferedType, &instanceWithInferedType.GongStructShapes, formGroup, playground)
		AssociationSliceToForm("GongEnumShapes", instanceWithInferedType, &instanceWithInferedType.GongEnumShapes, formGroup, playground)
		AssociationSliceToForm("NoteShapes", instanceWithInferedType, &instanceWithInferedType.NoteShapes, formGroup, playground)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DiagramPackage"
			rf.Fieldname = "Classdiagrams"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DiagramPackage),
					"Classdiagrams",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.DiagramPackage, *models.Classdiagram](
					nil,
					"Classdiagrams",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.DiagramPackage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("GongModelPath", instanceWithInferedType.GongModelPath, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Classdiagrams", instanceWithInferedType, &instanceWithInferedType.Classdiagrams, formGroup, playground)
		AssociationFieldToForm("SelectedClassdiagram", instanceWithInferedType.SelectedClassdiagram, formGroup, playground)
		AssociationSliceToForm("Umlscs", instanceWithInferedType, &instanceWithInferedType.Umlscs, formGroup, playground)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsReloaded", instanceWithInferedType.IsReloaded, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("AbsolutePathToDiagramPackage", instanceWithInferedType.AbsolutePathToDiagramPackage, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Field:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FieldTypeAsString", instanceWithInferedType.FieldTypeAsString, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Structname", instanceWithInferedType.Structname, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "Fields"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStructShape),
					"Fields",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongStructShape, *models.Field](
					nil,
					"Fields",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.GongEnumShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, playground)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("GongEnumValueEntrys", instanceWithInferedType, &instanceWithInferedType.GongEnumValueEntrys, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongEnumShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Classdiagram),
					"GongEnumShapes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Classdiagram, *models.GongEnumShape](
					nil,
					"GongEnumShapes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.GongEnumValueEntry:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnumShape"
			rf.Fieldname = "GongEnumValueEntrys"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongEnumShape),
					"GongEnumValueEntrys",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongEnumShape, *models.GongEnumValueEntry](
					nil,
					"GongEnumValueEntrys",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.GongStructShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, playground)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ShowNbInstances", instanceWithInferedType.ShowNbInstances, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("NbInstances", instanceWithInferedType.NbInstances, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Fields", instanceWithInferedType, &instanceWithInferedType.Fields, formGroup, playground)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongStructShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Classdiagram),
					"GongStructShapes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Classdiagram, *models.GongStructShape](
					nil,
					"GongStructShapes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FieldOffsetX", instanceWithInferedType.FieldOffsetX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("FieldOffsetY", instanceWithInferedType.FieldOffsetY, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("TargetMultiplicity", instanceWithInferedType.TargetMultiplicity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetX", instanceWithInferedType.TargetMultiplicityOffsetX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("TargetMultiplicityOffsetY", instanceWithInferedType.TargetMultiplicityOffsetY, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("SourceMultiplicity", instanceWithInferedType.SourceMultiplicity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetX", instanceWithInferedType.SourceMultiplicityOffsetX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("SourceMultiplicityOffsetY", instanceWithInferedType.SourceMultiplicityOffsetY, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Middlevertice", instanceWithInferedType.Middlevertice, formGroup, playground)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "Links"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStructShape),
					"Links",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.GongStructShape, *models.Link](
					nil,
					"Links",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Matched", instanceWithInferedType.Matched, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("NoteShapeLinks", instanceWithInferedType, &instanceWithInferedType.NoteShapeLinks, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "NoteShapes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Classdiagram),
					"NoteShapes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Classdiagram, *models.NoteShape](
					nil,
					"NoteShapes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.NoteShapeLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, playground.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "NoteShape"
			rf.Fieldname = "NoteShapeLinks"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.NoteShape),
					"NoteShapeLinks",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.NoteShape, *models.NoteShapeLink](
					nil,
					"NoteShapeLinks",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Position:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.UmlState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Umlsc"
			rf.Fieldname = "States"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Umlsc),
					"States",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Umlsc, *models.UmlState](
					nil,
					"States",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Umlsc:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, formGroup, playground)
		BasicFieldtoForm("Activestate", instanceWithInferedType.Activestate, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DiagramPackage"
			rf.Fieldname = "Umlscs"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DiagramPackage),
					"Umlscs",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.DiagramPackage, *models.Umlsc](
					nil,
					"Umlscs",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Vertice:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)

	default:
		_ = instanceWithInferedType
	}
}
