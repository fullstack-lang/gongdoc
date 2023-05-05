import { Injectable } from '@angular/core';
import { PointDB } from 'gongsvg';
import { Subject } from 'rxjs';
import { ShapeMouseEvent } from './shape.mouse.event';

@Injectable({
  providedIn: 'root'
})
export class LinkEventService {

  //
  // mouse events
  //
  private mouseMouseDownEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseDownEvent$ = this.mouseMouseDownEventSource.asObservable();
  emitMouseDownEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('LinkEventService, mouse down event', ShapeMouseEvent.shapeID)
    this.mouseMouseDownEventSource.next(ShapeMouseEvent);
  }


  private mouseMouseMoveEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseMoveEvent$ = this.mouseMouseMoveEventSource.asObservable();
  emitMouseMoveEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('LinkEventService, mouse move event', ShapeMouseEvent.shapeID)
    this.mouseMouseMoveEventSource.next(ShapeMouseEvent);
  }

  private mouseMouseLeaveEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseLeaveEvent$ = this.mouseMouseLeaveEventSource.asObservable();
  emitMouseLeaveEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('LinkEventService, mouse Leave event', ShapeMouseEvent.shapeID)
    this.mouseMouseLeaveEventSource.next(ShapeMouseEvent);
  }

  private mouseMouseUpEventSource = new Subject<ShapeMouseEvent>();
  mouseMouseUpEvent$ = this.mouseMouseUpEventSource.asObservable();
  emitMouseUpEvent(ShapeMouseEvent: ShapeMouseEvent) {
    // console.log('angleEventService, mouse down event, angle', ShapeMouseEvent)
    this.mouseMouseUpEventSource.next(ShapeMouseEvent);
  }
}
