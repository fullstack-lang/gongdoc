package models

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
)

type GongLinkImpl struct {
	gongLink *gong_models.GongLink
	node     *Node
	nodeCb   *NodeCB
}

func (gongLinkImpl *GongLinkImpl) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	classdiagram := gongLinkImpl.nodeCb.GetSelectedClassdiagram()

	// find the parent node to find the gongstruct to find the classshape
	// the node is field, one needs to find the gongstruct that contains it
	// get the parent node
	parentNode := gongLinkImpl.nodeCb.map_Children_Parent[stagedNode]

	gongNoteImpl := parentNode.impl.(*GongNoteImpl)
	gongNote := gongNoteImpl.gongNote

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

		noteShapeLink := (&NoteShapeLink{Name: stagedNode.GetName()}).Stage()
		noteShapeLink.Identifier =
			ShapeAndFieldnameToFieldIdentifier(gongNote.Name, stagedNode.Name)

		noteshape.NoteShapeLinks = append(noteshape.NoteShapeLinks, noteShapeLink)

		updateNodesStates(stage, gongLinkImpl.nodeCb)
	}

	// removing a note link
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stage.Checkout()

		// get the relevant gong note link
		var noteLink *NoteShapeLink
		for _, _noteLink := range noteshape.NoteShapeLinks {
			if IdentifierToFieldName(_noteLink.Identifier) ==
				stagedNode.Name {
				noteLink = _noteLink
			}
		}

		noteLink.Unstage()
		noteshape.NoteShapeLinks = remove(noteshape.NoteShapeLinks, noteLink)

		updateNodesStates(stage, gongLinkImpl.nodeCb)
	}

}

func (gongLinkImpl *GongLinkImpl) OnAfterDelete(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
}
