package node2gongdoc

import (
	"fmt"
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// NodeCB is the singloton callback implementation of CUD operations on node
type NodeCB struct {

	// diagramPackageNode is used for iterating on classdiagram nodes
	// and for getting the selected diagram
	diagramPackageNode *gongdoc_models.Node
	diagramPackage     *gongdoc_models.DiagramPackage

	// treeOfGongObjects is the root of all nodes related to gong objects
	// it is necessary for performing operation on all elements of the tree
	treeOfGongObjects *gongdoc_models.Tree

	// map_Children_Parent is a map to navigate from a children node to its parent node
	// it is set up once at the init phase
	map_Children_Parent map[*gongdoc_models.Node]*gongdoc_models.Node

	// map_Identifier_Node is a map to navigate from identifiers in the package
	// to nodes, and backforth
	// identifiers are unique in a package (that's the point of identifiers)
	// TO BE CHANGED, both a NoteLink and a GongField can have the same identifier
	// New form, a impl field to navigate from the node to the shape
	// a node field to navigate from a shape to the corresponding node
	map_Identifier_Node map[string]*gongdoc_models.Node

	map_gongObject_gongObjectImpl map[any]gongdoc_models.NodeImplInterface
}

// GetSelectedClassdiagram
func (nodesCb *NodeCB) GetSelectedClassdiagram() (classdiagram *gongdoc_models.Classdiagram) {

	classdiagram = nodesCb.diagramPackage.SelectedClassdiagram
	return
}

// OnAfterUpdate is called each time the end user interacts
// with any node. The front commit the state of the front node
// to the back ([frontNode] in the function).
// Therefore, the [stageNode] is now different from the [frontNode].
//
// This functiion and its siblings analyse which field of the
// node has changed and performs all necessary actions
func (nodeCb *NodeCB) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	stagedNode.Impl.OnAfterUpdate(stage, stagedNode, frontNode)
	nodeCb.updateNodesStates(stage)

	if stagedNode.IsExpanded != frontNode.IsExpanded {
		// setting the value of the staged node	to the new value
		stagedNode.IsExpanded = frontNode.IsExpanded
	}
}

// OnAfterCreate is another callback
func (nodeCb *NodeCB) OnAfterCreate(
	stage *gongdoc_models.StageStruct,
	node *gongdoc_models.Node) {

	log.Println("Node " + node.Name + " is created")

	node.HasCheckboxButton = true

	// check unicity of name, otherwise, add an index
	var hasNameCollision bool
	initialName := node.Name
	index := 0
	// loop until the name of the new diagram is not in collision with an existing
	// diagram name
	for index == 0 || hasNameCollision {
		index++
		hasNameCollision = false
		for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram]() {
			if classdiagram.Name == node.Name {
				hasNameCollision = true
			}
		}
		if hasNameCollision {
			node.Name = initialName + fmt.Sprintf("_%d", index)
		}
	}

	classdiagram := (&gongdoc_models.Classdiagram{Name: node.Name}).Stage()
	nodeCb.diagramPackage.Classdiagrams = append(nodeCb.diagramPackage.Classdiagrams, classdiagram)
	node.IsInEditMode = false
	node.IsInDrawMode = false
	node.HasEditButton = false

	nodeCb.diagramPackageNode.Children = append(nodeCb.diagramPackageNode.Children, node)

	// set up the back pointer from the shape to the node
	classdiagramImpl := new(ClassdiagramImpl)
	classdiagramImpl.node = node
	classdiagramImpl.classdiagram = classdiagram
	classdiagramImpl.nodeCb = nodeCb
	node.Impl = classdiagramImpl

	nodeCb.updateNodesStates(stage)

}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (nodeCb *NodeCB) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	switch impl := stagedNode.Impl.(type) {
	case *ClassdiagramImpl:
		impl.OnAfterDelete(stage, stagedNode, frontNode)
	}

	nodeCb.updateNodesStates(stage)
}

func (nodeCb *NodeCB) FillUpDiagramNodeTree(diagramPackage *gongdoc_models.DiagramPackage) {

	// generate tree of diagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc"}).Stage()

	// add the root of class diagrams
	diagramPackageNode := (&gongdoc_models.Node{Name: "class diagrams", Type: gongdoc_models.ROOT_OF_CLASS_DIAGRAMS}).Stage()
	diagramPackageNode.IsExpanded = true
	diagramPackageNode.HasAddChildButton = diagramPackage.IsEditable
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, diagramPackageNode)

	// add one node per class diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram]() {
		node := (&gongdoc_models.Node{Name: classdiagram.Name}).Stage()
		node.Type = gongdoc_models.CLASS_DIAGRAM
		node.HasCheckboxButton = true
		node.HasDeleteButton = true
		node.HasEditButton = true
		diagramPackageNode.Children = append(diagramPackageNode.Children, node)

		// set up the back pointer from the shape to the node
		classdiagramImpl := new(ClassdiagramImpl)
		classdiagramImpl.node = node
		classdiagramImpl.classdiagram = classdiagram
		classdiagramImpl.nodeCb = nodeCb
		node.Impl = classdiagramImpl
	}

	// set callbacks on node updates
	nodeCb.diagramPackageNode = diagramPackageNode
}

func SetNodeBackPointer[T1 gong_models.Gongstruct](gong_instance *T1, backPointer gongdoc_models.NodeImplInterface) {
	gong_models.SetBackPointer(&gong_models.Stage, gong_instance, backPointer)
}
func GetNodeBackPointer[T1 gong_models.Gongstruct](gong_instance *T1) (backPointer gongdoc_models.NodeImplInterface) {
	tmp := gong_models.GetBackPointer(&gong_models.Stage, gong_instance)

	backPointer = tmp.(gongdoc_models.NodeImplInterface)

	return
}

func (nodeCb *NodeCB) FillUpTreeOfGongObjects(pkgelt *gongdoc_models.DiagramPackage) {

	// set up the gongTree to display elements
	gongTree := (&gongdoc_models.Tree{Name: "gong"}).Stage()
	nodeCb.treeOfGongObjects = gongTree

	nodeCb.map_Identifier_Node = make(map[string]*gongdoc_models.Node)
	nodeCb.map_gongObject_gongObjectImpl = make(map[any]gongdoc_models.NodeImplInterface)

	gongstructRootNode := (&gongdoc_models.Node{Name: "gongstructs"}).Stage()
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		nodeGongstruct := (&gongdoc_models.Node{Name: gongStruct.Name}).Stage()
		nodeGongstruct.Type = gongdoc_models.GONG_STRUCT
		nodeGongstruct.HasCheckboxButton = true
		nodeGongstruct.IsExpanded = false

		// set up the back pointer from the shape to the node
		gongStructImpl := new(GongStructImpl)
		gongStructImpl.node = nodeGongstruct
		gongStructImpl.gongStruct = gongStruct
		gongStructImpl.nodeCb = nodeCb
		nodeGongstruct.Impl = gongStructImpl

		// append to the tree
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)
		nodeCb.map_Identifier_Node[gongdoc_models.ShapenameToIdentifier(gongStruct.Name)] = nodeGongstruct
		nodeCb.map_gongObject_gongObjectImpl[gongStruct] = gongStructImpl
		SetNodeBackPointer(gongStruct, gongStructImpl)

		for _, field := range gongStruct.Fields {
			nodeGongField := (&gongdoc_models.Node{Name: field.GetName()}).Stage()
			nodeGongField.Type = gongdoc_models.GONG_STRUCT_FIELD
			nodeGongField.HasCheckboxButton = true

			fieldImpl := new(FieldImpl)
			fieldImpl.node = nodeGongstruct
			fieldImpl.field = field
			fieldImpl.nodeCb = nodeCb
			nodeGongField.Impl = fieldImpl

			// append to tree
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeGongField)
			nodeCb.map_Identifier_Node[gongdoc_models.ShapeAndFieldnameToFieldIdentifier(gongStruct.Name, field.GetName())] = nodeGongField
			nodeCb.map_gongObject_gongObjectImpl[field] = fieldImpl
		}
	}

	gongenumRootNode := (&gongdoc_models.Node{Name: "gongenums"}).Stage()
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {

		node := (&gongdoc_models.Node{Name: gongEnum.Name}).Stage()
		node.HasCheckboxButton = true
		node.IsExpanded = false
		node.Type = gongdoc_models.GONG_ENUM

		gongEnumImpl := new(GongEnumImpl)
		gongEnumImpl.node = node
		gongEnumImpl.gongEnum = gongEnum
		gongEnumImpl.nodeCb = nodeCb
		node.Impl = gongEnumImpl

		// append to tree
		gongenumRootNode.Children = append(gongenumRootNode.Children, node)
		nodeCb.map_Identifier_Node[gongdoc_models.ShapenameToIdentifier(gongEnum.Name)] = node
		nodeCb.map_gongObject_gongObjectImpl[gongEnum] = gongEnumImpl

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			nodeGongEnumValue := (&gongdoc_models.Node{Name: gongEnumValue.GetName()}).Stage()
			nodeGongEnumValue.Type = gongdoc_models.GONG_ENUM_VALUE
			nodeGongEnumValue.HasCheckboxButton = true

			gongEnumValueImpl := new(GongEnumValueImpl)
			gongEnumValueImpl.node = node
			gongEnumValueImpl.gongEnumValue = gongEnumValue
			gongEnumValueImpl.nodeCb = nodeCb
			nodeGongEnumValue.Impl = gongEnumValueImpl

			// append to tree
			node.Children = append(node.Children, nodeGongEnumValue)
			nodeCb.map_Identifier_Node[gongdoc_models.ShapeAndFieldnameToFieldIdentifier(gongEnum.Name, gongEnumValue.GetName())] = nodeGongEnumValue
			nodeCb.map_gongObject_gongObjectImpl[gongEnumValue] = gongEnumValueImpl
		}
	}

	gongNotesRootNode := (&gongdoc_models.Node{Name: "notes"}).Stage()
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {

		node := (&gongdoc_models.Node{Name: gongNote.Name}).Stage()
		node.HasCheckboxButton = true
		node.Type = gongdoc_models.GONG_NOTE
		node.IsExpanded = true

		gongNoteImpl := new(GongNoteImpl)
		gongNoteImpl.node = node
		gongNoteImpl.gongNote = gongNote
		gongNoteImpl.nodeCb = nodeCb
		node.Impl = gongNoteImpl

		// append to tree
		gongNotesRootNode.Children = append(gongNotesRootNode.Children, node)
		nodeCb.map_Identifier_Node[gongdoc_models.ShapenameToIdentifier(gongNote.Name)] = node
		nodeCb.map_gongObject_gongObjectImpl[gongNote] = gongNoteImpl

		for _, gongLink := range gongNote.Links {
			nodeGongLink := (&gongdoc_models.Node{Name: gongLink.Name}).Stage()
			nodeGongLink.HasCheckboxButton = true
			nodeGongLink.Type = gongdoc_models.GONG_NOTE_LINK

			gongLinkImpl := new(GongLinkImpl)
			gongLinkImpl.node = node
			gongLinkImpl.gongLink = gongLink
			gongLinkImpl.nodeCb = nodeCb
			nodeGongLink.Impl = gongLinkImpl

			// append to tree
			node.Children = append(node.Children, nodeGongLink)
			nodeCb.map_Identifier_Node[gongdoc_models.ShapeAndFieldnameToFieldIdentifier(gongNote.Name, gongLink.Name)] = nodeGongLink
			nodeCb.map_gongObject_gongObjectImpl[gongLink] = gongLinkImpl
		}
	}

	// generate the map to navigate from children to parents
	fieldName := gongdoc_models.GetAssociationName[gongdoc_models.Node]().Children[0].Name
	nodeCb.map_Children_Parent = gongdoc_models.GetSliceOfPointersReverseMap[gongdoc_models.Node, gongdoc_models.Node](fieldName)
}

func (nodesCb *NodeCB) updateDiagramsNodes(stage *gongdoc_models.StageStruct) {

	// compute wether one of the diagrams is in draw/edit mode
	// if so, all diagram check need to be disabled
	var inModificationMode bool
	for _, classdiagramNode := range nodesCb.diagramPackageNode.Children {
		if classdiagramNode.IsInDrawMode || classdiagramNode.IsInEditMode {
			inModificationMode = true
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range nodesCb.diagramPackageNode.Children {

		// reset the state of the classdiagram node
		classdiagramNode.HasEditButton = false
		classdiagramNode.HasDeleteButton = false
		classdiagramNode.HasDrawButton = false
		classdiagramNode.HasDrawOffButton = false

		classdiagramNode.IsCheckboxDisabled = inModificationMode

		if !classdiagramNode.IsChecked {
			classdiagramNode.IsInEditMode = false
			classdiagramNode.IsInDrawMode = false
			continue
		}

		editable := nodesCb.diagramPackage.IsEditable && !classdiagramNode.IsInEditMode && !classdiagramNode.IsInDrawMode

		classdiagramNode.HasEditButton = editable
		classdiagramNode.HasDeleteButton = editable
		classdiagramNode.HasDrawButton = editable
	}
}

func (nodeCb *NodeCB) updateGongObjectsNodes(stage *gongdoc_models.StageStruct, classdiagram *gongdoc_models.Classdiagram) {

	// parse gong object nodes and disable the check box
	for _, rootNode := range nodeCb.treeOfGongObjects.RootNodes {
		for _, node := range rootNode.Children {
			node.IsCheckboxDisabled = !classdiagram.IsInDrawMode
			for _, node := range node.Children {
				node.IsCheckboxDisabled = !classdiagram.IsInDrawMode
			}
		}
	}

	// FIRST STAGE : for each identifiers node with a related visual element is in
	// the classdiagram: check the node and enable it if the diagram is in drawMode

	// 1. gongstructs / gongenum referenced by the classshape
	for _, gongStructShape := range classdiagram.GongStructShapes {

		var gongStructShapeNode *gongdoc_models.Node
		var ok bool
		gongStructShapeNode, ok = nodeCb.map_Identifier_Node[gongStructShape.Identifier]

		if !ok {
			log.Println("Unknown node, diagram not synchro with model ", gongStructShape.Identifier)
			continue
		}

		// get the gong object from the identifier
		gongStructName := gongdoc_models.IdentifierToShapename(gongStructShape.Identifier)
		gongStruct, ok :=
			(*gong_models.GetGongstructInstancesMap[gong_models.GongStruct]())[gongStructName]

		if !ok {
			log.Println("updateGongObjectsNodes: Unknown gong struct related to identifier:", gongStructShape.Identifier)
			continue
		}
		gongStructImplIF, ok := nodeCb.map_gongObject_gongObjectImpl[gongStruct]
		if !ok {
			log.Println("updateGongObjectsNodes: Unknown gong struct impl related to gong struct:", gongStruct.Name)
			continue
		}
		// set up the pointer to the shape
		gongStructImpl, ok := gongStructImplIF.(*GongStructImpl)

		gongStructImpl2 := GetNodeBackPointer(gongStruct)
		_ = gongStructImpl2

		gongStructImpl.HasDiagramElt = true

		gongStructShapeNode.IsChecked = true

		// disable checkbox of all children of the gongstruct
		for _, fieldOfClassshapeNode := range gongStructShapeNode.Children {
			fieldOfClassshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
		}

		for _, field := range gongStructShape.Fields {
			classshapeFieldNode, ok := nodeCb.map_Identifier_Node[field.Identifier]

			if !ok {
				log.Fatalln(field.Identifier, "unknown node")
			}

			classshapeFieldNode.IsChecked = true
			classshapeFieldNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		}
		for _, link := range gongStructShape.Links {
			linkNode, ok := nodeCb.map_Identifier_Node[link.Identifier]

			if !ok {
				log.Fatalln(link.Identifier, "unknown node")
			}

			linkNode.IsChecked = true
			linkNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode
		}
	}

	// disable checkbox field node that reference a gongstruct whose classshape is
	// not present in the diagram
	// first, construct map of all gongstructs present in the diagram
	set_of_classshape_names := make(map[string]bool)
	for _, classshape := range classdiagram.GongStructShapes {
		set_of_classshape_names[gongdoc_models.IdentifierToShapename(classshape.Identifier)] = true
	}

	// then iterate over all fields of all gongstructs node
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
		classshapeNode := nodeCb.map_Identifier_Node[gongdoc_models.ShapenameToIdentifier(gongStruct.Name)]

		for _, fieldOfClassshapeNode := range classshapeNode.Children {
			// then disable the checkbox
			if fieldOfClassshapeNode.Type == gongdoc_models.GONG_STRUCT_FIELD {

				fieldImpl := fieldOfClassshapeNode.Impl.(*FieldImpl)

				switch fieldWithRef := fieldImpl.field.(type) {
				case *gong_models.PointerToGongStructField:
					if _, ok := set_of_classshape_names[fieldWithRef.GongStruct.Name]; !ok {
						fieldOfClassshapeNode.IsCheckboxDisabled = true
					}
				case *gong_models.SliceOfPointerToGongStructField:
					if _, ok := set_of_classshape_names[fieldWithRef.GongStruct.Name]; !ok {
						fieldOfClassshapeNode.IsCheckboxDisabled = true
					}
				}
			}
		}
	}

	//
	// For notes
	//
	// 1. for each note of the model, enable the check button if the class diagram is in draw mode
	set_of_noteshape_names := make(map[string]bool)
	for _, noteshape := range classdiagram.NoteShapes {
		set_of_noteshape_names[gongdoc_models.IdentifierToShapename(noteshape.Identifier)] = true
	}
	for note := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {
		noteshapeNode := nodeCb.map_Identifier_Node[gongdoc_models.ShapenameToIdentifier(note.Name)]
		noteshapeNode.IsCheckboxDisabled = !classdiagram.IsInDrawMode

		for _, noteLink := range note.Links {

			noteLinkNode := nodeCb.map_Identifier_Node[gongdoc_models.ShapeAndFieldnameToFieldIdentifier(note.Name, noteLink.Name)]

			// disable check box of the note link if the note is not present
			if _, ok := set_of_noteshape_names[note.Name]; !ok {
				noteLinkNode.IsCheckboxDisabled = true
			}

			// disable check box of the note link if the target shape is not present
			if _, ok := set_of_classshape_names[noteLink.Name]; !ok {
				noteLinkNode.IsCheckboxDisabled = true
			}

		}
	}

	// 2. for each noteShape of the diagram, set the check button of the related node to true
	for _, noteshape := range classdiagram.NoteShapes {
		noteshapeNode, ok := nodeCb.map_Identifier_Node[noteshape.Identifier]

		if !ok {
			log.Println("Unknown identifier", noteshape.Identifier)
		}

		noteshapeNode.IsChecked = true

		// 2.1 for each link of the note in the diagram, set the check button to true
		for _, noteShapeLink := range noteshape.NoteShapeLinks {
			noteShapeLinkNode := nodeCb.map_Identifier_Node[noteShapeLink.Identifier]
			noteShapeLinkNode.IsChecked = true
		}
	}
}

// updateNodesStates updates the tree of symbols
// according to the selected diagram
//
// ## The algorithm is as follow
//
// ## For the diagram nodes
//
// For the identifiers nodes
func (nodeCb *NodeCB) updateNodesStates(stage *gongdoc_models.StageStruct) {

	nodeCb.treeOfGongObjects.UncheckAndDisable()

	nodeCb.updateDiagramsNodes(stage)

	// get the diagram
	classdiagram := nodeCb.diagramPackage.SelectedClassdiagram

	if classdiagram == nil {
		gongdoc_models.Stage.Commit()
		return
	}

	nodeCb.updateGongObjectsNodes(stage, classdiagram)

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("UpdateNodeStates, after  commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())

}
