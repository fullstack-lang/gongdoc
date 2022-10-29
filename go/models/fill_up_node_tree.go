package models

import (
	"log"
)

func FillUpNodeTree(pkgelt *Pkgelt) {

	FillUpDiagramNodeTree(pkgelt)

	FillUpCodeNodeTree(pkgelt)

	Stage.Commit()
	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))
	log.Printf("Server ready to serve on http://localhost:8080/")
}
