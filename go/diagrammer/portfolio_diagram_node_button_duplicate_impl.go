package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioDiagramNodeButtonDuplicateImpl struct {
	portfolioDiagramNode        PortfolioDiagramNode
	parentPortfolioCategoryNode PortfolioNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewPortfolioDiagramNodeButtonDuplicateImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	parentPortfolioCategoryNode PortfolioNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (buttonImpl *PortfolioDiagramNodeButtonDuplicateImpl) {

	buttonImpl = new(PortfolioDiagramNodeButtonDuplicateImpl)

	buttonImpl.portfolioDiagramNode = portfolioDiagramNode
	buttonImpl.parentPortfolioCategoryNode = parentPortfolioCategoryNode
	buttonImpl.diagrammer = diagrammer
	buttonImpl.treeNode = treeNode
	buttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonDuplicateImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	parentTreeNode := buttonImpl.diagrammer.GetPortfiolioNodeTreeNodeEntry(buttonImpl.parentPortfolioCategoryNode)

	if childrenPortfolioNode := buttonImpl.portfolioDiagramNode.DuplicateDiagram(); childrenPortfolioNode != nil {
		childrenTreeNode := buttonImpl.diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, buttonImpl.treeStage)
		parentTreeNode.Children = append(parentTreeNode.Children, childrenTreeNode)
		buttonImpl.diagrammer.AddPortfiolioNodeTreeNodeEntry(childrenPortfolioNode, childrenTreeNode)
		buttonImpl.treeStage.Commit()
	}

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.treeStage.Commit()
}
