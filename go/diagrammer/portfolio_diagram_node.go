package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

type PortfolioDiagramNode interface {
	PortfolioNode

	GetDiagram() Diagram

	// DisplayDiagram request the portfolio to display the diagram
	DisplayDiagram() map[ModelNode]Shape

	HasDiagramRenameButton() bool
	RenameDiagram(newName string)
	IsInRenameMode() bool
	SetIsInRenameMode(isInRenameMode bool)
}

type PortfolioDiagramNodeImpl struct {
	diagrammer           *Diagrammer
	portfolioDiagramNode PortfolioDiagramNode
}

// OnAfterUpdate implements models.NodeImplInterface
func (portfolioDiagramNodeImpl *PortfolioDiagramNodeImpl) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {
	if frontNode.IsChecked && !stagedNode.IsChecked {
		stagedNode.IsChecked = true
		portfolioDiagramNodeImpl.diagrammer.selectedDiagram = portfolioDiagramNodeImpl.portfolioDiagramNode.GetDiagram()
		map_ModelNode_Shape := portfolioDiagramNodeImpl.portfolioDiagramNode.DisplayDiagram()

		portfolioDiagramNodeImpl.diagrammer.generatePortfolioNodesButtons()
		portfolioDiagramNodeImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
		portfolioDiagramNodeImpl.diagrammer.treeStage.Commit()
	}
}
