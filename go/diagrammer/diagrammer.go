package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

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

func (diagrammer *Diagrammer) PortfolioNode2NodeTree(portfolioNode PortfolioNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: portfolioNode.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = portfolioNode.IsExpanded()
	treeNode.HasCheckboxButton = portfolioNode.HasCheckboxButton()
	treeNode.IsCheckboxDisabled = !diagrammer.Portfolio.IsInSelectionMode()
	treeNode.Impl = &PortfolioNodeImpl{
		portfolio:     diagrammer.Portfolio,
		portfolioNode: portfolioNode}

	for _, children := range portfolioNode.GetChildren() {
		childrenTreeNode := diagrammer.PortfolioNode2NodeTree(children, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}
