package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioDiagramNodeButtonSaveImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewPortfolioDiagramNodeButtonSaveImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (buttonImpl *PortfolioDiagramNodeButtonSaveImpl) {

	buttonImpl = new(PortfolioDiagramNodeButtonSaveImpl)

	buttonImpl.portfolioDiagramNode = portfolioDiagramNode
	buttonImpl.diagrammer = diagrammer
	buttonImpl.treeNode = treeNode
	buttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonSaveImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	buttonImpl.portfolioDiagramNode.SetIsInEditMode(false)

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.treeStage.Commit()
}