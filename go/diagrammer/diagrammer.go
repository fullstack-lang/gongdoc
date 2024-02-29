package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"
)

type Diagrammer struct {
	model     Model
	portfolio Portfolio
	treeStage *gongtree_models.StageStruct

	// selectedDiagram is the currently selected diagram
	selectedDiagram Diagram

	map_portfolioNode_treeNode map[PortfolioNode]*gongtree_models.Node
	map_elementNode_treeNode   map[ElementNode]*gongtree_models.Node
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
	diagrammer.map_elementNode_treeNode = make(map[ElementNode]*gongtree_models.Node)

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

func (diagrammer *Diagrammer) modelNode2NodeTree(modelNode ModelNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: modelNode.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = modelNode.IsExpanded()

	if elementNode, ok := modelNode.(ElementNode); ok {
		treeNode.HasCheckboxButton = true
		treeNode.IsCheckboxDisabled = true
		diagrammer.map_elementNode_treeNode[elementNode] = treeNode
	}

	treeNode.Impl = &ModelNodeImpl{
		diagrammer: diagrammer,
		modelNode:  modelNode}

	for _, childrenModelNode := range modelNode.GetChildren() {
		childrenTreeNode := diagrammer.modelNode2NodeTree(childrenModelNode, treeStage)
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
	}
	diagrammer.generatePortfolioNodesButtons()
}

func (diagrammer *Diagrammer) portfolioNode2NodeTree(portfolioNode PortfolioNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: portfolioNode.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = portfolioNode.IsExpanded()

	if diagramNode, ok := portfolioNode.(PortfolioDiagramNode); ok {
		_ = diagramNode
		treeNode.HasCheckboxButton = true
		treeNode.Impl = &PortfolioDiagramNodeImpl{
			diagrammer:           diagrammer,
			portfolioDiagramNode: diagramNode}
	}

	treeNode.IsCheckboxDisabled = !diagrammer.portfolio.IsInSelectionMode()
	diagrammer.map_portfolioNode_treeNode[portfolioNode] = treeNode

	for _, childrenPortfolioNode := range portfolioNode.GetChildren() {
		childrenTreeNode := diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)

	}

	return
}

func (diagrammer *Diagrammer) generatePortfolioNodesButtons() {
	for portfolioNode, treeNode := range diagrammer.map_portfolioNode_treeNode {

		// remove all buttons
		for _, _button := range treeNode.Buttons {
			_button.Unstage(diagrammer.treeStage)
		}
		treeNode.Buttons = make([]*gongtree_models.Button, 0)

		if portfolioCategoryNode, ok := portfolioNode.(PortfolioCategoryNode); ok {

			if portfolioCategoryNode.HasAddDiagramButton() {
				addDocButton := (&gongtree_models.Button{
					Name: portfolioCategoryNode.GetName() + " " + string(maticons.BUTTON_add),
					Icon: string(maticons.BUTTON_add)}).Stage(diagrammer.treeStage)
				treeNode.Buttons = append(treeNode.Buttons, addDocButton)
				addDocButton.Impl = NewAddButtonImpl(
					portfolioCategoryNode,
					diagrammer,
					treeNode,
					diagrammer.treeStage,
				)
			}
		}

		if diagramNode, ok := portfolioNode.(PortfolioDiagramNode); ok {
			if diagramNode.GetDiagram() == diagrammer.selectedDiagram {
				treeNode.IsChecked = true
			} else {
				treeNode.IsChecked = false
			}
		}

	}
	diagrammer.treeStage.Commit()
}

func (diagrammer *Diagrammer) CommitTreeStage() {
	diagrammer.treeStage.Commit()
}

func (diagrammer *Diagrammer) GetMap_elementNode_treeNode() map[ElementNode]*gongtree_models.Node {
	return diagrammer.map_elementNode_treeNode
}

// computeModelNodeStatus parses all nodes in the Model Tree
// and unchecks all nodes unless nodes matches an element in the diagram
func (diagrammer *Diagrammer) computeModelNodeStatus(map_ModelNode_Shape map[ModelNode]Shape) {
	for modelNode, treeNode := range diagrammer.map_elementNode_treeNode {
		treeNode.IsChecked = false
		if _, ok := map_ModelNode_Shape[modelNode]; ok {
			treeNode.IsChecked = true
		}
	}
}
