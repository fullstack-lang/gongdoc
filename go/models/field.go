package models

import (
	"log"
	"reflect"
	"strings"
)

// Field represent the UML Field of a Class (a "struct" in go)
// swagger:model Field
type Field struct {
	Name string

	// swagger:ignore
	// actual go field in the models that is diagrammed, for instance "models.Point{}.X"
	Field interface{} `gorm:"-"`

	Fieldname string

	// Identifier is the identifier of the struct field referenced
	// by the UML classshape in the modeled package
	//gong:ident
	Identifier string

	FieldTypeAsString string
	Structname        string
	Fieldtypename     string
}

// serialize the package and its elements to the Stage
// this is used if one Umlsc is dynamicaly created
func (field *Field) SerializeToStage() {

	field.Stage()

	// update name if not done
	if field.Name == "" {
		if field.Field != nil {
			// when a field can be is a struct
			typeofstruct := reflect.TypeOf(field.Field)

			if typeofstruct.Kind() == reflect.Ptr {
				log.Printf(typeofstruct.Elem().String())
				field.Name = strings.Split(typeofstruct.Elem().String(), ".")[1]
				field.Structname = field.Name
			}
		}
	}
}
