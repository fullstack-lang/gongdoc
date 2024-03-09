package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioDiagramNodeButtonEditCancelImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewPortfolioDiagramNodeButtonEditCancelImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (buttonImpl *PortfolioDiagramNodeButtonEditCancelImpl) {

	buttonImpl = new(PortfolioDiagramNodeButtonEditCancelImpl)

	buttonImpl.portfolioDiagramNode = portfolioDiagramNode
	buttonImpl.diagrammer = diagrammer
	buttonImpl.treeNode = treeNode
	buttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonEditCancelImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	map_ModelNode_Shape := buttonImpl.portfolioDiagramNode.CancelEdit()

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
	buttonImpl.treeStage.Commit()
}
