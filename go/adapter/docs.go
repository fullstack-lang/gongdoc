// Package adapter provides an implementation of the bridge interface
// for models defined in the gong "models" package.
//
// This package serves as an adapter
// allowing the models to be used wherever the bridge interface is required,
// facilitating a decoupled architecture between the system's core logic
// and its data models.
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
func (adapter *Adapter) GetRootNodes() (rootNodes []bridge.Node) {
	gongStructCategoryNode := NewGongStructCategoryNode(adapter.stage, "gongstructs")
	rootNodes = append(rootNodes, gongStructCategoryNode)

	return
}
