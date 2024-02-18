package node2gongdoc

import (
	"github.com/fullstack-lang/gongdoc/go/adapter"
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func FillUpNodeTree(
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
) {

	modelAdapter := adapter.NewModelAdapter(diagramPackage.ModelPkg.GetStage())
	portfolioAdapter := adapter.NewPortfolioAdapter(gongdocStage)
	diagrammer := diagrammer.NewDiagrammer(modelAdapter, portfolioAdapter, gongtreeStage)

	treeOfModelObjects := (&gongtree_models.Tree{Name: "model"}).Stage(gongtreeStage)
	diagrammer.FillUpModelTree(treeOfModelObjects)

	treeOfPortfolioObjects := (&gongtree_models.Tree{Name: "portfolio"}).Stage(gongtreeStage)
	diagrammer.FillUpPortfolioTree(treeOfPortfolioObjects)

	treeOfGongObjects := FillUpTreeOfGongObjects(gongdocStage, gongtreeStage, diagramPackage)
	rootOfClassdiagramsNode := FillUpTreeOfDiagramNodes(gongdocStage, gongtreeStage, diagramPackage, treeOfGongObjects)

	computeNodeConfs(
		gongtreeStage,
		gongdocStage,
		rootOfClassdiagramsNode,
		diagramPackage,
		treeOfGongObjects)

	gongdocStage.Commit()
	gongtreeStage.Commit()
}
