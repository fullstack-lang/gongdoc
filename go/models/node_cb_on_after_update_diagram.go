package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (nodesCb *NodeCB) OnAfterUpdateDiagram(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	// node has been checked by the end user
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// setting the value of the staged node	to the new value
		// and commit to the database,
		// The front will detect that the backend has been commited
		// It will refresh and fetch the node with checked value
		stagedNode.IsChecked = true
		nodesCb.diagramPackage.SelectedClassdiagram = stagedNode.Classdiagram

		// uncheck all other diagram nodes
		diagramNodes := append(
			nodesCb.diagramPackageNode.Children)

		for _, otherDiagramNode := range diagramNodes {
			if otherDiagramNode == stagedNode {
				continue
			}

			// uncheck the other node
			if otherDiagramNode.IsChecked {
				// log.Println("Node " + node.Name + " is checked and should be unchecked")
				otherDiagramNode.IsChecked = false
			}
		}

		// update the nodes in the tree of identifiers in order to update
		// which identifiers are present/absent in the selected diagram
		updateNodesStates(stage, nodesCb)
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
			oldClassdiagramFilePath := filepath.Join(nodesCb.diagramPackage.Path, "../diagrams", stagedNode.Classdiagram.Name) + ".go"
			newClassdiagramFilePath := filepath.Join(nodesCb.diagramPackage.Path, "../diagrams", frontNode.Name) + ".go"

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
		}
		switch stagedNode.Type {
		case CLASS_DIAGRAM, STATE_DIAGRAM:
			stagedNode.Name = frontNode.Name
			stagedNode.IsInEditMode = false
			updateNodesStates(stage, nodesCb)
		}
	}

	// the end user switch the edit mode
	if stagedNode.IsInEditMode != frontNode.IsInEditMode {
		stagedNode.IsInEditMode = frontNode.IsInEditMode
		updateNodesStates(stage, nodesCb)
	}

	if stagedNode.IsInDrawMode != frontNode.IsInDrawMode {
		stagedNode.IsInDrawMode = frontNode.IsInDrawMode
		switch stagedNode.Type {
		case CLASS_DIAGRAM:
			stagedNode.Classdiagram.IsInDrawMode = frontNode.IsInDrawMode
		}
		updateNodesStates(stage, nodesCb)
	}

	// marshall diagram to switch to saved state
	if !stagedNode.IsSaved && frontNode.IsSaved {
		switch stagedNode.Type {
		case CLASS_DIAGRAM:

			// checkout in order to get the latest version of the diagram before
			// modifying it updated by the front
			Stage.Checkout()

			Stage.Unstage()
			StageBranch(&Stage, stagedNode.Classdiagram)

			filepath := filepath.Join(
				filepath.Join(nodesCb.diagramPackage.AbsolutePathToDiagramPackage,
					"../diagrams"),
				stagedNode.Classdiagram.Name) + ".go"
			file, err := os.Create(filepath)
			if err != nil {
				log.Fatal("Cannot open diagram file" + err.Error())
			}
			defer file.Close()
			Stage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

			// restore the original stage
			Stage.Unstage()
			Stage.Checkout()
			stagedNode.IsSaved = false
			stage.Commit()
		case STATE_DIAGRAM:
			// stagedNode.Umlsc = stagedNode.Umlsc.Saved
		}

		updateNodesStates(stage, nodesCb)
	}
}
