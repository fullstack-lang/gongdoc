// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdoc/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.playground = playground
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Classdiagram:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.DiagramPackage:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Field:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongEnumShape:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongEnumValueEntry:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongStructShape:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Link:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.NoteShape:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.NoteShapeLink:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Position:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.UmlState:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Umlsc:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Vertice:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.playground.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.playground)
	fillUpTree(cellDeleteIconImpl.playground)
	cellDeleteIconImpl.playground.tableStage.Commit()
}

