package diagrammer

import gongtree_models "github.com/fullstack-lang/gongtree/go/models"

type ModelNode interface {
	GetChildren() []ModelNode
	GetName() string
	IsNameEditable() bool

	IsExpanded() bool
	SetIsExpanded(isExpanded bool)

	HasCheckboxButton() bool

	// GetElement return the pointer to the model element
	GetElement() any
}

type ModelNodeImpl struct {
	diagrammer *Diagrammer
	modelNode  ModelNode
}

// OnAfterUpdate implements models.NodeImplInterface
func (modelNodeImpl *ModelNodeImpl) OnAfterUpdate(stage *gongtree_models.StageStruct, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		modelNodeImpl.modelNode.SetIsExpanded(!modelNodeImpl.modelNode.IsExpanded())

		// we need to update the stage node because when the tree stage is committed later
		// the expanded configuration of the node will be correct
		// if not, the front node will be in the configuration of the stage node
		// whatever action the end user performs
		stagedNode.IsExpanded = frontNode.IsExpanded
	}

}
