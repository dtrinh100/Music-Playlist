import { NgModule }              from '@angular/core';
import { RouterModule, Routes }  from '@angular/router';

import { LoginComponent } from './login/login.component';
import { RegistrationComponent }   from './registration/registration.component';
import { HomepageComponent } from './homepage/homepage.component';
import { Status404Component } from './status-404/status-404.component';


const appRoutes: Routes = [
   { path: 'login', component: LoginComponent },
   { path: 'register', component: RegistrationComponent },
   { path: '', component: HomepageComponent },
   { path: '**', component: Status404Component}
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
