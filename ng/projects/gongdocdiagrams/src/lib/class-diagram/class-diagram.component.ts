import { Component, OnDestroy, OnInit } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';
import * as joint from 'jointjs';

import { ActivatedRoute, Router } from '@angular/router';

import * as gongdoc from 'gongdoc'
import * as gong from 'gong'

import { newUmlClassShape } from './newUmlClassShape'
import { ClassdiagramContextSubject, ClassdiagramContext } from '../diagram-displayed-gongstruct'

@Component({
  selector: 'lib-class-diagram',
  templateUrl: './class-diagram.component.html',
  styleUrls: ['./class-diagram.component.css']
})
export class ClassDiagramComponent implements OnInit, OnDestroy {


  /**
   * the class diagram component is refreshed both by direct input when the user moves vertices or positions
   * otherwise, modification are gotten from the back repo 
   * 
   * the checkCommitNbTimer polls the commit number of the back repo
   * if the commit number has increased, it pulls the front repo and redraw the diagram
   */
  checkGongdocCommitNbTimer: Observable<number> = timer(500, 500);
  lastCommitNb = -1
  lastDiagramId = 0
  currTime: number

  /**
   * jointjs stuff
   */
  namespace = joint.shapes;
  private paper: joint.dia.Paper;
  private graph: joint.dia.Graph;

  // the gong diagram of interest ot be drawn
  public classdiagram: gongdoc.ClassdiagramDB;

  /**
   * gongdoc manages 2 stacks: gongdoc and gong
   * 
   * gongdoc is for CRUDing the diagram elements (shapes, links, ...)
   * 
   * gong is for getting the elements of the go package (it has to be of an isntance of a gong model)
   * that are being modeled
   * 
   * instances in gongdoc refers to the gong model via the names of the elements.
   * a classshape instance modeling a "foo" gong struct will have a the value "foo" for the Structname field 
   */
  gongdocFrontRepo: gongdoc.FrontRepo
  gongFrontRepo: gong.FrontRepo

  // map of Classhapes according to the joint.shapes.uml.Class
  // it is used to save the position of the classhape in the diagram (which only know the ids)
  // into the Classshape object (via its Position field)
  public Map_CellId_ClassshapeDB = new Map<string, gongdoc.ClassshapeDB>();

  // map for storing which gong struct have a classshape
  // it is important for drawing links between shapes
  public Map_GongStructName_JointjsUMLClassShape = new Map<string, joint.shapes.uml.Class>();

  // important for storing the relation from the jointjs link to the 
  // gongdoc Link. When the user saves the position of the vertice, this enables
  // the saving it into the Link object
  public Map_CellId_LinkDB = new Map<string, gongdoc.LinkDB>();

  constructor(
    private route: ActivatedRoute,

    private PositionService: gongdoc.PositionService,
    private VerticeService: gongdoc.VerticeService,

    private gongdocFrontRepoService: gongdoc.FrontRepoService,
    private GongdocCommandService: gongdoc.GongdocCommandService,
    private gongdocCommitNbService: gongdoc.CommitNbService,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    // this.router.routeReuseStrategy.shouldReuseRoute = function () {
    //   return false;
    // };
  }

  // Since this component is not reused when a new diagram is selected, there can be many
  // instances of the diagram and each instance will stay alive. For instance,
  // the instance will be in the control flow if an observable the component subscribes to emits an event.
  // Therefore, it is mandatory to manage subscriptions in order to unscribe them on the ngOnDestroy hook
  subscriptionToDragAndDropEvent: Subscription
  subscriptionToRemoveFromDiagramEvent: Subscription

  // neccessary to unsubscribe
  ngOnDestroy() {
    console.log("on destroy")
  }

  ngOnInit(): void {

    // check loop for refresh from the back repo
    this.checkGongdocCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        const id = +this.route.snapshot.paramMap.get('id');

        this.gongdocCommitNbService.getCommitNb().subscribe(
          commitNb => {

            // condition for refresh
            if (this.lastCommitNb < commitNb || this.lastDiagramId != id) {

              console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
              console.log("last diagram id " + this.lastDiagramId + " new: " + id)
              this.pullGongdocAndDrawDiagram()
              this.lastCommitNb = commitNb
              this.lastDiagramId = id
            }
          }
        )
      }
    )
  }

  //
  // make a jointjs umlclass from a gong Classshape object
  //
  addClassshapeToGraph(classshape: gongdoc.ClassshapeDB) {
    //
    // Fetch fields 
    //
    const umlClassShape = newUmlClassShape(classshape)

    // structRectangle.attributes = ['firstName: String']
    umlClassShape.addTo(this.graph);

    // horrible hack because the TS compiles assets that umlclasshape.id is not a string but a 
    // an attribute of type joint.dia.Dimension
    var id: any;
    id = umlClassShape.id;
    var idstring: string
    idstring = id;
    this.Map_CellId_ClassshapeDB.set(idstring, classshape)
    this.Map_GongStructName_JointjsUMLClassShape.set(classshape.Structname, umlClassShape)
  }

  //
  // turn gong instances into a jointjs diagram
  //
  drawClassdiagram(): void {
    console.log("draw diagram")

    const namespace = joint.shapes;

    //
    // one hairy stuff is the computation of the drawing size
    //
    // this is a work in progress
    //
    let diagramWidth = 300
    if (this.classdiagram != undefined) {
      if (this.classdiagram.Classshapes != undefined) {
        diagramWidth = (this.classdiagram.Classshapes.length + 1) * 300
      }
    }

    //
    // a jointjs diagram is a Graph instance with a Paper instance
    //
    this.graph = new joint.dia.Graph(
      {},
      { cellNamespace: this.namespace } // critical piece of code. 
    );
    this.paper = new joint.dia.Paper(
      {
        el: document.getElementById('jointjs-holder'),

        model: this.graph,
        // Estimate the total width of the diagram >> 1 class width: (200px * 1.2 (css transform scale)) + margin: 30px * 2 (left-right)
        // The shapes are divided on 2 rows
        // MANAGE_ODD_NUMBERSHAPES = numbershapes % 2 ? 300 / 2 : 0
        // => Total width: numbershapes / 2 * 300 + MANAGE_ODD_NUMBERSHAPES
        width: diagramWidth,
        height: 1000,
        gridSize: 10,
        drawGrid: false,
        cellViewNamespace: namespace
      }
    );

    // draw class shapes from the gong classshapes
    if (this.classdiagram?.Classshapes != undefined) {
      for (let classshape of this.classdiagram.Classshapes) {
        this.addClassshapeToGraph(classshape)
      }

      // draw links of the diagram shapes
      for (let classshape of this.classdiagram.Classshapes) {
        if (classshape.Links != undefined) {
          for (let linkDB of classshape.Links) {

            // does from & to shapes exists?
            //
            // a gong st
            var fromShape = this.Map_GongStructName_JointjsUMLClassShape.get(linkDB.Structname)
            var toShape = this.Map_GongStructName_JointjsUMLClassShape.get(linkDB.Fieldtypename)

            var strockWidth = 2
            let LinkEndlabel = linkDB.Fieldname
            let distanceEndLabel = 0.75
            let LinkMuliplicity = linkDB.Multiplicity
            let distanceMultiplicity = 0.95

            if (toShape == undefined) {
              // the destination shape is not in the diagram
              continue;
            }

            let xFrom = fromShape.get('position').x
            let yFrom = fromShape.get('position').y
            let xTo = toShape.get('position').x
            let yTo = toShape.get('position').y
            var vertices = [{ x: (xFrom + yTo) / 2, y: (yFrom + yTo) / 2 }]

            if (linkDB.Middlevertice != undefined) {
              vertices = [{ x: linkDB.Middlevertice.X, y: linkDB.Middlevertice.Y }]
            }

            if (fromShape != undefined && toShape != undefined) {
              var link = new joint.shapes.standard.Link({
                source: fromShape,
                target: toShape,
                vertices: vertices,
                attrs: {
                  line: {
                    stroke: '#3c4260',
                    strokeWidth: strockWidth,
                    vertexMarker: {
                      'type': 'circle',
                      'r': 3,
                      'stroke-width': 2,
                      'fill': 'white'
                    },
                    targetMarker: { // no arrow at the end
                      'type': 'path',
                      'd': 'M 10 -5 0 0 10 5 z'
                    },
                  },
                },
                labels: [
                  {
                    attrs: { text: { text: LinkEndlabel } },
                    position: {
                      offset: 15,
                      distance: distanceEndLabel
                    }
                  },
                  {
                    attrs: { text: { text: LinkMuliplicity } },
                    position: {
                      offset: 15,
                      distance: distanceMultiplicity
                    }
                  }
                ],
              })
              link.addTo(this.graph);

              // later, we need to save the diagram
              // 
              // algo is 
              // - for each cells of the diagram, 
              //      get its id & position
              //      find the original LinkDB and updates its position
              //
              // Because each cell has an unique id
              // we create a map of cell id to LinkDB in order 
              // horrible hack because the TS compiles assets that umlclasshape.id is not a string but a 
              // an attribute of type joint.dia.Dimension
              var id: any;
              id = link.id;
              var idstring: string
              idstring = id;
              this.Map_CellId_LinkDB.set(idstring, linkDB)
            }
          }
        }
      }
    }

    // allow some observers to know what are the displayed structs
    if (this.classdiagram) {
      let classdiagramContext = new ClassdiagramContext()
      classdiagramContext.ClassdiagramID = this.classdiagram.ID
      ClassdiagramContextSubject.next(classdiagramContext)
    }
  }

  /**
   * saving the classdiagram
   * 
   * the challenge is to update the positions of classshapes and vertices
   */
  saveClassdiagram(): void {
    console.log("save diagram")

    // parse shapes positions
    var cells = this.graph.getCells()
    console.log(cells.length)

    cells.forEach(
      cell => {
        // ugly hack because cell.id is considered a Dimension by the ts compiler
        // vive golang
        var classshapeDB: gongdoc.ClassshapeDB;
        var cellId: any
        cellId = cell.id;

        // update position of classshapes
        if (this.Map_CellId_ClassshapeDB.get(cellId) != undefined) {

          // retrieve the shape.
          classshapeDB = this.Map_CellId_ClassshapeDB.get(cellId)

          if (classshapeDB.Position == undefined) {
            console.log("link position undefined")
          }

          classshapeDB.Position.X = cell.attributes.position.x
          classshapeDB.Position.Y = cell.attributes.position.y

          // update position to DB
          this.PositionService.updatePosition(classshapeDB.Position).subscribe(
            position => {
              console.log("position updated")
            }
          )
        }

        // update positions of links
        if (this.Map_CellId_LinkDB.has(cellId)) {

          // retrieve the shape.
          var linkDB = this.Map_CellId_LinkDB.get(cellId)

          if (linkDB.Middlevertice == undefined) {
            console.log("link middle vertice undefined")
          }

          // fetch corresponding position and update
          linkDB.Middlevertice.X = cell.attributes.vertices[0].x
          linkDB.Middlevertice.Y = cell.attributes.vertices[0].y

          // update position to DB
          var verticeDB = linkDB.Middlevertice
          this.VerticeService.updateVertice(verticeDB).subscribe(
            position => {
              console.log("vertice updated")
            }
          )
        }
      }
    )

    // send a marshalling command to the backend via GongdocCommandSingloton
    let gongdocCommandSingloton: gongdoc.GongdocCommandDB
    this.gongdocFrontRepo.GongdocCommands.forEach(
      gongdocCommand => {
        gongdocCommandSingloton = gongdocCommand
      }
    )
    gongdocCommandSingloton.Command = gongdoc.GongdocCommandType.MARSHALL_DIAGRAM
    gongdocCommandSingloton.DiagramName = this.classdiagram.Name
    gongdocCommandSingloton.Date = Date.now().toString()

    this.GongdocCommandService.updateGongdocCommand(gongdocCommandSingloton).subscribe(
      GongdocCommand => {
        console.log("GongdocCommand updated")
      }
    )
  }

  pullGongdocAndDrawDiagram() {
    this.gongdocFrontRepoService.pull().subscribe(
      frontRepo => {
        this.gongdocFrontRepo = frontRepo
        console.log("gongdoc front repo pull returned")

        const id = +this.route.snapshot.paramMap.get('id');
        this.classdiagram = frontRepo.Classdiagrams.get(id)

        this.drawClassdiagram();
      }
    )
  }
}

