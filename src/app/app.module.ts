import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';

import { LoginComponent } from './login/login.component';
import { AppComponent } from './app.component';
import { RegistrationComponent } from './registration/registration.component';
import { HomepageComponent } from './homepage/homepage.component';
import { SongsComponent } from './song/songs.component';
import { SongComponent } from './song/song.component';

import { Status404Component } from './status-404/status-404.component';

import { RegistrationDirective } from './registration/registration.directive';

import {
  SharedModule,
  NavbarComponent,
  FooterComponent,
  UserService,
  AuthService,
  ApiService,
  SongService,
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
    LoginComponent,
    SongsComponent,
    SongComponent
  ],
  imports: [
    AppRoutingModule,
    SharedModule
  ],
  providers: [
    ApiService,
    AuthService,
    UserService,
    SongService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
