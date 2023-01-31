package models

import gong_models "github.com/fullstack-lang/gong/go/models"

type GongStructImpl struct {
	gongStruct *gong_models.GongStruct
	NodeImpl
}

func (gongStructImpl *GongStructImpl) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		// remove the classshape from the selected diagram
		classDiagram := gongStructImpl.nodeCb.GetSelectedClassdiagram()
		classDiagram.RemoveClassshape(stagedNode.Name)

		updateNodesStates(stage, gongStructImpl.nodeCb)
	}

	// if node is checked, add classshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := gongStructImpl.nodeCb.GetSelectedClassdiagram()
		classDiagram.AddClassshape(frontNode.Name, REFERENCE_GONG_STRUCT)

		updateNodesStates(stage, gongStructImpl.nodeCb)
	}
}

func (gongStructImpl *GongStructImpl) OnAfterDelete(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
}
