package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

type DiagramNode interface {
	PortfolioNode

	GetDiagram() Diagram
}

type DiagramNodeImpl struct {
	diagrammer  *Diagrammer
	diagramNode DiagramNode
}

// OnAfterUpdate implements models.NodeImplInterface
func (diagramNodeImpl *DiagramNodeImpl) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {
	if frontNode.IsChecked && !stagedNode.IsChecked {
		stagedNode.IsChecked = true
		diagramNodeImpl.diagrammer.selectedDiagram = diagramNodeImpl.diagramNode.GetDiagram()
		map_ModelNode_Shape := diagramNodeImpl.diagrammer.portfolio.GenerateDiagram(diagramNodeImpl.diagramNode)
		diagramNodeImpl.diagrammer.generatePortfolioNodesButtons()
		diagramNodeImpl.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
		diagramNodeImpl.diagrammer.treeStage.Commit()
	}
}
