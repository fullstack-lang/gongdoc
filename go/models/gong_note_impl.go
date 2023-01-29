package models

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

type GongNoteImpl struct {
	gongNote *gong_models.GongNote
	node     *Node
	nodeCb   *NodeCB
}

func (gongNoteImpl *GongNoteImpl) OnAfterUpdate(
	stage *StageStruct,
	stagedNode, frontNode *Node) {

	classdiagram := gongNoteImpl.nodeCb.GetSelectedClassdiagram()

	// adding a note shape
	if !stagedNode.IsChecked && frontNode.IsChecked {
		stage.Checkout()

		noteShape := (&NoteShape{Name: stagedNode.Name}).Stage()

		mapOfGongNotes := *gong_models.GetGongstructInstancesMap[gong_models.GongNote]()

		gongNote, ok := mapOfGongNotes[noteShape.Name]
		if !ok {
			log.Fatal("Unkown note ", noteShape.Name)
		}

		noteShape.Identifier = ShapenameToIdentifier(noteShape.Name)

		noteShape.Body = gongNote.Body
		noteShape.X = 30
		noteShape.Y = 30
		noteShape.Width = 240
		noteShape.Heigth = 63

		classdiagram.NoteShapes = append(classdiagram.NoteShapes, noteShape)
		updateNodesStates(stage, gongNoteImpl.nodeCb)
	}

	// suppression a note
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stage.Checkout()
		foundNote := false
		var noteShape *NoteShape
		var _noteShape *NoteShape
		for _, _noteShape = range classdiagram.NoteShapes {

			// strange behavior when the note is removed within the loop
			if IdentifierToShapename(_noteShape.Identifier) == stagedNode.Name && !foundNote {
				foundNote = true
				noteShape = _noteShape
			}
		}
		if !foundNote {
			log.Panicf("Note %s of field not present ", stagedNode.Name)
		}

		// remove the links of the note shape
		for _, noteLink := range noteShape.NoteShapeLinks {
			noteLink.Unstage()
			noteShape.NoteShapeLinks = remove(noteShape.NoteShapeLinks, noteLink)
		}
		classdiagram.NoteShapes = remove(classdiagram.NoteShapes, noteShape)
		noteShape.Unstage()
		updateNodesStates(stage, gongNoteImpl.nodeCb)
	}

}

func (GongNoteImpl *GongNoteImpl) OnAfterDelete(
	stage *StageStruct,
	stagedNode, frontNode *Node) {
}
