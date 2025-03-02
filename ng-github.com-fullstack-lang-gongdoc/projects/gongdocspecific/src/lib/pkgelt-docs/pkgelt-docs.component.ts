import { Component, Input, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

import { AngularSplitModule } from 'angular-split';

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { SvgSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'


@Component({
    selector: 'lib-pkgelt-docs',
    templateUrl: './pkgelt-docs.component.html',
    styleUrls: ['./pkgelt-docs.component.css'],
    imports: [
        AngularSplitModule,
        TreeComponent,
        SvgSpecificComponent,
    ]
})
export class PkgeltDocsComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
  ) { }

  ngOnInit(): void {
    // console.log("PkgeltDocsComponent->GONG__StackPath : ", this.GONG__StackPath)
  }
}
