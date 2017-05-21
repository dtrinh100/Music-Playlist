import { BrowserModule } from '@angular/platform-browser';
import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from "@angular/router";

import { AppComponent } from './app.component';
import { HomeModule } from './homepage/homepage.module';
import {
  SharedModule,
  NavbarComponent,
  FooterComponent
} from './shared'


const rootRouting: ModuleWithProviders = RouterModule.forRoot([], { useHash: true });


@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    FooterComponent
  ],
  imports: [
    BrowserModule,
    rootRouting,
    SharedModule,
    HomeModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
