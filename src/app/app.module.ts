import { BrowserModule } from '@angular/platform-browser';
import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from "@angular/router";
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { AppRoutingModule } from './app-routing.module';

import { LoginComponent } from './login/login.component';
import { AppComponent } from './app.component';
import { RegistrationComponent } from './registration/registration.component';
import { HomepageComponent } from './homepage/homepage.component';
import { Status404Component } from './status-404/status-404.component';

import { RegistrationDirective } from './registration/registration.directive';

import {
  SharedModule,
  NavbarComponent,
  FooterComponent,
  UserService,
  ApiService,
  User
} from './shared';


@NgModule({
  declarations: [
    AppComponent,
    RegistrationComponent,
    NavbarComponent,
    FooterComponent,
    HomepageComponent,
    Status404Component,
    RegistrationDirective,
    LoginComponent
  ],
  imports: [
    AppRoutingModule,
    SharedModule
  ],
  providers: [
    ApiService,
    UserService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
