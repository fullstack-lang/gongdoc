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
func (diagramNodeImple *DiagramNodeImpl) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// manages the radio button stuff --> only one button at a time
		stagedNode.IsChecked = true
		diagramNodeImple.diagrammer.selectedDiagram = diagramNodeImple.diagramNode.GetDiagram()
		diagramNodeImple.diagrammer.portfolio.GenerateSVG(diagramNodeImple.diagramNode)
		diagramNodeImple.diagrammer.generatePortfolioNodesButtons()
	}
}
