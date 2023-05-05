import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import { ElementRef, Renderer2, AfterViewInit } from '@angular/core';


import { RectangleEventService } from '../rectangle-event.service';
import { SelectAreaConfig, SvgEventService, SweepDirection } from '../svg-event.service';

import * as gongsvg from 'gongsvg'
import { Subscription } from 'rxjs';
import { ShapeMouseEvent } from '../shape.mouse.event';
import { createPoint } from '../link/draw.segments';
import { MouseEventService } from '../mouse-event.service';

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit, OnDestroy, AfterViewInit {

  @Input() Rect: gongsvg.RectDB = new gongsvg.RectDB()
  @Input() GONG__StackPath: string = ""

  // In your component
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  anchorDragging: boolean = false;
  activeAnchor: 'left' | 'right' | 'top' | 'bottom' | null = null;

  rectDragging: boolean = false;

  // RectAtMouseDown is the clone of the Rect when mouse down
  private RectAtMouseDown: gongsvg.RectDB | undefined

  // to compute wether it was a select / dragging event
  distanceMoved = 0
  private mousePosRelativeToSvgAtMouseDown: { x: number; y: number } = { x: 0, y: 0 }
  private PointAtMouseDown: gongsvg.PointDB | undefined
  private dragThreshold = 5;

  //
  // for events management
  //
  private subscriptions: Subscription[] = [];

  constructor(
    private rectService: gongsvg.RectService,
    private rectangleEventService: RectangleEventService,
    private mouseEventService: MouseEventService,
    private svgEventService: SvgEventService,
    private elementRef: ElementRef) {

    this.subscriptions.push(
      rectangleEventService.mouseMouseDownEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          this.RectAtMouseDown = structuredClone(this.Rect)
          this.PointAtMouseDown = structuredClone(shapeMouseEvent.Point)

          this.mousePosRelativeToSvgAtMouseDown = {
            x: shapeMouseEvent.Point.X,
            y: shapeMouseEvent.Point.Y
          }
          this.distanceMoved = 0
        })
    )

    this.subscriptions.push(
      rectangleEventService.mouseMouseMoveEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          const deltaX = shapeMouseEvent.Point.X - this.mousePosRelativeToSvgAtMouseDown.x
          const deltaY = shapeMouseEvent.Point.Y - this.mousePosRelativeToSvgAtMouseDown.y
          this.distanceMoved = Math.sqrt(deltaX * deltaX + deltaY * deltaY)

          if (this.rectDragging || this.Rect.IsSelected) {
            if (this.Rect?.CanMoveHorizontaly) {
              this.Rect.X = this.RectAtMouseDown!.X +
                (shapeMouseEvent.Point.X - this.mousePosRelativeToSvgAtMouseDown.x)
            }
            if (this.Rect?.CanMoveVerticaly) {
              this.Rect.Y = this.RectAtMouseDown!.Y +
                (shapeMouseEvent.Point.Y - this.mousePosRelativeToSvgAtMouseDown.y)
            }
          }
        })
    )

    this.subscriptions.push(
      rectangleEventService.mouseMouseUpEvent$.subscribe(
        (shapeMouseEvent: ShapeMouseEvent) => {

          if (shapeMouseEvent.ShapeID != 0 && this.distanceMoved > this.dragThreshold) {
            this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
          } else {
            if (this.Rect?.IsSelectable && shapeMouseEvent.ShapeID == this.Rect.ID) {
              console.log("rect, onRectClick() toggle selected: ", this.Rect?.Name)
              this.Rect.IsSelected = !this.Rect.IsSelected
              this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
            }

            // mouseup emited from the backgrund let to unselect selected shapes
            if (this.Rect.IsSelectable && this.Rect.IsSelected && shapeMouseEvent.ShapeID == 0) {
              this.Rect.IsSelected = false
              this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
            }
          }
          // console.log('Rect ', this.Rect.Name, 'Mouse down event occurred on rectangle ', rectangleID);
        })
    );

    this.subscriptions.push(
      svgEventService.multiShapeSelectEndEvent$.subscribe(
        (selectAreaConfig: SelectAreaConfig) => {

          if (this.Rect.IsSelected) {
            return
          }

          switch (selectAreaConfig.SweepDirection) {
            case SweepDirection.LEFT_TO_RIGHT:

              // rectangle has to be in boxed in the rect
              if (
                this.Rect.X > selectAreaConfig.TopLeft[0] &&
                this.Rect.X + this.Rect.Width < selectAreaConfig.BottomRigth[0] &&
                this.Rect.Y > selectAreaConfig.TopLeft[1] &&
                this.Rect.Y + this.Rect.Height < selectAreaConfig.BottomRigth[1]
              ) {
                this.Rect.IsSelected = true
                this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
              }
              break
            case SweepDirection.RIGHT_TO_LEFT:

              // rectangle has to be partialy boxed in the rect
              if (
                this.Rect.X < selectAreaConfig.BottomRigth[0] &&
                this.Rect.X + this.Rect.Width > selectAreaConfig.TopLeft[0] &&
                this.Rect.Y < selectAreaConfig.BottomRigth[1] &&
                this.Rect.Y + this.Rect.Height > selectAreaConfig.TopLeft[1]
              ) {
                this.Rect.IsSelected = true
                this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
              }
              break
          }
        })
    );
  }

  ngOnInit(): void {
  }

  onSVGClick(event: MouseEvent) {
    console.log("rect, onSVGClick(): ", this.Rect?.Name)
  }

  onRectClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (this.rectDragging) {
      this.rectDragging = false;
      return
    }

    if (!event.altKey && !event.shiftKey) {


    }
  }

  rectMouseDown(event: MouseEvent): void {

    if (!event.altKey) {
      event.preventDefault();
      event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

      this.rectDragging = true;

      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect!.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: createPoint(x, y),
      }
      this.rectangleEventService.emitMouseDownEvent(shapeMouseEvent)
    } else {
      console.log("startRectDrag + shiftKey : ", this.Rect?.Name)

      const actualX = event.clientX - this.pageX
      const actualY = event.clientY - this.pageY

      this.rectangleEventService.emitRectAltKeyMouseDownEvent(
        this.Rect.ID,
        [actualX, actualY]);
    }
  }

  rectMouseMove(event: MouseEvent): void {

    let x = event.clientX - this.pageX
    let y = event.clientY - this.pageY

    // console.log("RectComponent DragRect : ", deltaX, deltaY, distanceMoved)

    // we want this event to bubble to the SVG element
    if (event.altKey) {
      console.log('RectComponent, Alt Mouse drag event occurred on rectangle ', this.Rect.Name, event.clientX, event.clientY);
      this.rectangleEventService.emitRectAltKeyMouseDragEvent(
        [x, y])
      return
    }

    if (event.shiftKey) {
      this.svgEventService.emitMultiShapeSelectDrag([x, y])
      return
    }

    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!event.altKey && !event.shiftKey) {
      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect!.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: createPoint(x, y),
      }
      this.rectangleEventService.emitMouseMoveEvent(shapeMouseEvent)
      this.mouseEventService.emitMouseMoveEvent(shapeMouseEvent)
    }
  }

  rectMouseUp(event: MouseEvent): void {

    if (!event.altKey && !event.shiftKey) {
      let x = event.clientX - this.pageX
      let y = event.clientY - this.pageY

      let shapeMouseEvent: ShapeMouseEvent = {
        ShapeID: this.Rect!.ID,
        ShapeType: gongsvg.RectDB.GONGSTRUCT_NAME,
        Point: createPoint(x, y),
      }

      this.rectangleEventService.emitMouseUpEvent(shapeMouseEvent)
      this.mouseEventService.emitMouseUpEvent(shapeMouseEvent)

      this.rectDragging = false
    }

    if (event.altKey) {
      console.log("endRectDrag + shiftKey rect : ", this.Rect?.Name)

      this.rectangleEventService.emitMouseRectAltKeyMouseUpEvent(this.Rect.ID);
    }

    if (event.shiftKey) {
      this.svgEventService.emitMouseShiftKeyMouseUpEvent([event.clientX, event.clientY])
    }
  }

  ngOnDestroy() {
    this.subscriptions.forEach((subscription) => subscription.unsubscribe());
  }

  // 
  pageX: number = 0
  pageY: number = 0
  getSvgTopLeftCoordinates() {
    const svgElement = this.elementRef.nativeElement.querySelector('svg');
    const svgRect = svgElement.getBoundingClientRect();
    this.pageX = svgRect.left + window.pageXOffset;
    this.pageY = svgRect.top + window.pageYOffset;

    // console.log('SVG Top-Left Corner:', this.pageX, this.pageY);
  }

  ngAfterViewInit() {
    this.getSvgTopLeftCoordinates();
  }


  anchorMouseDown(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.anchorDragging = true;
    this.activeAnchor = anchor;
  }

  anchorMouseMove(event: MouseEvent, anchor: 'left' | 'right' | 'top' | 'bottom'): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!this.anchorDragging) {
      return;
    }

    const newX = event.clientX - this.pageX
    const newY = event.clientY - this.pageY

    if (anchor === 'left') {
      const originalRightEdge = this.Rect.X + this.Rect.Width;
      this.Rect.X = newX;
      this.Rect.Width = originalRightEdge - newX;
    } else if (anchor === 'right') {
      this.Rect.Width = newX - this.Rect.X;
    } else if (anchor === 'top') {
      const originalBottomEdge = this.Rect.Y + this.Rect.Height;
      this.Rect.Y = newY;
      this.Rect.Height = originalBottomEdge - newY;
    } else if (anchor === 'bottom') {
      this.Rect.Height = newY - this.Rect.Y;
    }
  }

  anchorMouseUp(event: MouseEvent): void {
    if (!event.altKey && !event.shiftKey) {
      this.anchorDragging = false;
      this.activeAnchor = null;
      this.rectService.updateRect(this.Rect, this.GONG__StackPath).subscribe()
    }
  }
}
