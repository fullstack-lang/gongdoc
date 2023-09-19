// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GongStructShapes", instanceWithInferedType, &instanceWithInferedType.GongStructShapes, formGroup, playground)
		AssociationSliceToForm("GongEnumShapes", instanceWithInferedType, &instanceWithInferedType.GongEnumShapes, formGroup, playground)
		AssociationSliceToForm("NoteShapes", instanceWithInferedType, &instanceWithInferedType.NoteShapes, formGroup, playground)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, playground.formStage, formGroup)

	case *models.DiagramPackage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("GongModelPath", instanceWithInferedType.GongModelPath, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Classdiagrams", instanceWithInferedType, &instanceWithInferedType.Classdiagrams, formGroup, playground)
		AssociationFieldToForm("SelectedClassdiagram", instanceWithInferedType.SelectedClassdiagram, formGroup, playground)
		AssociationSliceToForm("Umlscs", instanceWithInferedType, &instanceWithInferedType.Umlscs, formGroup, playground)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsReloaded", instanceWithInferedType.IsReloaded, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("AbsolutePathToDiagramPackage", instanceWithInferedType.AbsolutePathToDiagramPackage, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Field:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FieldTypeAsString", instanceWithInferedType.FieldTypeAsString, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Structname", instanceWithInferedType.Structname, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongEnumShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, playground)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("GongEnumValueEntrys", instanceWithInferedType, &instanceWithInferedType.GongEnumValueEntrys, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongEnumValueEntry:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)

	case *models.GongStructShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, playground)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("ShowNbInstances", instanceWithInferedType.ShowNbInstances, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("NbInstances", instanceWithInferedType.NbInstances, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Fields", instanceWithInferedType, &instanceWithInferedType.Fields, formGroup, playground)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FieldOffsetX", instanceWithInferedType.FieldOffsetX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("FieldOffsetY", instanceWithInferedType.FieldOffsetY, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("TargetMultiplicity", instanceWithInferedType.TargetMultiplicity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetX", instanceWithInferedType.TargetMultiplicityOffsetX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetY", instanceWithInferedType.TargetMultiplicityOffsetY, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("SourceMultiplicity", instanceWithInferedType.SourceMultiplicity, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetX", instanceWithInferedType.SourceMultiplicityOffsetX, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetY", instanceWithInferedType.SourceMultiplicityOffsetY, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Middlevertice", instanceWithInferedType.Middlevertice, formGroup, playground)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, playground.formStage, formGroup)

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Matched", instanceWithInferedType.Matched, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("NoteShapeLinks", instanceWithInferedType, &instanceWithInferedType.NoteShapeLinks, formGroup, playground)

	case *models.NoteShapeLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Position:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	case *models.UmlState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Umlsc:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, formGroup, playground)
		BasicFieldtoForm("Activestate", instanceWithInferedType.Activestate, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Vertice:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}
