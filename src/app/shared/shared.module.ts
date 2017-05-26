import { NgModule } from '@angular/core';

import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpModule,
    RouterModule,
    BrowserAnimationsModule
  ],
  declarations: [],
  exports: [
    CommonModule,
    FormsModule,
    HttpModule,
    ReactiveFormsModule,
    RouterModule,
    BrowserAnimationsModule
  ]
})
export class SharedModule {}
