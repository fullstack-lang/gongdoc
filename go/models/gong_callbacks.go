package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Classdiagram:
		if stage.OnAfterClassdiagramCreateCallback != nil {
			stage.OnAfterClassdiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Classshape:
		if stage.OnAfterClassshapeCreateCallback != nil {
			stage.OnAfterClassshapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageCreateCallback != nil {
			stage.OnAfterDiagramPackageCreateCallback.OnAfterCreate(stage, target)
		}
	case *Field:
		if stage.OnAfterFieldCreateCallback != nil {
			stage.OnAfterFieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongStruct:
		if stage.OnAfterGongStructCreateCallback != nil {
			stage.OnAfterGongStructCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongdocCommand:
		if stage.OnAfterGongdocCommandCreateCallback != nil {
			stage.OnAfterGongdocCommandCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongdocStatus:
		if stage.OnAfterGongdocStatusCreateCallback != nil {
			stage.OnAfterGongdocStatusCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeCreateCallback != nil {
			stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Position:
		if stage.OnAfterPositionCreateCallback != nil {
			stage.OnAfterPositionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeCreateCallback != nil {
			stage.OnAfterTreeCreateCallback.OnAfterCreate(stage, target)
		}
	case *UmlState:
		if stage.OnAfterUmlStateCreateCallback != nil {
			stage.OnAfterUmlStateCreateCallback.OnAfterCreate(stage, target)
		}
	case *Umlsc:
		if stage.OnAfterUmlscCreateCallback != nil {
			stage.OnAfterUmlscCreateCallback.OnAfterCreate(stage, target)
		}
	case *Vertice:
		if stage.OnAfterVerticeCreateCallback != nil {
			stage.OnAfterVerticeCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Classdiagram:
		newTarget := any(new).(*Classdiagram)
		if stage.OnAfterClassdiagramUpdateCallback != nil {
			stage.OnAfterClassdiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Classshape:
		newTarget := any(new).(*Classshape)
		if stage.OnAfterClassshapeUpdateCallback != nil {
			stage.OnAfterClassshapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramPackage:
		newTarget := any(new).(*DiagramPackage)
		if stage.OnAfterDiagramPackageUpdateCallback != nil {
			stage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Field:
		newTarget := any(new).(*Field)
		if stage.OnAfterFieldUpdateCallback != nil {
			stage.OnAfterFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongStruct:
		newTarget := any(new).(*GongStruct)
		if stage.OnAfterGongStructUpdateCallback != nil {
			stage.OnAfterGongStructUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongdocCommand:
		newTarget := any(new).(*GongdocCommand)
		if stage.OnAfterGongdocCommandUpdateCallback != nil {
			stage.OnAfterGongdocCommandUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongdocStatus:
		newTarget := any(new).(*GongdocStatus)
		if stage.OnAfterGongdocStatusUpdateCallback != nil {
			stage.OnAfterGongdocStatusUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Node:
		newTarget := any(new).(*Node)
		if stage.OnAfterNodeUpdateCallback != nil {
			stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note:
		newTarget := any(new).(*Note)
		if stage.OnAfterNoteUpdateCallback != nil {
			stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Position:
		newTarget := any(new).(*Position)
		if stage.OnAfterPositionUpdateCallback != nil {
			stage.OnAfterPositionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tree:
		newTarget := any(new).(*Tree)
		if stage.OnAfterTreeUpdateCallback != nil {
			stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *UmlState:
		newTarget := any(new).(*UmlState)
		if stage.OnAfterUmlStateUpdateCallback != nil {
			stage.OnAfterUmlStateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Umlsc:
		newTarget := any(new).(*Umlsc)
		if stage.OnAfterUmlscUpdateCallback != nil {
			stage.OnAfterUmlscUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vertice:
		newTarget := any(new).(*Vertice)
		if stage.OnAfterVerticeUpdateCallback != nil {
			stage.OnAfterVerticeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Classdiagram:
		if stage.OnAfterClassdiagramDeleteCallback != nil {
			staged := any(staged).(*Classdiagram)
			stage.OnAfterClassdiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Classshape:
		if stage.OnAfterClassshapeDeleteCallback != nil {
			staged := any(staged).(*Classshape)
			stage.OnAfterClassshapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageDeleteCallback != nil {
			staged := any(staged).(*DiagramPackage)
			stage.OnAfterDiagramPackageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Field:
		if stage.OnAfterFieldDeleteCallback != nil {
			staged := any(staged).(*Field)
			stage.OnAfterFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongStruct:
		if stage.OnAfterGongStructDeleteCallback != nil {
			staged := any(staged).(*GongStruct)
			stage.OnAfterGongStructDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongdocCommand:
		if stage.OnAfterGongdocCommandDeleteCallback != nil {
			staged := any(staged).(*GongdocCommand)
			stage.OnAfterGongdocCommandDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongdocStatus:
		if stage.OnAfterGongdocStatusDeleteCallback != nil {
			staged := any(staged).(*GongdocStatus)
			stage.OnAfterGongdocStatusDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Node:
		if stage.OnAfterNodeDeleteCallback != nil {
			staged := any(staged).(*Node)
			stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note:
		if stage.OnAfterNoteDeleteCallback != nil {
			staged := any(staged).(*Note)
			stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Position:
		if stage.OnAfterPositionDeleteCallback != nil {
			staged := any(staged).(*Position)
			stage.OnAfterPositionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tree:
		if stage.OnAfterTreeDeleteCallback != nil {
			staged := any(staged).(*Tree)
			stage.OnAfterTreeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *UmlState:
		if stage.OnAfterUmlStateDeleteCallback != nil {
			staged := any(staged).(*UmlState)
			stage.OnAfterUmlStateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Umlsc:
		if stage.OnAfterUmlscDeleteCallback != nil {
			staged := any(staged).(*Umlsc)
			stage.OnAfterUmlscDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vertice:
		if stage.OnAfterVerticeDeleteCallback != nil {
			staged := any(staged).(*Vertice)
			stage.OnAfterVerticeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Classdiagram:
		if stage.OnAfterClassdiagramReadCallback != nil {
			stage.OnAfterClassdiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Classshape:
		if stage.OnAfterClassshapeReadCallback != nil {
			stage.OnAfterClassshapeReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageReadCallback != nil {
			stage.OnAfterDiagramPackageReadCallback.OnAfterRead(stage, target)
		}
	case *Field:
		if stage.OnAfterFieldReadCallback != nil {
			stage.OnAfterFieldReadCallback.OnAfterRead(stage, target)
		}
	case *GongStruct:
		if stage.OnAfterGongStructReadCallback != nil {
			stage.OnAfterGongStructReadCallback.OnAfterRead(stage, target)
		}
	case *GongdocCommand:
		if stage.OnAfterGongdocCommandReadCallback != nil {
			stage.OnAfterGongdocCommandReadCallback.OnAfterRead(stage, target)
		}
	case *GongdocStatus:
		if stage.OnAfterGongdocStatusReadCallback != nil {
			stage.OnAfterGongdocStatusReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeReadCallback != nil {
			stage.OnAfterNodeReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *Position:
		if stage.OnAfterPositionReadCallback != nil {
			stage.OnAfterPositionReadCallback.OnAfterRead(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeReadCallback != nil {
			stage.OnAfterTreeReadCallback.OnAfterRead(stage, target)
		}
	case *UmlState:
		if stage.OnAfterUmlStateReadCallback != nil {
			stage.OnAfterUmlStateReadCallback.OnAfterRead(stage, target)
		}
	case *Umlsc:
		if stage.OnAfterUmlscReadCallback != nil {
			stage.OnAfterUmlscReadCallback.OnAfterRead(stage, target)
		}
	case *Vertice:
		if stage.OnAfterVerticeReadCallback != nil {
			stage.OnAfterVerticeReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Classdiagram])
	
	case *Classshape:
		stage.OnAfterClassshapeUpdateCallback = any(callback).(OnAfterUpdateInterface[Classshape])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldUpdateCallback = any(callback).(OnAfterUpdateInterface[Field])
	
	case *GongStruct:
		stage.OnAfterGongStructUpdateCallback = any(callback).(OnAfterUpdateInterface[GongStruct])
	
	case *GongdocCommand:
		stage.OnAfterGongdocCommandUpdateCallback = any(callback).(OnAfterUpdateInterface[GongdocCommand])
	
	case *GongdocStatus:
		stage.OnAfterGongdocStatusUpdateCallback = any(callback).(OnAfterUpdateInterface[GongdocStatus])
	
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	
	case *Node:
		stage.OnAfterNodeUpdateCallback = any(callback).(OnAfterUpdateInterface[Node])
	
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	
	case *Position:
		stage.OnAfterPositionUpdateCallback = any(callback).(OnAfterUpdateInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeUpdateCallback = any(callback).(OnAfterUpdateInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateUpdateCallback = any(callback).(OnAfterUpdateInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscUpdateCallback = any(callback).(OnAfterUpdateInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeUpdateCallback = any(callback).(OnAfterUpdateInterface[Vertice])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramCreateCallback = any(callback).(OnAfterCreateInterface[Classdiagram])
	
	case *Classshape:
		stage.OnAfterClassshapeCreateCallback = any(callback).(OnAfterCreateInterface[Classshape])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageCreateCallback = any(callback).(OnAfterCreateInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldCreateCallback = any(callback).(OnAfterCreateInterface[Field])
	
	case *GongStruct:
		stage.OnAfterGongStructCreateCallback = any(callback).(OnAfterCreateInterface[GongStruct])
	
	case *GongdocCommand:
		stage.OnAfterGongdocCommandCreateCallback = any(callback).(OnAfterCreateInterface[GongdocCommand])
	
	case *GongdocStatus:
		stage.OnAfterGongdocStatusCreateCallback = any(callback).(OnAfterCreateInterface[GongdocStatus])
	
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	
	case *Node:
		stage.OnAfterNodeCreateCallback = any(callback).(OnAfterCreateInterface[Node])
	
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	
	case *Position:
		stage.OnAfterPositionCreateCallback = any(callback).(OnAfterCreateInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeCreateCallback = any(callback).(OnAfterCreateInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateCreateCallback = any(callback).(OnAfterCreateInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscCreateCallback = any(callback).(OnAfterCreateInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeCreateCallback = any(callback).(OnAfterCreateInterface[Vertice])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Classdiagram])
	
	case *Classshape:
		stage.OnAfterClassshapeDeleteCallback = any(callback).(OnAfterDeleteInterface[Classshape])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldDeleteCallback = any(callback).(OnAfterDeleteInterface[Field])
	
	case *GongStruct:
		stage.OnAfterGongStructDeleteCallback = any(callback).(OnAfterDeleteInterface[GongStruct])
	
	case *GongdocCommand:
		stage.OnAfterGongdocCommandDeleteCallback = any(callback).(OnAfterDeleteInterface[GongdocCommand])
	
	case *GongdocStatus:
		stage.OnAfterGongdocStatusDeleteCallback = any(callback).(OnAfterDeleteInterface[GongdocStatus])
	
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	
	case *Node:
		stage.OnAfterNodeDeleteCallback = any(callback).(OnAfterDeleteInterface[Node])
	
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	
	case *Position:
		stage.OnAfterPositionDeleteCallback = any(callback).(OnAfterDeleteInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeDeleteCallback = any(callback).(OnAfterDeleteInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateDeleteCallback = any(callback).(OnAfterDeleteInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscDeleteCallback = any(callback).(OnAfterDeleteInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeDeleteCallback = any(callback).(OnAfterDeleteInterface[Vertice])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramReadCallback = any(callback).(OnAfterReadInterface[Classdiagram])
	
	case *Classshape:
		stage.OnAfterClassshapeReadCallback = any(callback).(OnAfterReadInterface[Classshape])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageReadCallback = any(callback).(OnAfterReadInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldReadCallback = any(callback).(OnAfterReadInterface[Field])
	
	case *GongStruct:
		stage.OnAfterGongStructReadCallback = any(callback).(OnAfterReadInterface[GongStruct])
	
	case *GongdocCommand:
		stage.OnAfterGongdocCommandReadCallback = any(callback).(OnAfterReadInterface[GongdocCommand])
	
	case *GongdocStatus:
		stage.OnAfterGongdocStatusReadCallback = any(callback).(OnAfterReadInterface[GongdocStatus])
	
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	
	case *Node:
		stage.OnAfterNodeReadCallback = any(callback).(OnAfterReadInterface[Node])
	
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	
	case *Position:
		stage.OnAfterPositionReadCallback = any(callback).(OnAfterReadInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeReadCallback = any(callback).(OnAfterReadInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateReadCallback = any(callback).(OnAfterReadInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscReadCallback = any(callback).(OnAfterReadInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeReadCallback = any(callback).(OnAfterReadInterface[Vertice])
	
	}
}