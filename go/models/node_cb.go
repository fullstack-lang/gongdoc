package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type NodeCB struct {
	ClassdiagramsRootNode *Node

	idTree *Tree

	selectedClassdiagram *Classdiagram

	// map to navigate from a children node to its parent
	map_Children_Parent map[*Node]*Node

	// map to navigate from identifiers in the package
	// to nodes, and backforth
	// identifiers are unique in a package (that's the point of identifiers)
	map_Identifier_Node map[string]*Node
	map_Node_Identifier map[*Node]string

	diagramPackage *DiagramPackage
}

func (nodesCb *NodeCB) GetSelectedClassdiagram() (classdiagram *Classdiagram) {

	classdiagram = nodesCb.selectedClassdiagram

	return
}

// OnAfterUpdate is called each time the end user interacts
// with any node. The front commit the state of the front node
// to the back ([frontNode] in the function).
// Therefore, the [stageNode] is now different from the [frontNode].
//
// This functiion and its siblings analyse which field of the
// node has changed and performs all necessary actions
func (nodesCb *NodeCB) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:
		nodesCb.OnAfterUpdateDiagram(stage, stagedNode, frontNode)
	case GONG_STRUCT:
		nodesCb.OnAfterUpdateStruct(stage, stagedNode, frontNode)
	case GONG_STRUCT_FIELD:
		nodesCb.OnAfterUpdateStructField(stage, stagedNode, frontNode)
	case GONG_NOTE:
		nodesCb.OnAfterUpdateNote(stage, stagedNode, frontNode)
	case GONG_NOTE_LINK:
		nodesCb.OnAfterUpdateNoteLink(stage, stagedNode, frontNode)
	case GONG_ENUM:
		nodesCb.OnAfterUpdateEnum(stage, stagedNode, frontNode)
	case GONG_ENUM_VALUE:
		nodesCb.OnAfterUpdateEnumValue(stage, stagedNode, frontNode)
	}

	if stagedNode.IsExpanded != frontNode.IsExpanded {
		// setting the value of the staged node	to the new value
		stagedNode.IsExpanded = frontNode.IsExpanded
	}
}

func (nodesCb *NodeCB) OnAfterUpdateStruct(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		// remove the classshape from the selected diagram
		classDiagram := nodesCb.GetSelectedClassdiagram()
		classDiagram.RemoveClassshape(stagedNode.Name)

		updateNodesStates(stage, nodesCb)
	}

	// if node is checked, add classshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := nodesCb.GetSelectedClassdiagram()
		classDiagram.AddClassshape(nodesCb, frontNode.Name, REFERENCE_GONG_STRUCT)

		updateNodesStates(stage, nodesCb)
	}
}

func (nodesCb *NodeCB) OnAfterUpdateEnum(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		// remove the classshape from the selected diagram
		classDiagram := nodesCb.GetSelectedClassdiagram()

		// get the referenced gongstructs
		for _, classshape := range classDiagram.Classshapes {
			if IdentifierToShapename(classshape.Identifier) == stagedNode.Name {
				classDiagram.RemoveClassshape(IdentifierToShapename(classshape.Identifier))
			}

		}
		updateNodesStates(stage, nodesCb)
	}

	// if node is checked, add classshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := nodesCb.GetSelectedClassdiagram()
		classDiagram.AddClassshape(nodesCb, frontNode.Name, REFERENCE_GONG_ENUM)
		updateNodesStates(stage, nodesCb)
	}
}

func (nodesCb *NodeCB) OnAfterCreate(
	stage *StageStruct,
	newDiagramNode *Node) {

	log.Println("Node " + newDiagramNode.Name + " is created")

	switch newDiagramNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:
		newDiagramNode.HasCheckboxButton = true

		// check unicity of name, otherwise, add an index
		var hasNameCollision bool
		initialName := newDiagramNode.Name
		index := 0
		// loop until the name provide no collision
		for index == 0 || hasNameCollision {
			hasNameCollision = false
			for classdiagram := range *GetGongstructInstancesSet[Classdiagram]() {
				if classdiagram.Name == newDiagramNode.Name {
					hasNameCollision = true
				}
			}
			if hasNameCollision {
				newDiagramNode.Name = initialName + fmt.Sprintf("-%d", index)
				index++
			}
		}

		classdiagram := (&Classdiagram{Name: newDiagramNode.Name}).Stage()
		nodesCb.diagramPackage.Classdiagrams = append(nodesCb.diagramPackage.Classdiagrams, classdiagram)
		newDiagramNode.Classdiagram = classdiagram
		newDiagramNode.IsInEditMode = true
		newDiagramNode.IsInDrawMode = false
		newDiagramNode.HasEditButton = false

		nodesCb.ClassdiagramsRootNode.Children =
			append(nodesCb.ClassdiagramsRootNode.Children, newDiagramNode)

		updateNodesStates(stage, nodesCb)
		stage.Commit()
	}
}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (nodesCb *NodeCB) OnAfterDelete(
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
		nodesCb.diagramPackage.Classdiagrams = remove(nodesCb.diagramPackage.Classdiagrams, stagedNode.Classdiagram)
		stagedNode.Classdiagram.Unstage()

		// remove the actual classdiagram file if it exsits
		classdiagramFilePath := filepath.Join(nodesCb.diagramPackage.Path, "../diagrams", stagedNode.Classdiagram.Name) + ".go"
		if _, err := os.Stat(classdiagramFilePath); err == nil {
			if err := os.Remove(classdiagramFilePath); err != nil {
				log.Println("Error while deleting file " + classdiagramFilePath + " : " + err.Error())
			}
		}
	}
	switch stagedNode.Type {
	case CLASS_DIAGRAM, STATE_DIAGRAM:

		// commit will clean up the stage associations
		stage.Commit()
	}
}
