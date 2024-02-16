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
	GetRootNodes() []Node
}

type Node interface {
	GetChildren() []Node
	GetName() string
	IsNameEditable() bool
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

	for _, rootNode := range bridge.Model.GetRootNodes() {
		treeNode := (&gongtree_models.Node{Name: rootNode.GetName()}).Stage(bridge.treeStage)
		FillUpTreeRecursively(rootNode, treeNode, bridge.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func FillUpTreeRecursively(node Node, treeNode *gongtree_models.Node, treeStage *gongtree_models.StageStruct) {
	for _, children := range node.GetChildren() {
		childrenTreeNode := (&gongtree_models.Node{Name: children.GetName()}).Stage(treeStage)
		FillUpTreeRecursively(children, childrenTreeNode, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}
}
