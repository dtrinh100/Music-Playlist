import { NgModule } from '@angular/core';

import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { ShowAuthedDirective } from './directive'

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpModule,
    RouterModule,
    BrowserAnimationsModule,
  ],
  declarations: [
    ShowAuthedDirective,
  ],
  exports: [
    // Modules
    CommonModule,
    FormsModule,
    HttpModule,
    ReactiveFormsModule,
    RouterModule,
    BrowserAnimationsModule,
    // Directives
    ShowAuthedDirective,
  ]
})
export class SharedModule {
}
