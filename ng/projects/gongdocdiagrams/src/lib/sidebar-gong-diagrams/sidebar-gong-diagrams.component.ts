import { Component, OnInit, ViewChild } from '@angular/core';
import { CdkDragDrop } from '@angular/cdk/drag-drop';
import { ElementRef } from '@angular/core';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

// insertion point for per struct import code
import * as gong from 'gong'
import * as gongdoc from 'gongdoc'

import { ClassdiagramContextSubject, ClassdiagramContext } from '../diagram-displayed-gongstruct'
import { combineLatest, Observable, timer } from 'rxjs';
import { ClassshapeDB, FieldDB } from 'gongdoc';
import { NoDataRowOutlet } from '@angular/cdk/table';

/**
 * GongNode is the "data" node that is construed from the gongFrontRepo
 * 
 */
interface GongNode {
  name: string // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children?: GongNode[]
  type: gongdoc.GongdocNodeType
  structName: string
  id: number
  uniqueIdPerStack: number

  // specific field for gongdoc
  presentInDiagram?: boolean // if the corresponding graphic element is present in the diagram, this value is true
  gongBasicField?: gong.GongBasicFieldDB
  gongPointerToGongStructField?: gong.PointerToGongStructFieldDB
  gongSliceOfPointerToGongStructField?: gong.SliceOfPointerToGongStructFieldDB

  // if the node is for a link between two gong struct, both have to be present in the diagram
  // for canBeIncluded to be true
  canBeIncluded?: boolean

}

/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
export interface GongFlatNode {
  expandable: boolean
  name: string
  level: number
  type: gongdoc.GongdocNodeType
  structName: string
  id: number
  uniqueIdPerStack: number

  // specific field for gongdoc
  presentInDiagram?: boolean
  gongBasicField?: gong.GongBasicFieldDB
  gongPointerToGongStructField?: gong.PointerToGongStructFieldDB
  gongSliceOfPointerToGongStructField?: gong.SliceOfPointerToGongStructFieldDB
  canBeIncluded?: boolean
}

export interface DragDropPosition {
  x: number;
  y: number;
}

/**
 * SidebarGongDiagramsComponent is the component that displays all gongstructs of the gong model
 * that is modeled
 * 
 * each gong struct have:
 * - basic fields
 * - pointer to gong struct
 * - slice of pointer to gong struct
 * 
 * the component deals with actions on those elements
 * 
 * the sidebar component is bespoke rework of the default gong generated sidebar
 */
@Component({
  styleUrls: ['./sidebar-gong-diagrams.css'],
  selector: 'lib-sidebar-gong-diagams',
  templateUrl: './sidebar-gong-diagrams.html'
})
export class SidebarGongDiagramsComponent implements OnInit {

  /**
   * innerHTMLelement is the html elemnt of the diagram
   * it allows the Sidebar component to devise the drop position of the gongstruct
   */
  @ViewChild('innerHTMLelement') innerHTMLelement: ElementRef;

  // the classdiagram that is currently displayed
  currentClassdiagram: gongdoc.ClassdiagramDB

  /**
   * initial node expansion
   * 
   * by default, all nodes are collapsed. This timer expands root nodes
   */
  expandRootNodesTimer: Observable<number> = timer(1000)

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,

      // specific to gongdoc
      present: node.presentInDiagram,
      gongBasicField: node.gongBasicField,
      gongPointerToGongStructField: node.gongPointerToGongStructField,
      gongSliceOfPointerToGongStructField: node.gongSliceOfPointerToGongStructField,
      canBeIncluded: node.canBeIncluded,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => {
      return node.expandable
    }
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  gongFrontRepo: gong.FrontRepo
  gongdocFrontRepo: gongdoc.FrontRepo

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // provide the current display context
  classdiagramContext: ClassdiagramContext

  constructor(
    private gongFrontRepoService: gong.FrontRepoService,
    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private gongdocCommandService: gongdoc.GongdocCommandService,
  ) { }

  ngOnInit(): void {

    this.expandRootNodesTimer.subscribe(
      currTime => {
        // expand nodes that were exapanded before
        if (this.treeControl.dataNodes != undefined) {
          this.treeControl.dataNodes.forEach(
            node => {
              if (node.uniqueIdPerStack == 13) {
                this.treeControl.expand(node)
              }
            }
          )
        }
      }
    )

    this.gongdocFrontRepoService.pull().subscribe(
      gongdocFrontRepo => {
        this.gongdocFrontRepo = gongdocFrontRepo

        // listen to any new diagram draw in order to update the 
        // gong tree appaeance
        ClassdiagramContextSubject.subscribe(
          classdiagramContext => {
            this.classdiagramContext = classdiagramContext
            this.currentClassdiagram = this.gongdocFrontRepo.Classdiagrams.get(classdiagramContext.ClassdiagramID)
            this.refresh()
          }
        )
      }
    )
  }

  refresh(): void {
    this.gongFrontRepoService.pull().subscribe(frontRepo => {
      this.gongFrontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (this.treeControl.isExpanded(node)) {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = true
            } else {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = false
            }
          }
        )
      }

      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction

      /**
      * fill up the GongStruct part of the mat tree
      */
      let gongstructGongNodeStruct: GongNode = {
        name: "GongStruct",
        type: gongdoc.GongdocNodeType.ROOT_OF_GONG_STRUCTS,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "GongStruct",
        children: new Array<GongNode>(),

        // the root node is neither present not draggable
        presentInDiagram: false,
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(gongstructGongNodeStruct)

      // create the set of classshapes presents in the class diagram
      // important for knowing which shapes are already displayed are 
      let arrayOfDisplayedClassshape = new Map<string, ClassshapeDB>()
      let arrayOfDisplayedBasicField = new Map<string, FieldDB>()
      let arrayOfDisplayedLink = new Map<string, FieldDB>()

      this.currentClassdiagram.Classshapes?.forEach(
        classshape => {
          arrayOfDisplayedClassshape.set(classshape.Structname, classshape)
          classshape?.Fields?.forEach(
            field => {
              arrayOfDisplayedBasicField.set( classshape.Structname + "." + field.Fieldname, field)
              console.log("Adding " + classshape.Structname + "." + field.Fieldname)
            }
          )
          classshape?.Links?.forEach(
            link => {
              arrayOfDisplayedLink.set( classshape.Structname + "." + link.Fieldname, link)
            }
          )
        }
      )

      this.gongFrontRepo.GongStructs_array.forEach(
        gongstructDB => {

          let classshape = arrayOfDisplayedClassshape.get(gongstructDB.Name)

          let gongstructGongNodeInstance: GongNode = {
            name: gongstructDB.Name,
            type: gongdoc.GongdocNodeType.GONG_STRUCT,
            id: gongstructDB.ID,
            uniqueIdPerStack: gong.getGongStructUniqueID(gongstructDB.ID),
            structName: gongstructDB.Name,
            children: new Array<GongNode>(),

            // specific to gongdoc
            presentInDiagram: arrayOfDisplayedClassshape.has(gongstructDB.Name),

          }
          gongstructGongNodeStruct.children.push(gongstructGongNodeInstance)

          // insertion point for per field code 
          /**
          * let append a node for the slide of pointer GongBasicFields
          */
          let GongBasicFieldsGongNodeAssociation: GongNode = {
            name: "GongBasicFields",
            type: gongdoc.GongdocNodeType.ROOT_OF_BASIC_FIELDS,
            id: 0,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "GongStruct",
            children: new Array<GongNode>(),

          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(GongBasicFieldsGongNodeAssociation)

          gongstructDB.GongBasicFields?.forEach(gongbasicfieldDB => {

            let structIsPresent = arrayOfDisplayedClassshape.has(gongbasicfieldDB.GongStruct_GongBasicFields_reverse.Name)

            let presentInDiagram = arrayOfDisplayedBasicField.has(gongstructDB.Name + "." + gongbasicfieldDB.Name)

            let gongbasicfieldNode: GongNode = {
              name: gongbasicfieldDB.Name,
              type: gongdoc.GongdocNodeType.BASIC_FIELD,
              id: gongbasicfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getGongBasicFieldUniqueID(gongbasicfieldDB.ID),
              structName: gongstructDB.Name,
              gongBasicField: gongbasicfieldDB,
              children: new Array<GongNode>(),
              presentInDiagram: presentInDiagram,
              canBeIncluded: structIsPresent,
            }
            GongBasicFieldsGongNodeAssociation.children.push(gongbasicfieldNode)
          })

          /**
          * let append a node for the slide of pointer PointerToGongStructFields
          */
          let PointerToGongStructFieldsGongNodeAssociation: GongNode = {
            name: "PointerToGongStructFields",
            type: gongdoc.GongdocNodeType.ROOT_OF_POINTER_TO_STRUCT_FIELDS,
            id: 0,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: gongstructDB.Name,
            children: new Array<GongNode>(),

          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(PointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.PointerToGongStructFields?.forEach(pointerToGongstructFieldDB => {

            // console.log("field " + pointertogongstructfieldDB.Name)
            let sourceIsPresent = arrayOfDisplayedClassshape.has(pointerToGongstructFieldDB.GongStruct_PointerToGongStructFields_reverse.Name)
            // console.log("source " + pointertogongstructfieldDB.GongStruct_PointerToGongStructFields_reverse.Name + " is present " + sourceIsPresent)

            // compute wether the link can be included in the diagram
            let destinationIsPresent = arrayOfDisplayedClassshape.has(pointerToGongstructFieldDB.GongStruct.Name)
            // console.log("destination " + pointertogongstructfieldDB.GongStruct.Name + " is present " + destinationIsPresent)

            let canBeIncluded = sourceIsPresent && destinationIsPresent

            let pointertogongstructfieldNode: GongNode = {
              name: pointerToGongstructFieldDB.Name + " (" + pointerToGongstructFieldDB.GongStruct.Name + ")",
              type: gongdoc.GongdocNodeType.POINTER_TO_STRUCT,
              id: pointerToGongstructFieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getPointerToGongStructFieldUniqueID(pointerToGongstructFieldDB.ID),
              structName: gongstructDB.Name,
              children: new Array<GongNode>(),
              gongPointerToGongStructField: pointerToGongstructFieldDB,
              presentInDiagram: arrayOfDisplayedLink.has(gongstructDB.Name + "." + pointerToGongstructFieldDB.Name),
              canBeIncluded: canBeIncluded,
            }
            // console.log("can be included ? " + pointertogongstructfieldNode.name + " " +
            //   pointertogongstructfieldNode.canBeIncluded + " " +
            //   canBeIncluded)
            PointerToGongStructFieldsGongNodeAssociation.children.push(pointertogongstructfieldNode)
          })

          /**
          * let append a node for the slide of pointer SliceOfPointerToGongStructFields
          */
          let SliceOfPointerToGongStructFieldsGongNodeAssociation: GongNode = {
            name: "SliceOfPointerToGongStructFields",
            type: gongdoc.GongdocNodeType.ROOT_OF_SLICE_OF_POINTER_TO_GONG_STRUCT_FIELDS,
            id: 0,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: gongstructDB.Name,
            children: new Array<GongNode>(),

          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          gongstructGongNodeInstance.children.push(SliceOfPointerToGongStructFieldsGongNodeAssociation)

          gongstructDB.SliceOfPointerToGongStructFields?.forEach(sliceofpointertogongstructfieldDB => {

            let sourceIsPresent = arrayOfDisplayedClassshape.has(sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFields_reverse.Name)
            let destinationIsPresent = arrayOfDisplayedClassshape.has(sliceofpointertogongstructfieldDB.GongStruct.Name)
            let canBeIncluded = sourceIsPresent && destinationIsPresent
            let presentInDiagram = arrayOfDisplayedLink.has(gongstructDB.Name + "." + sliceofpointertogongstructfieldDB.Name)

            let sliceofpointertogongstructfieldNode: GongNode = {
              name: sliceofpointertogongstructfieldDB.Name + " (" + sliceofpointertogongstructfieldDB.GongStruct.Name + ")",
              type: gongdoc.GongdocNodeType.SLICE_OF_POINTER_TO_STRUCT,
              id: sliceofpointertogongstructfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * gong.getGongStructUniqueID(gongstructDB.ID)
                + 11 * gong.getSliceOfPointerToGongStructFieldUniqueID(sliceofpointertogongstructfieldDB.ID),
              structName: gongstructDB.Name,
              children: new Array<GongNode>(),
              gongSliceOfPointerToGongStructField: sliceofpointertogongstructfieldDB,
              presentInDiagram: presentInDiagram,
              canBeIncluded: canBeIncluded,
            }
            SliceOfPointerToGongStructFieldsGongNodeAssociation.children.push(sliceofpointertogongstructfieldNode)
          })
        })

      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (memoryOfExpandedNodes[node.uniqueIdPerStack] != undefined) {

              if (memoryOfExpandedNodes[node.uniqueIdPerStack]) {
                this.treeControl.expand(node)
              }
            }
          }
        )
      }


    });
  }

  /**
   * removeNodeFromDiagram is called from the html template
   * 
   * @param gongFlatNode 
   */
  removeNodeFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type

      switch (gongFlatNode.type) {
        case 'BASIC_FIELD':
          gongdocCommandSingloton.FieldName = gongFlatNode.name
      }

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  /**
   * removeBasicFieldFromDiagram is called from the html template
   * 
   * @param gongFlatNode 
   */
  removeBasicFieldFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
      gongdocCommandSingloton.FieldName = gongFlatNode.name

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  addBasicFieldToDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
      gongdocCommandSingloton.FieldName = gongFlatNode.name
      gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongBasicField?.BasicKindName

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  removePointerToStructFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
      gongdocCommandSingloton.FieldName = gongFlatNode.gongPointerToGongStructField.Name
      gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongPointerToGongStructField.GongStruct?.Name

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  addPointerToStructToDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
      gongdocCommandSingloton.FieldName = gongFlatNode.gongPointerToGongStructField?.Name
      gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongPointerToGongStructField?.GongStruct.Name

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  removeSliceOfPointerToStructFromDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_DELETE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
      gongdocCommandSingloton.FieldName = gongFlatNode.gongSliceOfPointerToGongStructField.Name
      gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongSliceOfPointerToGongStructField.GongStruct?.Name

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  addSliceOfPointerToStructToDiagram(gongFlatNode: GongFlatNode) {

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {
      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type
      gongdocCommandSingloton.FieldName = gongFlatNode.gongSliceOfPointerToGongStructField?.Name
      gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongSliceOfPointerToGongStructField?.GongStruct.Name

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  /**
   * dropped is called from the html template
   * 
   * @param event 
   * @param gongFlatNode 
   * @returns 
   */
  dropped(event: CdkDragDrop<ElementRef<HTMLElement>, any>, gongFlatNode: GongFlatNode): void {

    /**
     * dragDistance is the distance betweeen the point when the element is taken
     * and the point where the element is dropped
     */
    const dragVector = event.distance;

    /**
     * originDraggedElement is the position with the absolute coordinates of the starting point
     */
    const originDraggedElement = this.getPositionFromNativeElement(
      event.item.element.nativeElement
    );

    /**
     * droppedPosition = starting point + displacement
     */
    const droppedPosition = this.getDroppingPosition(
      dragVector,
      originDraggedElement
    );
    const paper = document.getElementById('jointjs-holder');
    if (paper == null) {
      console.warn(
        "DragDropService - Diagram paper must be definied in the HTML under the id 'jointjs-holder' (#jointjs-holder)"
      );
      return;
    }

    /**
     * the innerHTMLelement provide an offset in x
     */
    const sourceElement = this.getOffsetWidthHeightFromNativeElement(this.innerHTMLelement.nativeElement)

    /**
     * 64
     */
    const originPaper = this.getPositionFromNativeElement(paper);

    /**
     * dropOnPaperOffset is the computed position the jointjs paper
     */
    const dropOnPaperOffset = {
      x: droppedPosition.x - originPaper.x,
      y: droppedPosition.y - originPaper.y,
    };

    // get the GongdocCommandSingloton
    let gongdocCommandSingloton = this.gongdocFrontRepo.GongdocCommands.get(1)
    if (gongdocCommandSingloton != undefined) {

      gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.DIAGRAM_ELEMENT_CREATE
      gongdocCommandSingloton.StructName = gongFlatNode.structName
      gongdocCommandSingloton.DiagramName = this.currentClassdiagram.Name
      gongdocCommandSingloton.Date = Date.now().toString()
      gongdocCommandSingloton.PositionX = dropOnPaperOffset.x
      gongdocCommandSingloton.PositionY = dropOnPaperOffset.y
      gongdocCommandSingloton.GongdocNodeType = gongFlatNode.type

      switch (gongFlatNode.type) {
        case 'BASIC_FIELD':
          gongdocCommandSingloton.FieldName = gongFlatNode.name
          gongdocCommandSingloton.FieldTypeName = gongFlatNode.gongBasicField?.BasicKindName
      }

      this.gongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
        GongdocCommand => {
          console.log("GongdocCommand updated")
        }
      )
    }
  }

  getPositionFromNativeElement(el: HTMLElement): DragDropPosition {
    return {
      x: el.offsetLeft,
      y: el.offsetTop,
    };
  }

  getOffsetWidthHeightFromNativeElement(el: HTMLElement): DragDropPosition {
    return {
      x: el.offsetWidth,
      y: el.offsetHeight,
    };
  }

  getDroppingPosition(
    dragDropDistance: DragDropPosition,
    dragOrigin: DragDropPosition
  ): DragDropPosition {
    return {
      x: dragOrigin.x + dragDropDistance.x,
      y: dragOrigin.y + dragDropDistance.y,
    };

  }
}
