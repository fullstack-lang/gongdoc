import { AfterViewInit, Component, ElementRef, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import { Coordinate, RectangleEventService } from '../rectangle-event.service';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';

import * as gongsvg from 'gongsvg'
import { ShapeMouseEvent } from '../shape.mouse.event';
import { createPoint } from '../link/draw.segments';
import { MouseEventService } from '../mouse-event.service';

@Component({
  selector: 'lib-svg',
  templateUrl: './svg.component.html',
  styleUrls: ['./svg.component.css']
})
export class SvgComponent implements OnInit, OnDestroy, AfterViewInit {

  @Input() GONG__StackPath: string = ""

  public gongsvgFrontRepo?: gongsvg.FrontRepo


  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  svg = new gongsvg.SVGDB
  linkStartRectangleID: number = 0

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];


  // if true, the end user is shiftKey + mouse down from one rectangle
  // to another
  linkDrawing: boolean = false
  startX = 0;
  startY = 0;
  endX = 0;
  endY = 0;

  selectionRectDrawing: boolean = false
  rectX = 100;
  rectY = 100;
  width = 300;
  height = 40;

  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    private gongsvgPushFromFrontNbService: gongsvg.PushFromFrontNbService,
    private svgService: gongsvg.SVGService,
    private rectangleEventService: RectangleEventService,
    private svgEventService: SvgEventService,
    private mouseEventService: MouseEventService,
  ) {

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseDownEvent$.subscribe(
        ({ rectangleID: rectangleID, MousePosRelativeSVG: coordinate }) => {
          // console.log('SvgComponent, Mouse down event occurred on rectangle ', rectangleID, " at ", coordinate)
          this.linkStartRectangleID = rectangleID

          let rect = this.gongsvgFrontRepo?.Rects.get(rectangleID)

          if (rect == undefined) {
            return
          }

          this.linkDrawing = true
          this.startX = coordinate[0]
          this.startY = coordinate[1]
        })
    );

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseDragEvent$.subscribe((coordinate: Coordinate) => {

        this.endX = coordinate[0]
        this.endY = coordinate[1]
        // console.log('SvgComponent, Received Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);
      })
    )

    this.subscriptions.push(
      rectangleEventService.mouseRectAltKeyMouseUpEvent$.subscribe((rectangleID: number) => {
        // console.log('SvgComponent, Mouse up event occurred on rectangle ', rectangleID);
        this.linkDrawing = false

        this.onEndOfLinkDrawing(this.linkStartRectangleID, rectangleID)
      })
    )

    this.subscriptions.push(
      svgEventService.multiShapeSelectDragEvent$.subscribe(
        (coordinate: Coordinate) => {

          let actualX = coordinate[0]
          let actualY = coordinate[1]

          this.rectX = Math.min(this.startX, actualX);
          this.rectY = Math.min(this.startY, actualY);
          this.width = Math.abs(actualX - this.startX);
          this.height = Math.abs(actualY - this.startY);
        })
    );

    this.subscriptions.push(
      svgEventService.mouseShiftKeyMouseUpEvent$.subscribe(
        (coordinate: Coordinate) => {

          let actualX = coordinate[0]
          let actualY = coordinate[1]

          this.selectionRectDrawing = false
          this.endX = actualX - this.pageX
          this.endY = actualY - this.pageY

          let selectAreaConfig: SelectAreaConfig = new SelectAreaConfig()

          if (this.endX > this.startX) {
            selectAreaConfig.SweepDirection = SweepDirection.LEFT_TO_RIGHT
            selectAreaConfig.TopLeft = [this.startX, this.startY]
            selectAreaConfig.BottomRigth = [this.endX, this.endY]
          } else {
            selectAreaConfig.SweepDirection = SweepDirection.RIGHT_TO_LEFT
            selectAreaConfig.TopLeft = [this.endX, this.endY]
            selectAreaConfig.BottomRigth = [this.startX, this.startY]
          }

          this.svgEventService.emitMultiShapeSelectEnd(selectAreaConfig)
        }
      )
    )
  }



  ngOnInit(): void {

    // console.log("Svg component->ngOnInit : GONG__StackPath, " + this.GONG__StackPath)

    // see above for the explanation
    this.gongsvgNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          // console.log("last commit nb " + this.lastCommiNbFromBagetCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
          this.refresh()
          this.lastCommitNbFromBack = commiNbFromBagetCommitNbFromBack
        }
      }
    )

    this.gongsvgPushFromFrontNbService.getPushNbFromFront(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          // console.log("last commit nb " + this.lastCommiNbFromBagetCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
          this.refresh()
          this.lastCommitNbFromBack = commiNbFromBagetCommitNbFromBack
        }
      }
    )
  }

  refresh(): void {

    this.gongsvgFrontRepoService.pull(this.GONG__StackPath).subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo

        if (this.gongsvgFrontRepo.SVGs_array.length == 1) {
          this.svg = this.gongsvgFrontRepo.SVGs_array[0]
        } else {
          return
        }

        if (this.svg.Layers == undefined) {
          return
        }

        this.svg.Layers.sort((t1, t2) => {
          let t1_revPointerID_Index = t1.SVG_LayersDBID_Index
          let t2_revPointerID_Index = t2.SVG_LayersDBID_Index

          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }

    )
  }

  ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
  }

  onEndOfLinkDrawing(startRectangleID: number, endRectangleID: number) {

    let svgArray = this.gongsvgFrontRepo?.SVGs_array
    // update the svg
    if (svgArray?.length == 1) {
      this.svg = svgArray[0]
    } else {
      return
    }

    if (this.svg.Layers == undefined) {
      return
    }

    if (this.svg.DrawingState != gongsvg.DrawingState.NOT_DRAWING_LINE) {
      // console.log("problem with svg, length ", this.svg.DrawingState, " is not ", gongsvg.DrawingState.NOT_DRAWING_LINE)
    }

    this.svg.DrawingState = gongsvg.DrawingState.DRAWING_LINE

    this.svg.StartRectID.Valid = true
    this.svg.StartRectID.Int64 = startRectangleID

    this.svg.EndRectID.Valid = true
    this.svg.EndRectID.Int64 = endRectangleID

    this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe(
      () => {
        // back to normal state
        this.svg.DrawingState = gongsvg.DrawingState.NOT_DRAWING_LINE
        this.svgService.updateSVG(this.svg, this.GONG__StackPath).subscribe()
      }
    )
  }

  onClick(event: MouseEvent) {
    // an event is emitted for all rects to go on a unselect mode
    if (!event.altKey && !event.shiftKey) {
      console.log("SVG : on click()")

      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: 0,
        ShapeType: "",
        Point: createPoint(x, y),
      }

      this.rectangleEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  mousedown(event: MouseEvent): void {
    if (event.shiftKey) {
      this.selectionRectDrawing = true
      this.startX = event.clientX - this.pageX
      this.startY = event.clientY - this.pageY
    }
  }

  mousemove(event: MouseEvent): void {
    const x = event.clientX - this.pageX
    const y = event.clientY - this.pageY

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      // console.log('SvgComponent, ALT Mouse drag event occurred', this.linkDrawing, this.startX, this.startY, this.endX, this.endY);

      this.rectangleEventService.emitRectAltKeyMouseDragEvent([x, y])
      return
    }
    if (event.shiftKey) {

      this.svgEventService.emitMultiShapeSelectDrag([x, y])
      // console.log('SvgComponent, SHIFT Mouse drag event occurred', this.selectionRectDrawing, this.rectX, this.rectY, this.width, this.height);
    }

    if (!event.shiftKey && !event.altKey) {
      // in case of dragging something, when the mouse moves fast, it can reach the SVG background
      // in this case, one forward the mouse event on the event bus
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: 0,
        ShapeType: "",
        Point: createPoint(x, y),
      }
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }


  onmouseup(event: MouseEvent): void {

    const x = event.clientX - this.pageX
    const y = event.clientY - this.pageY

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent([event.clientX, event.clientY])
    }

    if (!event.shiftKey && !event.altKey) {

      // in case of dragging something, when the mouse moves fast, it can reach the SVG background
      // in this case, one forward the mouse event on the event bus
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: 0,
        ShapeType: "",
        Point: createPoint(x, y),
      }
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)
    }
  }

  pageX: number = 0
  pageY: number = 0
  @ViewChild('drawingArea') drawingArea: ElementRef<HTMLDivElement> | undefined

  ngAfterViewInit() {
    const offset = this.getDivOffset(this.drawingArea!.nativeElement);
    this.pageX = offset.offsetX
    this.pageY = offset.offsetY
  }

  getDivOffset(div: HTMLDivElement): { offsetX: number; offsetY: number } {
    const rect = div.getBoundingClientRect();
    const offsetX = rect.left + window.pageXOffset;
    const offsetY = rect.top + window.pageYOffset;
    return { offsetX, offsetY };
  }

}
