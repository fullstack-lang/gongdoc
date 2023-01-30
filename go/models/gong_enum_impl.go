package models

import gong_models "github.com/fullstack-lang/gong/go/models"

type GongEnumImpl struct {
	gongEnum *gong_models.GongEnum
	NodeImpl
}

func (gongEnumImpl *GongEnumImpl) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		// remove the classshape from the selected diagram
		classDiagram := gongEnumImpl.nodeCb.GetSelectedClassdiagram()

		// get the referenced gongstructs
		for _, classshape := range classDiagram.GongStructShapes {
			if IdentifierToShapename(classshape.Identifier) == stagedNode.Name {
				classDiagram.RemoveClassshape(IdentifierToShapename(classshape.Identifier))
			}

		}
		updateNodesStates(stage, gongEnumImpl.nodeCb)
	}

	// if node is checked, add classshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := gongEnumImpl.nodeCb.GetSelectedClassdiagram()
		classDiagram.AddClassshape(gongEnumImpl.nodeCb, frontNode.Name, REFERENCE_GONG_ENUM)
		updateNodesStates(stage, gongEnumImpl.nodeCb)
	}

}

func (gongEnumImpl *GongEnumImpl) OnAfterDelete(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
}
