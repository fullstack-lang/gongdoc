import { Component } from '@angular/core';

import * as gongdoc from 'gongdoc';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {

  // choices for the top radio button
  view = 'Default view'
  default = 'Default view'
  meta = 'Meta view'
  gong = 'Gong view'
  views: string[] = [this.default, this.meta, this.gong];

  diagramPackage: gongdoc.DiagramPackageDB = new (gongdoc.DiagramPackageDB);

  constructor(
    private diagramPackageService: gongdoc.DiagramPackageService,
  ) {
    // create a new GongDoc instance
    this.diagramPackageService.getDiagramPackages().subscribe(diagramPackages => {
      this.diagramPackage = diagramPackages[0];
    }
    )
  }

  refresh() {
    // refresh the view
    this.diagramPackage.IsReloaded = true
    this.diagramPackageService.updateDiagramPackage(this.diagramPackage).subscribe(diagramPackage => {
      console.log('diagram package refreshed')
    })
  }
}
