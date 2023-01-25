package models

func (nodesCb *NodeCB) OnAfterUpdateNoteLink(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	classdiagram := nodesCb.GetSelectedClassdiagram()

	// find the parent node to find the gongstruct to find the classshape
	// the node is field, one needs to find the gongstruct that contains it
	// get the parent node
	parentNode := nodesCb.map_Children_Parent[stagedNode]
	gongNote := parentNode.GongNote

	// find the classhape in the classdiagram
	foundNoteshape := false
	var noteshape *NoteShape
	_ = noteshape

	for _, _noteshape := range classdiagram.NoteShapes {
		// strange behavior when the classshape is remove within the loop
		if IdentifierToShapename(_noteshape.Identifier) ==
			gongNote.Name && !foundNoteshape {
			noteshape = _noteshape
		}
	}

	// adding a note link
	if !stagedNode.IsChecked && frontNode.IsChecked {
		stage.Checkout()

		noteLink := (&NoteLink{Name: stagedNode.GetName()}).Stage()
		noteLink.Identifier =
			ShapeAndFieldnameToFieldIdentifier(gongNote.Name, stagedNode.Name)

		noteshape.NoteLinks = append(noteshape.NoteLinks, noteLink)

		Stage.Commit()
	}

	// removing a note link
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stage.Checkout()

		// get the relevant gong note link
		var noteLink *NoteLink
		for _, _noteLink := range noteshape.NoteLinks {
			if IdentifierToFieldName(_noteLink.Identifier) ==
				stagedNode.Name {
				noteLink = _noteLink
			}
		}

		noteshape.NoteLinks = remove(noteshape.NoteLinks, noteLink)

		Stage.Commit()
	}
}
