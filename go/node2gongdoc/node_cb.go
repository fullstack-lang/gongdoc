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
func (nodesCb *NodeCB) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	stagedNode.Impl.OnAfterUpdate(stage, stagedNode, frontNode)

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

	nodeCb.diagramPackageNode.Children =
		append(nodeCb.diagramPackageNode.Children, node)

	// set up the back pointer from the shape to the node
	classdiagramImpl := new(ClassdiagramImpl)
	classdiagramImpl.node = node
	classdiagramImpl.classdiagram = classdiagram
	classdiagramImpl.nodeCb = nodeCb
	node.Impl = classdiagramImpl

	updateNodesStates(stage, nodeCb)

}

// OnAfterDelete is called after a node is deleted
// notice that the fontNode only have its basic fields updated
// its pointers are not ok
func (nodesCb *NodeCB) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	switch impl := stagedNode.Impl.(type) {
	case *ClassdiagramImpl:
		impl.OnAfterDelete(stage, stagedNode, frontNode)
	}
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
