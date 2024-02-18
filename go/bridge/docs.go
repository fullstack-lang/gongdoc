// bridge is a package that manages
//
// a model contains element
// a folder that contains diagrams with shapes
// a "diagrams" tree displays all diagrams in the folder
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
	GetChildren() []ModelNode
}

type ModelNode interface {
	GetChildren() []ModelNode
	GetName() string
	IsNameEditable() bool
	IsExpanded() bool
	HasCheckboxButton() bool
}

type Portfolio interface {
	GetChildren() []PortfolioNode
}

type PortfolioNode interface {
	GetChildren() []PortfolioNode
	GetName() string
	IsNameEditable() bool
	IsExpanded() bool
	HasCheckboxButton() bool
}

type Bridge struct {
	Model     Model
	Portfolio Portfolio
	treeStage *gongtree_models.StageStruct
}

func NewBridge(
	model Model,
	portfolio Portfolio,
	treeStage *gongtree_models.StageStruct,
) (bridge *Bridge) {
	bridge = new(Bridge)

	bridge.Model = model
	bridge.Portfolio = portfolio
	bridge.treeStage = treeStage
	return
}

// FillUpModelTree ranges over Model Root Nodes
// and recursively fill up the modelTree
func (bridge *Bridge) FillUpModelTree(modelTree *gongtree_models.Tree) {

	for _, node := range bridge.Model.GetChildren() {
		treeNode := ModelNode2NodeTree(node, bridge.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

// FillUpPortfolioTree ranges over Portfolio Root Nodes
// and recursively fill up the PortfolioTree
func (bridge *Bridge) FillUpPortfolioTree(modelTree *gongtree_models.Tree) {

	for _, node := range bridge.Portfolio.GetChildren() {
		treeNode := PortfolioNode2NodeTree(node, bridge.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func ModelNode2NodeTree(node ModelNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: node.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = node.IsExpanded()
	treeNode.HasCheckboxButton = node.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = true

	for _, children := range node.GetChildren() {
		childrenTreeNode := ModelNode2NodeTree(children, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}

func PortfolioNode2NodeTree(node PortfolioNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: node.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = node.IsExpanded()
	treeNode.HasCheckboxButton = node.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = true

	for _, children := range node.GetChildren() {
		childrenTreeNode := PortfolioNode2NodeTree(children, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}
