package probe

import (
	"fmt"
	"sort"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	playground.treeStage.Reset()

	// create tree
	sidebar := (&gongtree_models.Tree{Name: "gong"}).Stage(playground.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](playground.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", playground.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&gongtree_models.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		switch gongStruct.Name {
		// insertion point
		case "Classdiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Classdiagram](playground.stageOfInterest)
			for classdiagram := range set {
				nodeInstance := (&gongtree_models.Node{Name: classdiagram.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(classdiagram, "Classdiagram", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DiagramPackage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DiagramPackage](playground.stageOfInterest)
			for diagrampackage := range set {
				nodeInstance := (&gongtree_models.Node{Name: diagrampackage.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(diagrampackage, "DiagramPackage", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Field":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Field](playground.stageOfInterest)
			for field := range set {
				nodeInstance := (&gongtree_models.Node{Name: field.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(field, "Field", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumShape](playground.stageOfInterest)
			for gongenumshape := range set {
				nodeInstance := (&gongtree_models.Node{Name: gongenumshape.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(gongenumshape, "GongEnumShape", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumValueEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumValueEntry](playground.stageOfInterest)
			for gongenumvalueentry := range set {
				nodeInstance := (&gongtree_models.Node{Name: gongenumvalueentry.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(gongenumvalueentry, "GongEnumValueEntry", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongStructShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongStructShape](playground.stageOfInterest)
			for gongstructshape := range set {
				nodeInstance := (&gongtree_models.Node{Name: gongstructshape.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(gongstructshape, "GongStructShape", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Link](playground.stageOfInterest)
			for link := range set {
				nodeInstance := (&gongtree_models.Node{Name: link.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(link, "Link", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "NoteShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.NoteShape](playground.stageOfInterest)
			for noteshape := range set {
				nodeInstance := (&gongtree_models.Node{Name: noteshape.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(noteshape, "NoteShape", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "NoteShapeLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.NoteShapeLink](playground.stageOfInterest)
			for noteshapelink := range set {
				nodeInstance := (&gongtree_models.Node{Name: noteshapelink.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(noteshapelink, "NoteShapeLink", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Position":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Position](playground.stageOfInterest)
			for position := range set {
				nodeInstance := (&gongtree_models.Node{Name: position.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(position, "Position", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "UmlState":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.UmlState](playground.stageOfInterest)
			for umlstate := range set {
				nodeInstance := (&gongtree_models.Node{Name: umlstate.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(umlstate, "UmlState", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Umlsc":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Umlsc](playground.stageOfInterest)
			for umlsc := range set {
				nodeInstance := (&gongtree_models.Node{Name: umlsc.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(umlsc, "Umlsc", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Vertice":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Vertice](playground.stageOfInterest)
			for vertice := range set {
				nodeInstance := (&gongtree_models.Node{Name: vertice.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(vertice, "Vertice", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(playground.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			playground,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}
	playground.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	playground     *Playground
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	playground *Playground) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.playground = playground
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
