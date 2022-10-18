package models

import (
	"log"
)

func FillUpNodeTree(pkgelt *Pkgelt) {

	// FillUpCodeNodeTree(pkgelt)
	FillUpDiagramNodeTree(pkgelt)

	Stage.Commit()
	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))
	log.Printf("Server ready to serve on http://localhost:8080/")
}
