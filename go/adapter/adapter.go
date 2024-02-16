package adapter

import (
	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdoc/go/bridge"
)

type Adapter struct {
	stage *gong_models.StageStruct
}

func NewAdapter(stage *gong_models.StageStruct) (adapter *Adapter) {
	adapter = new(Adapter)
	adapter.stage = stage
	return
}

// GetRootNodes implements bridge.Model.
func (adapter *Adapter) GetChildren() (rootNodes []bridge.Node) {
	gongStructCategoryNode := NewGongStructCategoryNode(adapter.stage, "gongstructs")
	rootNodes = append(rootNodes, gongStructCategoryNode)

	gongEnumCategoryNode := NewGongEnumCategoryNode(adapter.stage, "gongenums")
	rootNodes = append(rootNodes, gongEnumCategoryNode)

	gongNoteCategoryNode := NewGongNoteCategoryNode(adapter.stage, "gongnotes")
	rootNodes = append(rootNodes, gongNoteCategoryNode)

	return
}
