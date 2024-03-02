package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioDiagramNodeButtonDuplicateImpl struct {
	portfolioDiagramNode            PortfolioDiagramNode
	containingPortfolioCategoryNode PortfolioCategoryNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewPortfolioDiagramNodeButtonDuplicateImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	containingPortfolioCategoryNode PortfolioCategoryNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (buttonImpl *PortfolioDiagramNodeButtonDuplicateImpl) {

	buttonImpl = new(PortfolioDiagramNodeButtonDuplicateImpl)

	buttonImpl.portfolioDiagramNode = portfolioDiagramNode
	buttonImpl.diagrammer = diagrammer
	buttonImpl.treeNode = treeNode
	buttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonDuplicateImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	buttonImpl.portfolioDiagramNode.DuplicateDiagram()

	if childrenPortfolioNode := buttonImpl.portfolioDiagramNode.DuplicateDiagram(); childrenPortfolioNode != nil {
		childrenTreeNode := buttonImpl.diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, buttonImpl.treeStage)
		buttonImpl.treeNode.Children = append(buttonImpl.treeNode.Children, childrenTreeNode)
		buttonImpl.diagrammer.map_portfolioNode_treeNode[childrenPortfolioNode] = childrenTreeNode
		buttonImpl.treeStage.Commit()
	}

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.treeStage.Commit()
}
