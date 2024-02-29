package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type AddDiagramButtonImpl struct {
	portfolioCategoryNode PortfolioCategoryNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewAddButtonImpl(
	portfolioCategoryNode PortfolioCategoryNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (adbuttonImpl *AddDiagramButtonImpl) {

	adbuttonImpl = new(AddDiagramButtonImpl)

	adbuttonImpl.portfolioCategoryNode = portfolioCategoryNode
	adbuttonImpl.diagrammer = diagrammer
	adbuttonImpl.treeNode = treeNode
	adbuttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *AddDiagramButtonImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	if childrenPortfolioNode := buttonImpl.portfolioCategoryNode.AddDiagram(); childrenPortfolioNode != nil {
		childrenTreeNode := buttonImpl.diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, buttonImpl.treeStage)
		buttonImpl.treeNode.Children = append(buttonImpl.treeNode.Children, childrenTreeNode)
		buttonImpl.diagrammer.map_portfolioNode_treeNode[childrenPortfolioNode] = childrenTreeNode
		buttonImpl.treeStage.Commit()
	}
}
