package adapter

import gong_models "github.com/fullstack-lang/gong/go/models"

type ModelCategoryNodeBase struct {
	stage      *gong_models.StageStruct
	Name       string
	isExpanded bool
}

func (base *ModelCategoryNodeBase) IsExpanded() bool {
	return true
}

func (base *ModelCategoryNodeBase) SetIsExpanded(isExpanded bool) {
	base.isExpanded = isExpanded
}

func (base *ModelCategoryNodeBase) IsNameEditable() bool {
	return false
}

func (base *ModelCategoryNodeBase) GetName() string {
	return base.Name
}
