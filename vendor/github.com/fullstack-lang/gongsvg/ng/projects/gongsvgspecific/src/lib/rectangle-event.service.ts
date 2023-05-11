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