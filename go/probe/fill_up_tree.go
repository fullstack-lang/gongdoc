package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range playground.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	playground.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(playground.treeStage)

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

		nodeGongstruct := (&tree.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}
		
		switch gongStruct.Name {
		// insertion point
		case "Classdiagram":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Classdiagram](playground.stageOfInterest)
			for _classdiagram := range set {
				nodeInstance := (&tree.Node{Name: _classdiagram.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_classdiagram, "Classdiagram", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DiagramPackage":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DiagramPackage](playground.stageOfInterest)
			for _diagrampackage := range set {
				nodeInstance := (&tree.Node{Name: _diagrampackage.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_diagrampackage, "DiagramPackage", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Field":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Field](playground.stageOfInterest)
			for _field := range set {
				nodeInstance := (&tree.Node{Name: _field.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_field, "Field", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumShape](playground.stageOfInterest)
			for _gongenumshape := range set {
				nodeInstance := (&tree.Node{Name: _gongenumshape.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumshape, "GongEnumShape", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongEnumValueEntry":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongEnumValueEntry](playground.stageOfInterest)
			for _gongenumvalueentry := range set {
				nodeInstance := (&tree.Node{Name: _gongenumvalueentry.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongenumvalueentry, "GongEnumValueEntry", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "GongStructShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.GongStructShape](playground.stageOfInterest)
			for _gongstructshape := range set {
				nodeInstance := (&tree.Node{Name: _gongstructshape.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_gongstructshape, "GongStructShape", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Link":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Link](playground.stageOfInterest)
			for _link := range set {
				nodeInstance := (&tree.Node{Name: _link.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_link, "Link", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "NoteShape":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.NoteShape](playground.stageOfInterest)
			for _noteshape := range set {
				nodeInstance := (&tree.Node{Name: _noteshape.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_noteshape, "NoteShape", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "NoteShapeLink":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.NoteShapeLink](playground.stageOfInterest)
			for _noteshapelink := range set {
				nodeInstance := (&tree.Node{Name: _noteshapelink.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_noteshapelink, "NoteShapeLink", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Position":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Position](playground.stageOfInterest)
			for _position := range set {
				nodeInstance := (&tree.Node{Name: _position.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_position, "Position", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "UmlState":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.UmlState](playground.stageOfInterest)
			for _umlstate := range set {
				nodeInstance := (&tree.Node{Name: _umlstate.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_umlstate, "UmlState", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Umlsc":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Umlsc](playground.stageOfInterest)
			for _umlsc := range set {
				nodeInstance := (&tree.Node{Name: _umlsc.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_umlsc, "Umlsc", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Vertice":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Vertice](playground.stageOfInterest)
			for _vertice := range set {
				nodeInstance := (&tree.Node{Name: _vertice.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_vertice, "Vertice", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&tree.Button{
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
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
