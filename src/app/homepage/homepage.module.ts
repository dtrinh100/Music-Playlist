import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { HomepageComponent } from './homepage.component';
import { SharedModule } from '../shared';


const homeRouting: ModuleWithProviders = RouterModule.forChild([
  {
    path: '',
    component: HomepageComponent
  }
]);


@NgModule({
  imports: [
    homeRouting,
    SharedModule
  ],
  declarations: [
    HomepageComponent
  ],
  providers: []
})
export class HomeModule {}
