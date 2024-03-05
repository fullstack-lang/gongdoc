package diagrammer

import (
	"slices"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type PortfolioDiagramNodeButtonRemoveImpl struct {
	portfolioDiagramNode PortfolioDiagramNode

	diagrammer *Diagrammer

	treeNode *gongtree_models.Node

	treeStage *gongtree_models.StageStruct
}

func NewPortfolioDiagramNodeButtonRemoveImpl(
	portfolioDiagramNode PortfolioDiagramNode,
	diagrammer *Diagrammer,
	treeNode *gongtree_models.Node,
	treeStage *gongtree_models.StageStruct,
) (buttonImpl *PortfolioDiagramNodeButtonRemoveImpl) {

	buttonImpl = new(PortfolioDiagramNodeButtonRemoveImpl)

	buttonImpl.portfolioDiagramNode = portfolioDiagramNode
	buttonImpl.diagrammer = diagrammer
	buttonImpl.treeNode = treeNode
	buttonImpl.treeStage = treeStage

	return
}

func (buttonImpl *PortfolioDiagramNodeButtonRemoveImpl) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	parentTreeNode := buttonImpl.diagrammer.GetPortfiolioNodeFromTreeNode(buttonImpl.portfolioDiagramNode.GetParent())

	buttonImpl.portfolioDiagramNode.DeleteDiagram()

	// remove the node from the parent
	buttonImpl.portfolioDiagramNode.GetParent().RemoveChildren(buttonImpl.portfolioDiagramNode)

	idx := slices.Index(parentTreeNode.Children, buttonImpl.treeNode)
	parentTreeNode.Children = slices.Delete(parentTreeNode.Children, idx, idx+1)

	buttonImpl.diagrammer.RemovePortfiolioNodeTreeNodeEntry(buttonImpl.portfolioDiagramNode)
	buttonImpl.treeStage.Commit()

	buttonImpl.diagrammer.generatePortfolioNodesButtons()
	buttonImpl.treeStage.Commit()
}
