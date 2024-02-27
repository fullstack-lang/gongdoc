package adapter

import (
	"github.com/fullstack-lang/gongdoc/go/diagrammer"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type ClassdiagramAdapter struct {
	classdiagram *gongdoc_models.Classdiagram
}

func (classdiagramAdapter *ClassdiagramAdapter) GetShapes() []diagrammer.Shape {
	return nil
}

func (classdiagramAdapter *ClassdiagramAdapter) GetName() string {
	return classdiagramAdapter.classdiagram.GetName()
}

func (classdiagramAdapter *ClassdiagramAdapter) GetClassdiagral() *gongdoc_models.Classdiagram {
	return classdiagramAdapter.classdiagram
}
