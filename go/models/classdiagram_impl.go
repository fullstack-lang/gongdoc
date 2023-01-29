package models

type ClassdiagramImpl struct {
	classdiagram *Classdiagram
	node         *Node
}

func (classdiagramImpl *ClassdiagramImpl) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

}
