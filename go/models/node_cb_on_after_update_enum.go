package models

func (nodesCb *NodeCB) OnAfterUpdateEnum(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		// remove the classshape from the selected diagram
		classDiagram := nodesCb.GetSelectedClassdiagram()

		// get the referenced gongstructs
		for _, classshape := range classDiagram.Classshapes {
			if IdentifierToShapename(classshape.Identifier) == stagedNode.Name {
				classDiagram.RemoveClassshape(IdentifierToShapename(classshape.Identifier))
			}

		}
		updateNodesStates(stage, nodesCb)
	}

	// if node is checked, add classshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classDiagram := nodesCb.GetSelectedClassdiagram()
		classDiagram.AddClassshape(nodesCb, frontNode.Name, REFERENCE_GONG_ENUM)
		updateNodesStates(stage, nodesCb)
	}
}
