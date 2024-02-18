package adapter

import gong_models "github.com/fullstack-lang/gong/go/models"

type CategoryNodeBase struct {
	stage *gong_models.StageStruct
	Name  string
}

func (base *CategoryNodeBase) IsExpanded() bool {
	return true
}

func (base *CategoryNodeBase) HasCheckboxButton() bool {
	return false
}

func (base *CategoryNodeBase) IsNameEditable() bool {
	return false
}

func (base *CategoryNodeBase) GetName() string {
	return base.Name
}
