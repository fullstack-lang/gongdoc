package models

import "log"

type ClassshapeCallbacksSingloton struct {
	ClassshapeCallback ClassshapeCallback
}

func (classshapeCallbacksSingloton *ClassshapeCallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	stagedClassshape, frontClassshape *Classshape) {

	if stagedClassshape.IsSelected != frontClassshape.IsSelected {
		log.Println("UML Shape selected ", stagedClassshape.GongStruct.Name)
		gongStruct, ok := Stage.GongStructs_mapString[stagedClassshape.GongStruct.Name]
		if ok {
			if classshapeCallbacksSingloton.ClassshapeCallback != nil {
				classshapeCallbacksSingloton.ClassshapeCallback.HasSelected(gongStruct.Name)
			}
		}
	}
}

type ClassshapeCallback interface {
	HasSelected(gongstructName string)
}
