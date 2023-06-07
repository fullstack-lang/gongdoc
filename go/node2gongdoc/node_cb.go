package node2gongdoc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// NodeCB is the singloton callback implementation of CUD operations on node
type NodeCB struct {

	// diagramPackageNode is used for iterating on classdiagram nodes
	// and for getting the selected diagram
	diagramPackageNode *gongdoc_models.Node
	diagramPackage     *gongdoc_models.DiagramPackage

	// treeOfGongObjects is the root of all nodes related to gong objects
	// it is necessary for performing operation on all elements of the tree
	treeOfGongObjects *gongdoc_models.Tree

	// map_Children_Parent is a map to navigate from a children node to its parent node
	// it is set up once at the init phase
	map_Children_Parent map[*gongdoc_models.Node]*gongdoc_models.Node
}

// GetSelectedClassdiagram
func (nodesCb *NodeCB) GetSelectedClassdiagram() (classdiagram *gongdoc_models.Classdiagram) {

	classdiagram = nodesCb.diagramPackage.SelectedClassdiagram
	return
}

// OnAfterUpdate is called each time the end user interacts
// with any node. The front commit the state of the front node
// to the back ([frontNode] in the function).
// Therefore, the [stageNode] is now different from the [frontNode].
//
// This functiion and its siblings analyse which field of the
// node has changed and performs all necessary actions
func (nodeCb *NodeCB) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	if stagedNode.IsExpanded != frontNode.IsExpanded {
		// setting the value of the staged node	to the new value
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	if stagedNode.Impl != nil {
		stagedNode.Impl.OnAfterUpdate(stage, stagedNode, frontNode)
		nodeCb.computeNodesConfiguration(stage)
	}

}

// OnAfterCreate is another callback
func (nodeCb *NodeCB) OnAfterCreate(
	gongdocStage *gongdoc_models.StageStruct,
	node *gongdoc_models.Node) {

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
		for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
			if classdiagram.Name == node.Name {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			node.Name = initialName + fmt.Sprintf("_%d", index)
		}
	}

	classdiagram := (&gongdoc_models.Classdiagram{Name: node.Name}).Stage(nodeCb.diagramPackage.Stage_)
	nodeCb.diagramPackage.Classdiagrams = append(nodeCb.diagramPackage.Classdiagrams, classdiagram)
	node.IsInEditMode = false
	node.IsInDrawMode = false
	node.HasEditButton = false
	node.HasDuplicateButton = false

	// append buttons
	drawButton := (&gongdoc_models.Button{Name: string(BUTTON_draw)}).Stage(gongdocStage)
	drawButton.Icon = "draw"
	drawButton.Displayed = false

	nodeCb.diagramPackageNode.Children = append(nodeCb.diagramPackageNode.Children, node)

	// set up the back pointer from the shape to the node
	classdiagramImpl := new(ClassdiagramImpl)
	classdiagramImpl.node = node
	classdiagramImpl.classdiagram = classdiagram
	classdiagramImpl.nodeCb = nodeCb
	node.Impl = classdiagramImpl

	filepath := filepath.Join(
		filepath.Join(classdiagramImpl.nodeCb.diagramPackage.AbsolutePathToDiagramPackage,
			"../diagrams"),
		classdiagramImpl.classdiagram.Name) + ".go"
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("Cannot open diagram file" + err.Error())
	}
	defer file.Close()

	// save the diagram
	// checkout in order to get the latest version of the diagram before
	// modifying it updated by the front
	nodeCb.computeNodesConfiguration(gongdocStage)
	gongdocStage.Commit()

	// now save the diagram
	gongdocStage.Checkout()
	gongdocStage.Unstage()
	gongdoc_models.StageBranch(gongdocStage, classdiagramImpl.classdiagram)

	gongdoc_models.SetupMapDocLinkRenaming(nodeCb.diagramPackage.ModelPkg.Stage_, gongdocStage)

	gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

	// restore the original stage
	gongdocStage.Unstage()
	gongdocStage.Checkout()

}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (nodeCb *NodeCB) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	switch impl := stagedNode.Impl.(type) {
	case *ClassdiagramImpl:
		impl.OnAfterDelete(stage, stagedNode, frontNode)
	}

	nodeCb.computeNodesConfiguration(stage)
}

// computeNodesConfiguration computes both trees
func (nodeCb *NodeCB) computeNodesConfiguration(gongdocStage *gongdoc_models.StageStruct) {

	computeNodeConfs(gongdocStage,
		nodeCb.diagramPackageNode,
		nodeCb.diagramPackage,
		nodeCb.treeOfGongObjects)

}

func (nodesCb *NodeCB) computeDiagramNodesConfigurations(stage *gongdoc_models.StageStruct) {

	computeClassdiagramNodesConfigurations(
		nodesCb.diagramPackageNode,
		nodesCb.diagramPackage,
		stage)
}
