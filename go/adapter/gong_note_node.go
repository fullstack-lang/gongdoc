package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type GongNoteNode struct {
	ElementNodeBase
	gongNote *gong_models.GongNote
}

// AddToDiagram implements diagrammer.ElementNode.
func (gongNoteNode *GongNoteNode) AddToDiagram() {
	panic("unimplemented")
}

var _ diagrammer.ModelElementNode = &GongNoteNode{}

func NewGongNoteNode(
	stage *gong_models.StageStruct,
	gongNote *gong_models.GongNote) (gongNoteNode *GongNoteNode) {
	gongNoteNode = &GongNoteNode{ElementNodeBase: ElementNodeBase{stage: stage}}

	gongNoteNode.stage = stage
	gongNoteNode.gongNote = gongNote
	return
}

// GenerateChildren implements diagrammer.Node.
func (gongNoteNode *GongNoteNode) GenerateChildren() (children []diagrammer.ModelNode) {

	for _, link := range gongNoteNode.gongNote.Links {
		linkNode := NewLinkNode(gongNoteNode.stage, link)
		children = append(children, linkNode)
	}

	return
}

// GetName implements diagrammer.Node.
func (gongNoteNode *GongNoteNode) GetName() string {
	return gongNoteNode.gongNote.GetName()
}

func (gongNoteNode *GongNoteNode) GetElement() any {
	return gongNoteNode.gongNote
}
