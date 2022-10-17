package models

import (
	"log"
)

func FillUpDiagramNodeTree(pkgelt *Pkgelt) {

	// generate tree of diagrams
	gongdocTree := (&Tree{Name: "gongdoc", Type: TREE_OF_DIAGRAMS}).Stage()

	// add the root of class diagrams
	classdiagramsRootNode := (&Node{Name: "class diagrams", Type: ROOT_OF_CLASS_DIAGRAMS}).Stage()
	classdiagramsRootNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, classdiagramsRootNode)

	// add one node per class diagram
	for classdiagram := range *GetGongstructInstancesSet[Classdiagram]() {
		node := (&Node{Name: classdiagram.Name}).Stage()
		node.Classdiagram = classdiagram
		node.Type = CLASS_DIAGRAM
		node.HasCheckboxButton = true
		classdiagramsRootNode.Children = append(classdiagramsRootNode.Children, node)
	}

	// root of state diagrams
	stateDiagramssRootNode := (&Node{Name: "state diagrams", Type: ROOT_OF_STATE_DIAGRAMS}).Stage()
	stateDiagramssRootNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, stateDiagramssRootNode)

	// add one node per state diagram
	for statediagram := range *GetGongstructInstancesSet[Umlsc]() {
		node := (&Node{Name: statediagram.Name}).Stage()
		node.Umlsc = statediagram
		node.Type = STATE_DIAGRAM
		node.HasCheckboxButton = true
		stateDiagramssRootNode.Children = append(stateDiagramssRootNode.Children, node)
	}

	// set callbacks on node updates
	onNodeCallbackStruct := new(CallbacksSingloton)
	onNodeCallbackStruct.ClassdiagramsRootNode = classdiagramsRootNode
	onNodeCallbackStruct.StateDiagramsRootNode = stateDiagramssRootNode
	Stage.OnAfterNodeUpdateCallback = onNodeCallbackStruct
}

func FillUpNodeTree(pkgelt *Pkgelt) {

	FillUpCodeNodeTree(pkgelt)
	FillUpDiagramNodeTree(pkgelt)

	Stage.Commit()
	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))
	log.Printf("Server ready to serve on http://localhost:8080/")
}
