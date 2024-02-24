package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

// PortfolioNode has to be implemented by the Adapter Nodes
type PortfolioNode interface {

	// GetName() returns the Name that is displayed on the node
	GetName() string

	// GetChildren returns the nodes below the node
	GetChildren() []PortfolioNode

	// IsNameEditable is the adapter provides function for updating the name
	IsNameEditable() bool

	// IsExpanded is true if the the node is visualy expanded
	IsExpanded() bool

	// HasCheckboxButton is true if the node has a checkbox button
	HasCheckboxButton() bool

	// OnCheckboxButtonCheck is call if the check button is checked
	OnCheckboxButtonCheck()

	// HasAddButton is true if a "Add" button has to be displayed
	HasAddButton() bool
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
		portfolioNodeImpl.diagrammer.selectedDiagram = portfolioNodeImpl.portfolioNode
		portfolioNodeImpl.diagrammer.generatePortfolioNodesButtons()
	}
}
