import { Component, OnInit, ViewChild } from '@angular/core';

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import { PkgeltDocsComponent } from '../../projects/gongdocspecific/src/lib/pkgelt-docs/pkgelt-docs.component'

import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';

import * as gongdoc from '../../projects/gongdoc/src/public-api'

import * as svg from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svg/src/public-api'

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    imports: [
        MatRadioModule,
        FormsModule,
        CommonModule,
        MatButtonModule,
        MatIconModule,
        AngularSplitModule,

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
