// bridge is a package that manages
//
// a model contains element
// a diagram package that contains diagrams with shapes
// a "diagrams" tree displays all diagrams in the diagram package
// a "model" tree displays all elements of the models
//
// the "diagrams" tree allows for the selection of a diagram within the diagram package
// the "model" tree allows for the addition/suppression of shapes in the diagram
//
// the "model" tree is construed by getting informations from the model
package bridge

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type Model interface {
	GetChildren() []Node
}

type Node interface {
	GetChildren() []Node
	GetName() string
	IsNameEditable() bool
	IsExpanded() bool
	HasCheckboxButton() bool
}

type Bridge struct {
	Model     Model
	treeStage *gongtree_models.StageStruct
}

func NewBridge(
	model Model,
	treeStage *gongtree_models.StageStruct,
) (bridge *Bridge) {
	bridge = new(Bridge)

	bridge.Model = model
	bridge.treeStage = treeStage
	return
}

// FillUpTree ranges over Model Root Nodes
// and recursively fill up the modelTree
func (bridge *Bridge) FillUpTree(modelTree *gongtree_models.Tree) {

	for _, node := range bridge.Model.GetChildren() {
		treeNode := Node2NodeTree(node, bridge.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func Node2NodeTree(node Node, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: node.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = node.IsExpanded()
	treeNode.HasCheckboxButton = node.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = true

	for _, children := range node.GetChildren() {
		childrenTreeNode := Node2NodeTree(children, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}
