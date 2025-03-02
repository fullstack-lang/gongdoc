import { Component, OnInit, ViewChild } from '@angular/core';

import { PkgeltDocsComponent } from '../../projects/gongdocspecific/src/lib/pkgelt-docs/pkgelt-docs.component'

import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';

import * as gongdoc from '../../projects/gongdoc/src/public-api'

import * as svg from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svg/src/public-api'

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    imports: [
        PkgeltDocsComponent,

    ]
})
export class AppComponent implements OnInit {

  default = 'DOC Data/Model'
  diagrammer = 'Diagram Edition'
  probe = "probe"
  view = this.probe

  views: string[] = [this.diagrammer, this.probe];

  GONG__MODEL__StacksPath = "github.com/fullstack-lang/gongdoc/go/models"
  GONG__DATA__StackPath = "gongdoc"

  loading = true

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackType = gongdoc.StackType

  StackName = "gongdoc"
  TableExtraPathEnum = gongtable.TableExtraPathEnum
  StacksNames = gongdoc.StacksNames
  SVGStackType = svg.StackType


  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
