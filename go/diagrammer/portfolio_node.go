package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

type PortfolioNode interface {
	GetChildren() []PortfolioNode
	GetName() string
	IsNameEditable() bool
	IsExpanded() bool
	HasCheckboxButton() bool
	OnCheckboxButtonCheck()
	HasHadButton() bool
}

type PortfolioNodeImpl struct {
	diagrammer    *Diagrammer
	portfolioNode PortfolioNode
}

// OnAfterUpdate implements models.NodeImplInterface
func (portfolioNodeImpl *PortfolioNodeImpl) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {
	if frontNode.IsChecked && !stagedNode.IsChecked {

		// let the adapter do what it has to to
		portfolioNodeImpl.portfolioNode.OnCheckboxButtonCheck()

		// manages the radio button stuff --> only one button at a time
		stagedNode.IsChecked = true
		portfolioNodeImpl.diagrammer.selectedPortfolioNode = portfolioNodeImpl.portfolioNode
		portfolioNodeImpl.diagrammer.computeCheckedStatusOfNodes()

		// for _, otherDiagramNode := range nodeImplClasssiagram.diagramPackageNode.Children {
		// 	if otherDiagramNode == stagedNode {
		// 		continue
		// 	}

		// 	// uncheck the other node
		// 	if otherDiagramNode.IsChecked {
		// 		// log.Println("Node " + node.Name + " is checked and should be unchecked")
		// 		otherDiagramNode.IsChecked = false
		// 		otherDiagramNode.Commit(gongtreeStage)
		// 	}
		// }
	}
}
