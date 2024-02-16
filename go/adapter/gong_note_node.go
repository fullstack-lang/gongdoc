package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/bridge"
)

func NewGongNoteNode(
	stage *gong_models.StageStruct,
	gongNote *gong_models.GongNote) (gongNoteNode *GongNoteNode) {
	gongNoteNode = new(GongNoteNode)

	gongNoteNode.stage = stage
	gongNoteNode.gongNote = gongNote
	return
}

type GongNoteNode struct {
	stage    *gong_models.StageStruct
	gongNote *gong_models.GongNote
}

func (gongNoteNode *GongNoteNode) IsExpanded() bool {
	return false
}

func (gongNoteNode *GongNoteNode) HasCheckboxButton() bool {
	return true
}

// GetChildren implements bridge.Node.
func (gongNoteNode *GongNoteNode) GetChildren() (children []bridge.Node) {

	for _, link := range gongNoteNode.gongNote.Links {
		fieldNode := NewLinkNode(gongNoteNode.stage, link)
		children = append(children, fieldNode)
	}

	return
}

// GetName implements bridge.Node.
func (gongNoteNode *GongNoteNode) GetName() string {
	return gongNoteNode.gongNote.GetName()
}

// IsNameEditable implements bridge.Node.
func (gongNoteNode *GongNoteNode) IsNameEditable() bool {
	return false
}
