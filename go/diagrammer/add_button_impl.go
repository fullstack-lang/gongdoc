package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type AddButtonImpl struct {
	portfolioNode PortfolioNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewAddButtonImpl(
	portfolioNode PortfolioNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (AddbuttonImpl *AddButtonImpl) {

	AddbuttonImpl = new(AddButtonImpl)

	AddbuttonImpl.diagrammer = diagrammer
	AddbuttonImpl.treeNode = treeNode
	AddbuttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *AddButtonImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	if childrenPortfolioNode := buttonImpl.diagrammer.portfolio.AddDiagram(buttonImpl.portfolioNode); childrenPortfolioNode != nil {
		childrenTreeNode := buttonImpl.diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, buttonImpl.treeStage)
		buttonImpl.treeNode.Children = append(buttonImpl.treeNode.Children, childrenTreeNode)
		buttonImpl.diagrammer.map_portfolioNode_treeNode[childrenPortfolioNode] = childrenTreeNode
		buttonImpl.treeStage.Commit()
	}
}
