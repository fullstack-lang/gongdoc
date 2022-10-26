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
			stagedNode.IsChecked = true
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

		// node was checked and user wants to uncheck it. This is not possible
		// from a application logic point of view
		// on need to commit the staged node for the front to reconstruct
		// the node as checked and overides the unchecking action
		if stagedNode.IsChecked && !frontNode.IsChecked {
			stagedNode.Commit()
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

// remove node from slice
func remove[T Gongstruct](slice []*T, t *T) []*T {

	// get the index of the t
	rank := -1
	for i, t_ := range slice {
		if t_ == t {
			rank = i
		}
	}
	return append(slice[:rank], slice[rank+1:]...)
}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (callbacksSingloton CallbacksSingloton) OnAfterDelete(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// checkout the stage, it shall remove the link between
		// the parent node and the staged node because 0..1->0..N association
		// is stored in the staged node as a reverse pointer
		stage.Checkout()

		// remove the child node from the parent node
		// fieldName := GetAssociationName[Node]().Children[0].Name
		// mapReverse := GetSliceOfPointersReverseMap[Node, Node](fieldName)
		// parentNode := mapReverse[stagedNode]
		// parentNode.Children = remove(parentNode.Children, stagedNode)
	}
	switch stagedNode.Type {
	case CLASS_DIAGRAM:
		// remove the classdiagram node from the pkg element node
		fieldName := GetAssociationName[Pkgelt]().Classdiagrams[0].Name
		mapReverse := GetSliceOfPointersReverseMap[Pkgelt, Classdiagram](fieldName)
		pkgelt := mapReverse[stagedNode.Classdiagram]
		pkgelt.Classdiagrams = remove(pkgelt.Classdiagrams, stagedNode.Classdiagram)
		stagedNode.Classdiagram.Unstage()
	case STATE_DIAGRAM:

		// remove the umlsc node from the pkg element node
		fieldName := GetAssociationName[Pkgelt]().Umlscs[0].Name
		mapReverse := GetSliceOfPointersReverseMap[Pkgelt, Umlsc](fieldName)
		pkgelt := mapReverse[stagedNode.Umlsc]
		pkgelt.Umlscs = remove(pkgelt.Umlscs, stagedNode.Umlsc)
		stagedNode.Umlsc.Unstage()
	}
	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// commit will clean up the stage associations
		stage.Commit()
	}
}
