import { Component } from '@angular/core';

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
}
