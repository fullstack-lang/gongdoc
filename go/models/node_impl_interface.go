package models

type NodeImplInterface interface {
	OnAfterUpdate(stage *StageStruct, stagedNode, frontNode *Node)
	OnAfterDelete(stage *StageStruct, stagedNode, frontNode *Node)
}
