package models

import "log"

type ClassshapeCallbacksSingloton struct {
	ClassshapeCallback ClassshapeCallback
}

func (classshapeCallbacksSingloton *ClassshapeCallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	stagedClassshape, frontClassshape *Classshape) {

	if stagedClassshape.IsSelected != frontClassshape.IsSelected {

		// reset the IsSelected to false
		stagedClassshape.Checkout()
		stagedClassshape.IsSelected = false
		stagedClassshape.Commit()

		log.Println("UML Shape selected ", stagedClassshape.Identifier)
		classshapeCallbacksSingloton.ClassshapeCallback.HasSelected(stagedClassshape.Identifier)
	}
}

type ClassshapeCallback interface {
	HasSelected(gongstructName string)
}
