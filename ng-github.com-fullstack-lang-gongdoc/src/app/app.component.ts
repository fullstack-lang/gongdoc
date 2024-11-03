import { Component, OnInit, ViewChild } from '@angular/core';

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import { PanelComponent } from '../../projects/gongdocspecific/src/public-api';

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { MaterialTableComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-table/material-table.component';
import { MaterialFormComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-form/material-form.component';
import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';
import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'

import * as gongdoc from '../../projects/gongdoc/src/public-api'
import * as gongsvg from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvg/src/public-api'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: true,
  imports: [
    MatRadioModule,
    FormsModule,
    CommonModule,
    MatButtonModule,
    MatIconModule,

    AngularSplitModule,

    TreeComponent,
    MaterialTableComponent,
    MaterialFormComponent,

    PanelComponent,
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
  SVGStackType = gongsvg.StackType


  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
