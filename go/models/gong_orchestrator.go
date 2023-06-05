package models

// insertion point
// NodeOrchestrator
type NodeOrchestrator struct {
}

func (orchestrator *NodeOrchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedNode, backRepoNode *Node) {

	stagedNode.OnAfterUpdate(gongsvgStage, stagedNode, backRepoNode)
}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point
	case Node:
		stage.OnAfterNodeUpdateCallback = new(NodeOrchestrator)

	}

}
