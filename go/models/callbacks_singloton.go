package models

import (
	"log"
	"strconv"
)

type CallbacksSingloton struct {
	ClassdiagramsRootNode *Node
	StateDiagramsRootNode *Node
}

func (callbacksSingloton CallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// if a diagram is selected, you cannot unselect it
		if !stagedNode.IsChecked && frontNode.IsChecked {

			// setting the value of the staged node	to the new value
			stagedNode.IsChecked = frontNode.IsChecked
			stagedNode.Commit()

			// parse all nodes and uncheck them if necessary
			diagramNodes := append(
				callbacksSingloton.ClassdiagramsRootNode.Children,
				callbacksSingloton.StateDiagramsRootNode.Children...)

			for _, otherDiagramNode := range diagramNodes {
				if otherDiagramNode == stagedNode {
					continue
				}

				// uncheck the other node
				if otherDiagramNode.IsChecked {
					// log.Println("Node " + node.Name + " is checked and should be unchecked")
					otherDiagramNode.IsChecked = false
					otherDiagramNode.Commit()
				}
			}
		}

		// in case the front change the name of the diagram
		// one need to indicate the change as an increase in the commit
		// from the back, otherwise, the front wont be able to detect
		// the change
		// change the name of the diagram
		if stagedNode.Name != frontNode.Name {

			switch stagedNode.Type {
			case CLASS_DIAGRAM:
				stagedNode.Classdiagram.Name = frontNode.Name
				stagedNode.Classdiagram.Commit()
			case STATE_DIAGRAM:
				stagedNode.Umlsc.Name = frontNode.Name
				stagedNode.Umlsc.Commit()
			}
			switch stagedNode.Type {
			case CLASS_DIAGRAM, STATE_DIAGRAM:
				stagedNode.Name = frontNode.Name
				stagedNode.IsInEditMode = false
				stagedNode.Commit()
			}
		}

	}

	if stagedNode.IsExpanded != frontNode.IsExpanded {
		log.Println("Node " + stagedNode.Name + " is updated with value IsExpanded to " + strconv.FormatBool(frontNode.IsExpanded))

		// setting the value of the staged node	to the new value
		stagedNode.IsExpanded = frontNode.IsExpanded
	}
}

func (callbacksSingloton CallbacksSingloton) OnAfterCreate(
	stage *StageStruct,
	newDiagramNode *Node) {

	log.Println("Node " + newDiagramNode.Name + " is created")

	switch newDiagramNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:
		newDiagramNode.HasCheckboxButton = true

		// fetch the only package
		var pkgelt *Pkgelt
		for _pkgelt := range *GetGongstructInstancesSet[Pkgelt]() {
			pkgelt = _pkgelt
		}

		classdiagram := (&Classdiagram{Name: newDiagramNode.Name}).Stage()
		pkgelt.Classdiagrams = append(pkgelt.Classdiagrams, classdiagram)
		newDiagramNode.Classdiagram = classdiagram

		stage.Commit()
	}
}

func (callbacksSingloton CallbacksSingloton) OnAfterDelete(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

	}
}
