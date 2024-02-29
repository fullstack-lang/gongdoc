package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type DiagramButtonRenameCancelImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewDiagramButtonRenameCancelImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (adbuttonImpl *DiagramButtonRenameCancelImpl) {

	adbuttonImpl = new(DiagramButtonRenameCancelImpl)

	adbuttonImpl.portfolioDiagramNode = portfolioDiagramNode
	adbuttonImpl.diagrammer = diagrammer
	adbuttonImpl.treeNode = treeNode
	adbuttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *DiagramButtonRenameCancelImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	buttonImpl.portfolioDiagramNode.SetIsInRenameMode(false)
	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.treeStage.Commit()

}
