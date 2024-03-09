package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioDiagramNodeButtonEditImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewPortfolioDiagramNodeButtonEditImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (buttonImpl *PortfolioDiagramNodeButtonEditImpl) {

	buttonImpl = new(PortfolioDiagramNodeButtonEditImpl)

	buttonImpl.portfolioDiagramNode = portfolioDiagramNode
	buttonImpl.diagrammer = diagrammer
	buttonImpl.treeNode = treeNode
	buttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonEditImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	map_ModelNode_Shape := buttonImpl.portfolioDiagramNode.EditDiagram()

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
	buttonImpl.treeStage.Commit()
}
