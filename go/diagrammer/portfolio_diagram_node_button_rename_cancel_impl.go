package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioDiagramNodeButtonRenameCancelImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewPortfolioDiagramNodeButtonRenameCancelImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (adbuttonImpl *PortfolioDiagramNodeButtonRenameCancelImpl) {

	adbuttonImpl = new(PortfolioDiagramNodeButtonRenameCancelImpl)

	adbuttonImpl.portfolioDiagramNode = portfolioDiagramNode
	adbuttonImpl.diagrammer = diagrammer
	adbuttonImpl.treeNode = treeNode
	adbuttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonRenameCancelImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	buttonImpl.portfolioDiagramNode.SetIsInRenameMode(false)
	buttonImpl.treeNode.IsInEditMode = false

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.treeStage.Commit()

}
