package models

type NodeImplInterface interface {

	// OnAfterUpdate function is called each time a node is modified
	OnAfterUpdate(stage *StageStruct, stagedNode, frontNode *Node)

	// OnAfterUpdate function is called each time a node is deleted
	OnAfterDelete(stage *StageStruct, stagedNode, frontNode *Node)

	// here the node will get from the implementation the instruction
	// of wether is has to be checked or not
	HasToBeChecked() bool
	SetHasToBeCheckedValue(value bool)

	HasToBeDisabled() bool
	SetHasToBeDisabledValue(value bool)
}
