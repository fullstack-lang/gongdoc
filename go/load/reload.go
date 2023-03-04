package load

import (
	"path/filepath"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_node2gongdoc "github.com/fullstack-lang/gongdoc/go/node2gongdoc"
)

func Reload(diagramPackage *gongdoc_models.DiagramPackage) {

	gong_models.Stage.Checkout()
	gong_models.Stage.Reset()
	modelPkg, _ := gong_models.LoadSource(
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../models"))
	gong_models.Stage.Commit()

	diagramPackage.Stage_.Checkout()
	diagramPackage.Stage_.Reset()
	diagramPackage.Stage_.Commit()

	diagramPackage.Classdiagrams = nil
	diagramPackage.Umlscs = nil
	diagramPackage.ModelPkg = modelPkg

	diagramPackage, _ = LoadDiagramPackage(
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../models"),
		modelPkg, true)

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	gongdoc_models.SetupMapDocLinkRenaming(diagramPackage.Stage_)
	// end of the be removed
	gongdoc_node2gongdoc.FillUpNodeTree(diagramPackage)
	diagramPackage.Stage_.Commit()
}
