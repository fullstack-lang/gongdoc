package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type CallbacksSingloton struct {
	ClassdiagramsRootNode *Node
	StateDiagramsRootNode *Node

	GongstructsRootNode *Node
	GongenumsRootNode   *Node
}

func (callbacksSingloton CallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// if node is selected
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

			// update the tree of identifiers in order to show
			// which identifiers are presents in the selected diagram
			updateNodesStates(stage, &callbacksSingloton)
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

				// check that the name is a correct identifer in go
				for _, b := range frontNode.Name {
					if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_' || '0' <= b && b <= '9' {
						// Avoid assigning a rune for the common case of an ascii character.
						continue
					} else {
						log.Println("The name of the diagram is not a correct identifier in go: " + frontNode.Name)
						stagedNode.Commit()
						return
					}
				}

				// rename the diagram file if it exists
				// remove the actual classdiagram file if it exsits
				fieldName := GetAssociationName[DiagramPackage]().Classdiagrams[0].Name
				mapReverse := GetSliceOfPointersReverseMap[DiagramPackage, Classdiagram](fieldName)
				pkgelt := mapReverse[stagedNode.Classdiagram]

				oldClassdiagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", stagedNode.Classdiagram.Name) + ".go"
				newClassdiagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", frontNode.Name) + ".go"

				if _, err := os.Stat(oldClassdiagramFilePath); err == nil {
					if err := os.Rename(oldClassdiagramFilePath, newClassdiagramFilePath); err != nil {
						log.Println("Error while renaming file " + oldClassdiagramFilePath + " : " + err.Error())
					} else {
						// change the name of the go variable that describes the diagram
						// in the package
						input, err := ioutil.ReadFile(newClassdiagramFilePath)
						if err != nil {
							log.Fatalln(err)
						}

						lines := strings.Split(string(input), "\n")

						for i, line := range lines {
							if strings.Contains(line, "var "+stagedNode.Classdiagram.Name) {
								lines[i] = strings.Replace(line, stagedNode.Classdiagram.Name, frontNode.Name, 1)
								continue
							}
						}
						output := strings.Join(lines, "\n")
						err = ioutil.WriteFile(newClassdiagramFilePath, []byte(output), 0644)
						if err != nil {
							log.Fatalln(err)
						}

						stagedNode.Classdiagram.Name = frontNode.Name
						stagedNode.Classdiagram.Commit()
					}
				}

			case STATE_DIAGRAM:
				stagedNode.Umlsc.Name = frontNode.Name
				stagedNode.Umlsc.Commit()
			}
			switch stagedNode.Type {
			case CLASS_DIAGRAM, STATE_DIAGRAM:
				stagedNode.Name = frontNode.Name
				stagedNode.IsInEditMode = false
				updateNodesStates(stage, &callbacksSingloton)
			}
		}

		if stagedNode.IsInEditMode != frontNode.IsInEditMode {
			stagedNode.IsInEditMode = frontNode.IsInEditMode
			updateNodesStates(stage, &callbacksSingloton)
		}

		if stagedNode.IsInDrawMode != frontNode.IsInDrawMode {
			stagedNode.IsInDrawMode = frontNode.IsInDrawMode
			switch stagedNode.Type {
			case CLASS_DIAGRAM:
				stagedNode.Classdiagram.IsInDrawMode = frontNode.IsInDrawMode
			case STATE_DIAGRAM:
				stagedNode.Umlsc.IsInDrawMode = frontNode.IsInDrawMode
			}
			updateNodesStates(stage, &callbacksSingloton)
		}

		// marshall diagram to switch to saved state
		if !stagedNode.IsSaved && frontNode.IsSaved {
			// fetch the only package
			var pkgelt *DiagramPackage
			for _pkgelt := range *GetGongstructInstancesSet[DiagramPackage]() {
				pkgelt = _pkgelt
			}
			switch stagedNode.Type {
			case CLASS_DIAGRAM:

				// checkout in order to get the latest version of the diagram updated
				// by the front
				Stage.Checkout()
				stagedNode.Classdiagram.Marshall(pkgelt, filepath.Join(pkgelt.Path, "../diagrams"))
				stage.Commit()
			case STATE_DIAGRAM:
				// stagedNode.Umlsc = stagedNode.Umlsc.Saved
			}
			stagedNode.IsSaved = true
			updateNodesStates(stage, &callbacksSingloton)
		}
	case GONG_STRUCT:

		// if node is unchecked
		if stagedNode.IsChecked && !frontNode.IsChecked {

			// get the latest version of the diagram
			stage.Checkout()

			// remove the classshape from the selected diagram
			for _, classdiagramNode := range callbacksSingloton.ClassdiagramsRootNode.Children {
				if classdiagramNode.IsChecked {
					// get the diagram
					classDiagram := classdiagramNode.Classdiagram

					// get the referenced gongstructs
					for _, classshape := range classDiagram.Classshapes {
						gongstruct := classshape.GongStruct
						if gongstruct.Name == stagedNode.Name {
							classDiagram.RemoveClassshape(gongstruct.Name)
						}
					}
				}

				updateNodesStates(stage, &callbacksSingloton)
			}
		}

		// if node is checked, add classshape
		if !stagedNode.IsChecked && frontNode.IsChecked {

			// get the latest version of the diagram
			stage.Checkout()

			for _, classdiagramNode := range callbacksSingloton.ClassdiagramsRootNode.Children {
				if classdiagramNode.IsChecked {
					// get the diagram
					classDiagram := classdiagramNode.Classdiagram
					classDiagram.AddClassshape(frontNode.Name)
				}

				updateNodesStates(stage, &callbacksSingloton)
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
		var pkgelt *DiagramPackage
		for _pkgelt := range *GetGongstructInstancesSet[DiagramPackage]() {
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
	}
	switch stagedNode.Type {
	case CLASS_DIAGRAM:
		// remove the classdiagram node from the pkg element node
		fieldName := GetAssociationName[DiagramPackage]().Classdiagrams[0].Name
		mapReverse := GetSliceOfPointersReverseMap[DiagramPackage, Classdiagram](fieldName)
		pkgelt := mapReverse[stagedNode.Classdiagram]
		pkgelt.Classdiagrams = remove(pkgelt.Classdiagrams, stagedNode.Classdiagram)
		stagedNode.Classdiagram.Unstage()

		// remove the actual classdiagram file if it exsits
		classdiagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", stagedNode.Classdiagram.Name) + ".go"
		if _, err := os.Stat(classdiagramFilePath); err == nil {
			if err := os.Remove(classdiagramFilePath); err != nil {
				log.Println("Error while deleting file " + classdiagramFilePath + " : " + err.Error())
			}
		}

	case STATE_DIAGRAM:
		// remove the umlsc node from the pkg element node
		fieldName := GetAssociationName[DiagramPackage]().Umlscs[0].Name
		mapReverse := GetSliceOfPointersReverseMap[DiagramPackage, Umlsc](fieldName)
		pkgelt := mapReverse[stagedNode.Umlsc]
		pkgelt.Umlscs = remove(pkgelt.Umlscs, stagedNode.Umlsc)
		stagedNode.Umlsc.Unstage()

		// remove the actual classdiagram file if it exsits
		statediagramFilePath := filepath.Join(pkgelt.Path, "../diagrams", stagedNode.Umlsc.Name) + ".go"
		if _, err := os.Stat(statediagramFilePath); err == nil {
			if err := os.Remove(statediagramFilePath); err != nil {
				log.Println("Error while deleting file " + statediagramFilePath + " : " + err.Error())
			}
		}
	}
	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// commit will clean up the stage associations
		stage.Commit()
	}
}