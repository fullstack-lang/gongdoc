package node2gongdoc

import (
	"log"
	"os"
	"path/filepath"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ButtonImplClassdiagram struct {
	// one need to access the daigramPackage to set the selected diagram
	diagramPackage *gongdoc_models.DiagramPackage
	classdiagram   *gongdoc_models.Classdiagram

	// one needs to access the node of the diagram package to manage the childern nodes
	diagramPackageNode *gongdoc_models.Node

	legacyDiagramPackageNode *gongdoc_models.Node

	// one needs to perform computation of node confs after the update
	treeOfGongObjects *gongdoc_models.Tree

	classdiagramNode     *gongdoc_models.Node
	nodeImplClassdiagram *NodeImplClasssiagram

	// type of button
	Icon ButtonType
}

func NewButtonImplClassdiagram(
	diagramPackage *gongdoc_models.DiagramPackage,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackageNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
	classdiagramNode *gongdoc_models.Node,
	nodeImplClassdiagram *NodeImplClasssiagram,
	icon ButtonType,
) (buttonImplClassdiagram *ButtonImplClassdiagram) {

	buttonImplClassdiagram = new(ButtonImplClassdiagram)

	buttonImplClassdiagram.diagramPackage = diagramPackage
	buttonImplClassdiagram.classdiagram = classdiagram
	buttonImplClassdiagram.diagramPackageNode = diagramPackageNode
	buttonImplClassdiagram.treeOfGongObjects = treeOfGongObjects
	buttonImplClassdiagram.classdiagramNode = classdiagramNode
	buttonImplClassdiagram.nodeImplClassdiagram = nodeImplClassdiagram
	buttonImplClassdiagram.Icon = icon

	return
}

func (buttonImplClassdiagram *ButtonImplClassdiagram) ButtonUpdated(
	gongdocStage *gongdoc_models.StageStruct,
	stageButton, front *gongdoc_models.Button) {

	log.Println("ButtonImplClassdiagramDraw, ButtonUpdated", front.Name)

	switch buttonImplClassdiagram.Icon {
	case BUTTON_draw:
		buttonImplClassdiagram.classdiagram.IsInDrawMode = true
		buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode = true
	case BUTTON_edit_off:
		buttonImplClassdiagram.classdiagram.IsInDrawMode = false
		buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode = false
	case BUTTON_save:
		buttonImplClassdiagram.classdiagram.IsInDrawMode = false
		buttonImplClassdiagram.nodeImplClassdiagram.IsInDrawMode = false

		// checkout in order to get the latest version of the diagram before
		// modifying it updated by the front
		gongdocStage.Checkout()
		gongdocStage.Unstage()
		gongdoc_models.StageBranch(gongdocStage, buttonImplClassdiagram.classdiagram)

		filepath := filepath.Join(
			filepath.Join(buttonImplClassdiagram.diagramPackage.AbsolutePathToDiagramPackage,
				"../diagrams"),
			buttonImplClassdiagram.classdiagram.Name) + ".go"
		file, err := os.Create(filepath)
		if err != nil {
			log.Fatal("Cannot open diagram file" + err.Error())
		}
		defer file.Close()

		mapDocLinkRemaping := &gongdocStage.Map_DocLink_Renaming
		_ = mapDocLinkRemaping

		gongdoc_models.SetupMapDocLinkRenaming(buttonImplClassdiagram.diagramPackage.ModelPkg.Stage_, gongdocStage)

		gongdocStage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

		// restore the original stage
		gongdocStage.Unstage()
		gongdocStage.Checkout()
		gongdocStage.Commit()
	}

	computeNodeConfs(gongdocStage,
		buttonImplClassdiagram.diagramPackageNode,
		buttonImplClassdiagram.diagramPackage,
		buttonImplClassdiagram.treeOfGongObjects)
}
