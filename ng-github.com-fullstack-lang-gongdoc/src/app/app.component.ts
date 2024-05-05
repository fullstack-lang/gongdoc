import { Component, OnInit, ViewChild } from '@angular/core';

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import { PanelComponent } from '../../projects/gongdocspecific/src/public-api';

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

    PanelComponent,
  ]
})
export class AppComponent implements OnInit {

  default = 'DOC Data/Model'
  diagrammer = 'Diagram Edition'
  view = this.diagrammer

  views: string[] = [this.diagrammer];

  GONG__MODEL__StacksPath = "github.com/fullstack-lang/gongdoc/go/models"
  GONG__DATA__StackPath = "gongdoc"

  loading = true

  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
