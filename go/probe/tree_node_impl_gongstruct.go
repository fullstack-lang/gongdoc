// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Classdiagram" {
		fillUpTable[models.Classdiagram](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DiagramPackage" {
		fillUpTable[models.DiagramPackage](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Field" {
		fillUpTable[models.Field](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumShape" {
		fillUpTable[models.GongEnumShape](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValueEntry" {
		fillUpTable[models.GongEnumValueEntry](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStructShape" {
		fillUpTable[models.GongStructShape](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShape" {
		fillUpTable[models.NoteShape](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "NoteShapeLink" {
		fillUpTable[models.NoteShapeLink](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Position" {
		fillUpTable[models.Position](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "UmlState" {
		fillUpTable[models.UmlState](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Umlsc" {
		fillUpTable[models.Umlsc](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Vertice" {
		fillUpTable[models.Vertice](nodeImplGongstruct.playground)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.playground.tableStage.Commit()
}