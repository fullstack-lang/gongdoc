package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

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

	// renaming
	if stagedNode.Name != frontNode.Name {

		if err := portfolioDiagramNodeImpl.portfolioDiagramNode.RenameDiagram(frontNode.Name); err != nil {

			// force the stage node name to the front
			portfolioDiagramNodeImpl.portfolioDiagramNode.SetIsInRenameMode(false)
			stagedNode.IsInEditMode = false
			portfolioDiagramNodeImpl.diagrammer.generatePortfolioNodesButtons()
			portfolioDiagramNodeImpl.diagrammer.treeStage.Commit()
			return
		}

		stagedNode.Name = frontNode.Name
		stagedNode.IsInEditMode = false
		portfolioDiagramNodeImpl.portfolioDiagramNode.SetIsInRenameMode(false)
		portfolioDiagramNodeImpl.diagrammer.generatePortfolioNodesButtons()
		portfolioDiagramNodeImpl.diagrammer.treeStage.Commit()
	}
}