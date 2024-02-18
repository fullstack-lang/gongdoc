// diagrammer is a package that manages
//
// A model that contains elements of different categories
// A portfolio that contains diagrams with shapes.
// Each shape is related to one element, one kind of shape for each kind of element
//
// Two trees:
//
// - a "model" tree displays all elements of the models
// - a "portfolio" tree displays all diagrams in the folder
//
// the "diagrams" tree allows for the selection of a diagram within the diagram package
// the "model" tree allows for the addition/suppression of shapes in the diagram
//
// the "model" tree is construed by getting informations from the model
package diagrammer

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

// Porfolio
//
// IsInSelectionMode()
type Portfolio interface {
	GetChildren() []PortfolioNode
	GetSelectedDiagram() Diagram
	IsInSelectionMode() bool // the end user can select a diagram to display
}

type Diagram interface {
}

type PortfolioNode interface {
	GetChildren() []PortfolioNode
	GetName() string
	IsNameEditable() bool
	IsExpanded() bool
	HasCheckboxButton() bool
}

type Diagrammer struct {
	Model     Model
	Portfolio Portfolio
	treeStage *gongtree_models.StageStruct
}

func NewDiagrammer(
	model Model,
	portfolio Portfolio,
	treeStage *gongtree_models.StageStruct,
) (diagrammer *Diagrammer) {
	diagrammer = new(Diagrammer)

	diagrammer.Model = model
	diagrammer.Portfolio = portfolio
	diagrammer.treeStage = treeStage
	return
}

// FillUpModelTree ranges over Model Root Nodes
// and recursively fill up the modelTree
func (diagrammer *Diagrammer) FillUpModelTree(modelTree *gongtree_models.Tree) {

	for _, node := range diagrammer.Model.GetChildren() {
		treeNode := diagrammer.ModelNode2NodeTree(node, diagrammer.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func (diagrammer *Diagrammer) ModelNode2NodeTree(node ModelNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: node.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = node.IsExpanded()
	treeNode.HasCheckboxButton = node.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = true

	for _, children := range node.GetChildren() {
		childrenTreeNode := diagrammer.ModelNode2NodeTree(children, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}

// FillUpPortfolioTree ranges over Portfolio Root Nodes
// and recursively fill up the PortfolioTree
func (diagrammer *Diagrammer) FillUpPortfolioTree(modelTree *gongtree_models.Tree) {

	for _, node := range diagrammer.Portfolio.GetChildren() {
		treeNode := diagrammer.PortfolioNode2NodeTree(node, diagrammer.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func (diagrammer *Diagrammer) PortfolioNode2NodeTree(node PortfolioNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: node.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = node.IsExpanded()
	treeNode.HasCheckboxButton = node.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = !diagrammer.Portfolio.IsInSelectionMode()

	for _, children := range node.GetChildren() {
		childrenTreeNode := diagrammer.PortfolioNode2NodeTree(children, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}
