import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { ShapeMouseEvent } from './shape.mouse.event';

export type Coordinate = [number, number]

// Define the interface for the event object
interface RectMouseEvent {
  rectangleID: number;
  MousePosRelativeSVG: [number, number]
}

@Injectable({
  providedIn: 'root'
})
export class RectangleEventService {

  //
  // mouse events
  //
  private mouseDownEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseDownEvent$ = this.mouseDownEventSource.asObservable();
  emitMouseDownEvent(shapeMouseEvent: ShapeMouseEvent) {
    // console.log('RectangleEventService, rect mouse down event, rectangle', rectangleID)
    this.mouseDownEventSource.next(shapeMouseEvent);
  }

  private mouseMoveEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseMoveEvent$ = this.mouseMoveEventSource.asObservable();
  emitMouseMoveEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('RectangleEventService, rect mouse drag event, rectangle', rectangleID)
    this.mouseMoveEventSource.next(ShapeMouseEvent);
  }

  private mouseMouseUpEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseUpEvent$ = this.mouseMouseUpEventSource.asObservable();
  emitMouseUpEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('RectangleEventService, rect mouse down event, rectangle', rectangleID)
    this.mouseMouseUpEventSource.next(ShapeMouseEvent);
  }

  //
  // mouse ALT events
  //

  private mouseRectAltKeyMouseDownEventSource = new Subject<RectMouseEvent>();
  mouseRectAltKeyMouseDownEvent$ = this.mouseRectAltKeyMouseDownEventSource.asObservable();
  emitRectAltKeyMouseDownEvent(rectangleID: number, coordinate: [number, number]) {
    this.mouseRectAltKeyMouseDownEventSource.next({ rectangleID, MousePosRelativeSVG: coordinate });
  }

  private mouseRectAltKeyMouseUpEventSource = new Subject<number>();
  mouseRectAltKeyMouseUpEvent$ = this.mouseRectAltKeyMouseUpEventSource.asObservable();
  emitMouseRectAltKeyMouseUpEvent(rectangleID: number) {
    this.mouseRectAltKeyMouseUpEventSource.next(rectangleID);
  }

  private mouseRectAltKeyMouseDragEventSource = new Subject<Coordinate>()
  mouseRectAltKeyMouseDragEvent$ = this.mouseRectAltKeyMouseDragEventSource.asObservable()
  emitRectAltKeyMouseDragEvent(coordinate: Coordinate) {
    this.mouseRectAltKeyMouseDragEventSource.next(coordinate)
  }
}