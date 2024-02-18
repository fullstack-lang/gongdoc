package adapter

import gong_models "github.com/fullstack-lang/gong/go/models"

type ElementNodeBase struct {
	stage *gong_models.StageStruct
}

func (base *ElementNodeBase) IsExpanded() bool {
	return false
}

func (base *ElementNodeBase) HasCheckboxButton() bool {
	return true
}

func (base *ElementNodeBase) IsNameEditable() bool {
	return false
}
