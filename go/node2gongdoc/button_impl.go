package node2gongdoc

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ButtonImpl struct {
}

func (buttonImpl *ButtonImpl) ButtonUpdated(
	stage *gongdoc_models.StageStruct,
	front *gongdoc_models.Button) {

	log.Println("ButtonImpl, ButtonUpdated", front.Name)

}
