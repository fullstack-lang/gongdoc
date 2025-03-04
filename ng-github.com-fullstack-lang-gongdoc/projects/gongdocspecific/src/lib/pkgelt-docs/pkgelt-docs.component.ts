import { Component, Input, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

import { AngularSplitModule } from 'angular-split';

import { TreeSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { SvgSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'

import * as gongdoc from '../../../../gongdoc/src/public-api'

@Component({
    selector: 'lib-pkgelt-docs',
    templateUrl: './pkgelt-docs.component.html',
    styleUrls: ['./pkgelt-docs.component.css'],
    imports: [
        AngularSplitModule,
        TreeSpecificComponent,
        SvgSpecificComponent,
    ]
})
export class PkgeltDocsComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  TreeNames = gongdoc.TreeNames

  constructor(
    private router: Router,
  ) { }

  ngOnInit(): void {
    // console.log("PkgeltDocsComponent->GONG__StackPath : ", this.GONG__StackPath)
  }
}
