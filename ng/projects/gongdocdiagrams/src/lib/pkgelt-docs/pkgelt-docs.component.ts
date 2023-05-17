import { Component, Input, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

import { AngularDragEndEventService } from '../../../../../../vendor/github.com/fullstack-lang/gongsvg/ng/projects/gongsvgspecific/src/lib/angular-drag-end-event.service';


@Component({
  selector: 'lib-pkgelt-docs',
  templateUrl: './pkgelt-docs.component.html',
  styleUrls: ['./pkgelt-docs.component.css']
})
export class PkgeltDocsComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private angularDragEndEventService: AngularDragEndEventService,
  ) { }

  ngOnInit(): void {
    console.log("PkgeltDocsComponent->GONG__StackPath : ", this.GONG__StackPath)
  }

  onDragEnd(): void {
    console.log("angular split : on drag end")
    this.angularDragEndEventService.emitMouseUpEvent(0)
  }
}
