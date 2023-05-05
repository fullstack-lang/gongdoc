import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { Coordinate } from './rectangle-event.service';

export enum SweepDirection {
  LEFT_TO_RIGHT = "LEFT_TO_RIGHT",
  RIGHT_TO_LEFT = "RIGHT_TO_LEFT",
}

export class SelectAreaConfig {
  TopLeft: Coordinate = [0, 0]
  BottomRigth: Coordinate = [0, 0]
  SweepDirection: SweepDirection = SweepDirection.LEFT_TO_RIGHT
}

@Injectable({
  providedIn: 'root'
})
export class SvgEventService {

  private mouseShiftKeyMouseUpEventSource = new Subject<Coordinate>();
  mouseShiftKeyMouseUpEvent$ = this.mouseShiftKeyMouseUpEventSource.asObservable();
  emitMouseShiftKeyMouseUpEvent(Coordinate: Coordinate) {
    this.mouseShiftKeyMouseUpEventSource.next(Coordinate);
  }

  private multiShapeSelectEndSource = new Subject<SelectAreaConfig>();
  multiShapeSelectEndEvent$ = this.multiShapeSelectEndSource.asObservable();
  emitMultiShapeSelectEnd(selectAreaConfig: SelectAreaConfig) {
    this.multiShapeSelectEndSource.next(selectAreaConfig);
  }

  private multiShapeSelectDragSource = new Subject<Coordinate>();
  multiShapeSelectDragEvent$ = this.multiShapeSelectDragSource.asObservable();
  emitMultiShapeSelectDrag(Coordinate: Coordinate) {
    this.multiShapeSelectDragSource.next(Coordinate);
  }
}
