// generated code - do not edit
package probe

import (
	"embed"
	"sort"

	"github.com/gin-gonic/gin"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"
)

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	stackPath string,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct) {

	gongStage, _ := gong_fullstack.NewStackInstance(r, stackPath)

	gong_models.LoadEmbedded(gongStage, goModelsDir)

	// treeForSelectingDate that is on the sidebar
	stageForSidebarTree, _ := gongtree_fullstack.NewStackInstance(r, stackPath+"-sidebar")

	// stage for main table
	tableStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath)
	tableStage.Commit()

	// stage for reusable form
	formStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath+"-form")
	formStage.Commit()

	playground := (&Playground{
		stageOfInterest:    stageOfInterest,
		backRepoOfInterest: backRepoOfInterest,
		r:                  r,
		formStage:          formStage,
		tableStage:         tableStage,
	})

	// create tree
	treeOfGongStructs := (&gongtree_models.Tree{Name: "gong"}).Stage(stageForSidebarTree)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](gongStage)

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

		nodeGongstruct := (&gongtree_models.Node{Name: gongStruct.Name}).Stage(stageForSidebarTree)
		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(stageForSidebarTree)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			playground,
		)

		treeOfGongStructs.RootNodes = append(treeOfGongStructs.RootNodes, nodeGongstruct)
	}
	stageForSidebarTree.Commit()
}