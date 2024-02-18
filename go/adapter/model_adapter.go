package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/diagrammer"
)

type ModelAdapter struct {
	stage *gong_models.StageStruct
}

func NewModelAdapter(stage *gong_models.StageStruct) (adapter *ModelAdapter) {
	adapter = new(ModelAdapter)
	adapter.stage = stage
	return
}

// GetRootNodes implements bridge.Model.
func (modelAdapter *ModelAdapter) GetChildren() (rootNodes []diagrammer.ModelNode) {
	gongStructCategoryNode := NewGongStructCategoryNode(modelAdapter.stage, "gongstructs")
	rootNodes = append(rootNodes, gongStructCategoryNode)

	gongEnumCategoryNode := NewGongEnumCategoryNode(modelAdapter.stage, "gongenums")
	rootNodes = append(rootNodes, gongEnumCategoryNode)

	gongNoteCategoryNode := NewGongNoteCategoryNode(modelAdapter.stage, "gongnotes")
	rootNodes = append(rootNodes, gongNoteCategoryNode)

	return
}
