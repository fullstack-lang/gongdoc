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
	diagramPackage      *DiagramPackage
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

func (nodesCb *NodeCB) OnAfterCreate(
	stage *StageStruct,
	node *Node) {

	log.Println("Node " + node.Name + " is created")

	node.HasCheckboxButton = true

	// check unicity of name, otherwise, add an index
	var hasNameCollision bool
	initialName := node.Name
	index := 0
	// loop until the name of the new diagram is not in collision with an existing
	// diagram name
	for index == 0 || hasNameCollision {
		index++
		hasNameCollision = false
		for classdiagram := range *GetGongstructInstancesSet[Classdiagram]() {
			if classdiagram.Name == node.Name {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			node.Name = initialName + fmt.Sprintf("_%d", index)
		}
	}

	classdiagram := (&Classdiagram{Name: node.Name}).Stage()
	nodesCb.diagramPackage.Classdiagrams = append(nodesCb.diagramPackage.Classdiagrams, classdiagram)
	node.Classdiagram = classdiagram
	node.IsInEditMode = false
	node.IsInDrawMode = false
	node.HasEditButton = false

	nodesCb.ClassdiagramsRootNode.Children =
		append(nodesCb.ClassdiagramsRootNode.Children, node)

	updateNodesStates(stage, nodesCb)

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
		UnstageBranch(stage, stagedNode.Classdiagram)

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
		updateNodesStates(stage, nodesCb)
	}
}
