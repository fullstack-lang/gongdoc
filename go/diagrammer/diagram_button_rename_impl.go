package diagrammer

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type DiagramButtonRenameImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewDiagramButtonRenameImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (adbuttonImpl *DiagramButtonRenameImpl) {

	adbuttonImpl = new(DiagramButtonRenameImpl)

	adbuttonImpl.portfolioDiagramNode = portfolioDiagramNode
	adbuttonImpl.diagrammer = diagrammer
	adbuttonImpl.treeNode = treeNode
	adbuttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *DiagramButtonRenameImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	buttonImpl.portfolioDiagramNode.SetIsInRenameMode(true)
	buttonImpl.treeNode.IsInEditMode = true

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.treeStage.Commit()

}
