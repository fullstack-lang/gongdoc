import { Component, OnInit } from '@angular/core';

import * as gongdoc from 'gongdoc';
import { StackConfigs, StacksService } from './stacks.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  // choices for the top radio button
  view = 'Default view'
  default = 'Default view'
  meta = 'Meta view'
  gong = 'Gong view'
  views: string[] = [this.default, this.meta, this.gong];

  stacks: string[] = []
  // soloStack : string = "fullstack-lang/gongdoc/go/tests/geometry/go/models"
  soloStack : string = ""
  

  diagramPackage: gongdoc.DiagramPackageDB = new (gongdoc.DiagramPackageDB);

  loading = true

  constructor(
    private diagramPackageService: gongdoc.DiagramPackageService,
    private stacksService: StacksService
  ) {

  }

  ngOnInit(): void {

    // get all stacks to analyse
    this.stacksService.getStacks().subscribe(
      (stacks : string[]) => {
        for (let stack of stacks) {
          console.log( "Gongdoc component Stack ", stack)
        }
        this.stacks = stacks
        this.soloStack = this.stacks[0]
        console.log( "Gongdoc AppComponent solo stack ", this.soloStack)
        this.loading = false
      }
    )

    // create a new GongDoc instance
    this.diagramPackageService.getDiagramPackages().subscribe(diagramPackages => {
      this.diagramPackage = diagramPackages[0];
    }
    )
  }

  refresh() {
    // refresh the view
    this.diagramPackage.IsReloaded = true
    this.diagramPackageService.updateDiagramPackage(this.diagramPackage, "").subscribe(diagramPackage => {
      console.log('diagram package refreshed')
    })
  }
}
