package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type Diagrammer struct {
	model     Model
	portfolio Portfolio
	treeStage *gongtree_models.StageStruct

	selectedPortfolioNode      PortfolioNode
	map_portfolioNode_treeNode map[PortfolioNode]*gongtree_models.Node
}

func NewDiagrammer(
	model Model,
	portfolio Portfolio,
	treeStage *gongtree_models.StageStruct,

) (diagrammer *Diagrammer) {
	diagrammer = new(Diagrammer)

	diagrammer.model = model
	diagrammer.portfolio = portfolio
	diagrammer.treeStage = treeStage

	diagrammer.map_portfolioNode_treeNode = make(map[PortfolioNode]*gongtree_models.Node)
	return
}

// FillUpModelTree ranges over Model Root Nodes
// and recursively fill up the modelTree
func (diagrammer *Diagrammer) FillUpModelTree(modelTree *gongtree_models.Tree) {

	for _, node := range diagrammer.model.GetChildren() {
		treeNode := diagrammer.modelNode2NodeTree(node, diagrammer.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func (diagrammer *Diagrammer) modelNode2NodeTree(node ModelNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: node.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = node.IsExpanded()
	treeNode.HasCheckboxButton = node.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = true

	for _, children := range node.GetChildren() {
		childrenTreeNode := diagrammer.modelNode2NodeTree(children, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}

// FillUpPortfolioTree ranges over Portfolio Root Nodes
// and recursively fill up the PortfolioTree
func (diagrammer *Diagrammer) FillUpPortfolioTree(modelTree *gongtree_models.Tree) {

	for _, portfolioNode := range diagrammer.portfolio.GetChildren() {
		treeNode := diagrammer.portfolioNode2NodeTree(portfolioNode, diagrammer.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
		diagrammer.map_portfolioNode_treeNode[portfolioNode] = treeNode
	}
}

func (diagrammer *Diagrammer) portfolioNode2NodeTree(portfolioNode PortfolioNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: portfolioNode.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = portfolioNode.IsExpanded()
	treeNode.HasCheckboxButton = portfolioNode.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = !diagrammer.portfolio.IsInSelectionMode()
	treeNode.Impl = &PortfolioNodeImpl{
		diagrammer:    diagrammer,
		portfolioNode: portfolioNode}

	for _, childrenPortfolioNode := range portfolioNode.GetChildren() {
		childrenTreeNode := diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
		diagrammer.map_portfolioNode_treeNode[childrenPortfolioNode] = childrenTreeNode
	}

	return
}

func (diagrammer *Diagrammer) computeCheckedStatusOfNodes() {
	for portfolioNode, treeNode := range diagrammer.map_portfolioNode_treeNode {
		if portfolioNode == diagrammer.selectedPortfolioNode {
			treeNode.IsChecked = true
		} else {
			treeNode.IsChecked = false
		}
	}
	diagrammer.treeStage.Commit()
}

func (diagrammer *Diagrammer) CommitTreeStage() {
	diagrammer.treeStage.Commit()
}

func (diagrammer *Diagrammer) GenerateSVG() {
	diagrammer.portfolio.GenerateSVG()
}
