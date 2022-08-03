package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case Classdiagram:
		fieldCoder := Classdiagram{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.IsEditable = false
		return (any)(fieldCoder).(Type)
	case Classshape:
		fieldCoder := Classshape{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Structname = "2"
		fieldCoder.ShowNbInstances = false
		fieldCoder.NbInstances = 5
		fieldCoder.Width = 8.000000
		fieldCoder.Heigth = 9.000000
		fieldCoder.ClassshapeTargetType = "10"
		return (any)(fieldCoder).(Type)
	case Field:
		fieldCoder := Field{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Fieldname = "1"
		fieldCoder.FieldTypeAsString = "2"
		fieldCoder.Structname = "3"
		fieldCoder.Fieldtypename = "4"
		return (any)(fieldCoder).(Type)
	case GongStruct:
		fieldCoder := GongStruct{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.NbInstances = 1
		return (any)(fieldCoder).(Type)
	case GongdocCommand:
		fieldCoder := GongdocCommand{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Command = "1"
		fieldCoder.DiagramName = "2"
		fieldCoder.Date = "3"
		fieldCoder.GongdocNodeType = "4"
		fieldCoder.StructName = "5"
		fieldCoder.FieldName = "6"
		fieldCoder.FieldTypeName = "7"
		fieldCoder.PositionX = 8
		fieldCoder.PositionY = 9
		fieldCoder.NoteName = "10"
		return (any)(fieldCoder).(Type)
	case GongdocStatus:
		fieldCoder := GongdocStatus{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Status = "1"
		fieldCoder.CommandCompletionDate = "2"
		return (any)(fieldCoder).(Type)
	case Link:
		fieldCoder := Link{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Fieldname = "1"
		fieldCoder.Structname = "2"
		fieldCoder.Fieldtypename = "3"
		fieldCoder.TargetMultiplicity = "4"
		fieldCoder.SourceMultiplicity = "5"
		return (any)(fieldCoder).(Type)
	case Note:
		fieldCoder := Note{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Body = "1"
		fieldCoder.X = 2.000000
		fieldCoder.Y = 3.000000
		fieldCoder.Width = 4.000000
		fieldCoder.Heigth = 5.000000
		return (any)(fieldCoder).(Type)
	case Pkgelt:
		fieldCoder := Pkgelt{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Path = "1"
		fieldCoder.GongModelPath = "2"
		fieldCoder.Editable = false
		return (any)(fieldCoder).(Type)
	case Position:
		fieldCoder := Position{}
		// insertion point for field dependant code
		fieldCoder.X = 0.000000
		fieldCoder.Y = 1.000000
		fieldCoder.Name = "2"
		return (any)(fieldCoder).(Type)
	case UmlState:
		fieldCoder := UmlState{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.X = 1.000000
		fieldCoder.Y = 2.000000
		return (any)(fieldCoder).(Type)
	case Umlsc:
		fieldCoder := Umlsc{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Activestate = "2"
		return (any)(fieldCoder).(Type)
	case Vertice:
		fieldCoder := Vertice{}
		// insertion point for field dependant code
		fieldCoder.X = 0.000000
		fieldCoder.Y = 1.000000
		fieldCoder.Name = "2"
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Classdiagram | []*Classdiagram | *Classshape | []*Classshape | *Field | []*Field | *GongStruct | []*GongStruct | *GongdocCommand | []*GongdocCommand | *GongdocStatus | []*GongdocStatus | *Link | []*Link | *Note | []*Note | *Pkgelt | []*Pkgelt | *Position | []*Position | *UmlState | []*UmlState | *Umlsc | []*Umlsc | *Vertice | []*Vertice
}

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case *Classdiagram:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "IsEditable"
			}
		}
	case *Classshape:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "2" {
				return "Structname"
			}
			if field == "10" {
				return "ClassshapeTargetType"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 5 {
				return "NbInstances"
			}
		case float64:
			// insertion point for field dependant name
			if field == 8.000000 {
				return "Width"
			}
			if field == 9.000000 {
				return "Heigth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "ShowNbInstances"
			}
		}
	case *Field:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Fieldname"
			}
			if field == "2" {
				return "FieldTypeAsString"
			}
			if field == "3" {
				return "Structname"
			}
			if field == "4" {
				return "Fieldtypename"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *GongStruct:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "NbInstances"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *GongdocCommand:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Command"
			}
			if field == "2" {
				return "DiagramName"
			}
			if field == "3" {
				return "Date"
			}
			if field == "4" {
				return "GongdocNodeType"
			}
			if field == "5" {
				return "StructName"
			}
			if field == "6" {
				return "FieldName"
			}
			if field == "7" {
				return "FieldTypeName"
			}
			if field == "10" {
				return "NoteName"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 8 {
				return "PositionX"
			}
			if field == 9 {
				return "PositionY"
			}
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *GongdocStatus:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Status"
			}
			if field == "2" {
				return "CommandCompletionDate"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Link:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Fieldname"
			}
			if field == "2" {
				return "Structname"
			}
			if field == "3" {
				return "Fieldtypename"
			}
			if field == "4" {
				return "TargetMultiplicity"
			}
			if field == "5" {
				return "SourceMultiplicity"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Note:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Body"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 2.000000 {
				return "X"
			}
			if field == 3.000000 {
				return "Y"
			}
			if field == 4.000000 {
				return "Width"
			}
			if field == 5.000000 {
				return "Heigth"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Pkgelt:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "1" {
				return "Path"
			}
			if field == "2" {
				return "GongModelPath"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
			if field == false {
				return "Editable"
			}
		}
	case *Position:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "2" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 0.000000 {
				return "X"
			}
			if field == 1.000000 {
				return "Y"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *UmlState:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 1.000000 {
				return "X"
			}
			if field == 2.000000 {
				return "Y"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Umlsc:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "2" {
				return "Activestate"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *Vertice:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "2" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
		case float64:
			// insertion point for field dependant name
			if field == 0.000000 {
				return "X"
			}
			if field == 1.000000 {
				return "Y"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
