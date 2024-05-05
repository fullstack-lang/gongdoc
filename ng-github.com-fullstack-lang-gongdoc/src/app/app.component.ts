import { Component, OnInit, ViewChild } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
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
