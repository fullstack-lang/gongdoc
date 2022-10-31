package models

import (
	"log"
)

func FillUpNodeTree(pkgelt *DiagramPackage) {

	onNodeCallbackStruct := new(CallbacksSingloton)
	FillUpDiagramNodeTree(pkgelt, onNodeCallbackStruct)
	FillUpTreeOfIdentifiers(pkgelt, onNodeCallbackStruct)
	updateNodesStates(&Stage, onNodeCallbackStruct)

	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))
	log.Printf("Server ready to serve on http://localhost:8080/")
}
