import { NgModule } from '@angular/core';
import { GongsvgspecificComponent } from './gongsvgspecific.component';
import { LayerComponent } from './layer/layer.component';
import { RectComponent } from './rect/rect.component';
import { BrowserModule } from '@angular/platform-browser';
import { TextComponent } from './text/text.component';
import { CircleComponent } from './circle/circle.component';
import { LineComponent } from './line/line.component';
import { EllipseComponent } from './ellipse/ellipse.component';
import { PolylineComponent } from './polyline/polyline.component';
import { PathComponent } from './path/path.component';
import { PolygoneComponent } from './polygone/polygone.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component'

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

import { GongsvgModule } from 'gongsvg';
import { SvgComponent } from './svg/svg.component';
import { LinkComponent } from './link/link.component';



@NgModule({
  declarations: [
    GongsvgspecificComponent,
    LayerComponent,
    RectComponent,
    TextComponent,
    CircleComponent,
    LineComponent,
    EllipseComponent,
    PolylineComponent,
    PathComponent,
    PolygoneComponent,
    DataModelPanelComponent,
    SvgComponent,
    LinkComponent
  ],
  imports: [
    BrowserModule,
    GongdocModule,
    GongdocdiagramsModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongsvgModule
  ],
  exports: [
    GongsvgspecificComponent,
    SvgComponent,
    DataModelPanelComponent,
  ]
})
export class GongsvgspecificModule { }
