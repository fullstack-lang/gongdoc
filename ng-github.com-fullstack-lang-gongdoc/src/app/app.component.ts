import { Component, OnInit, ViewChild } from '@angular/core';

import { PkgeltDocsComponent } from '../../projects/gongdocspecific/src/lib/pkgelt-docs/pkgelt-docs.component'

import * as gongdoc from '../../projects/gongdoc/src/public-api'

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    imports: [
        PkgeltDocsComponent,

    ]
})
export class AppComponent implements OnInit {

  StackType = gongdoc.StackType

  ngOnInit(): void {

  }
}
