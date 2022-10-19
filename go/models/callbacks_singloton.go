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
	staged, new *Node) {

	switch staged.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:
		if staged.IsChecked != new.IsChecked {

			// setting the value of the staged node	to the new value
			staged.IsChecked = new.IsChecked

			// parse all nodes and check wether, if checked/not checked, the is the updated node
			diagramNodes := append(
				callbacksSingloton.ClassdiagramsRootNode.Children,
				callbacksSingloton.StateDiagramsRootNode.Children...)

			for _, node := range diagramNodes {
				if node == staged {
					continue
				}
				if new.IsChecked && node.IsChecked {
					// log.Println("Node " + node.Name + " is checked and should unchecked")
					node.IsChecked = false
					node.Commit()
				}
				if !new.IsChecked && !node.IsChecked {
					// log.Println("Node " + node.Name + " is checked and should unchecked")
					node.IsChecked = true
					node.Commit()
				}
			}
		}

		// in case the front change the name of the diagram
		// one need to indicate the change as an increase in the commit
		// from the back, otherwise, the front wont be able to detect
		// the change
		// change the name of the diagram
		if staged.Name != new.Name {

			switch staged.Type {
			case CLASS_DIAGRAM:
				staged.Classdiagram.Name = new.Name
				staged.Classdiagram.Commit()
			case STATE_DIAGRAM:
				staged.Umlsc.Name = new.Name
				staged.Umlsc.Commit()
			}
			switch staged.Type {
			case CLASS_DIAGRAM, STATE_DIAGRAM:
				staged.Name = new.Name
				staged.IsInEditMode = false
				staged.Commit()
			}
		}

	}

	if staged.IsExpanded != new.IsExpanded {
		log.Println("Node " + staged.Name + " is updated with value IsExpanded to " + strconv.FormatBool(new.IsExpanded))

		// setting the value of the staged node	to the new value
		staged.IsExpanded = new.IsExpanded
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
