import { NgModule }              from '@angular/core';
import { RouterModule, Routes }  from '@angular/router';

import { RegistrationComponent }   from './registration/registration.component';
import { HomepageComponent } from './homepage/homepage.component';



const appRoutes: Routes = [

   { path: 'register', component: RegistrationComponent },
   { path: '', component: HomepageComponent }
];
@NgModule({
  imports: [
    RouterModule.forRoot(appRoutes)
  ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule {}
