package models

import (
	"log"
	"strings"
)

const RefPrefixReferencedPackage = "ref_"
const RefPackagePlusPeriod = "models."

// ToFieldName take an ident in the forms
// "ref_models.Foo.Name" and returns "Name"
func ToFieldName(fieldIdentifier string) (fieldName string) {

	if !strings.Contains(fieldIdentifier, RefPackagePlusPeriod) {
		log.Fatalln("ToFieldName: missing", RefPackagePlusPeriod, "in", fieldIdentifier)
	}

	structNameWithFieldName := strings.TrimPrefix(fieldIdentifier, RefPrefixReferencedPackage+RefPackagePlusPeriod)

	subStrings := strings.Split(structNameWithFieldName, ".")
	if len(subStrings) != 2 {
		log.Fatalln("ToFieldName: wrong number of substrings in ", structNameWithFieldName)
	}

	fieldName = subStrings[1]

	return
}

// ToFieldIdentifier takes "Foo" "Name" and returns "ref_models.Foo.Name"
func ToFieldIdentifier(structName string, fieldName string) (fieldIdentifier string) {

	fieldIdentifier = RefPrefixReferencedPackage + RefPackagePlusPeriod +
		structName + "." + fieldName

	return
}

// ToStructName take an ident in the forms
// "ref_models.Foo" and returns "Foo"
func ToStructName(structIdentifier string) (structName string) {

	structName = strings.TrimPrefix(structIdentifier, RefPrefixReferencedPackage+"models.")

	return
}
