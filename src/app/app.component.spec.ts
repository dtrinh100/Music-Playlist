import {TestBed, async} from '@angular/core/testing';
import {ModuleWithProviders} from '@angular/core';
import {RouterModule} from "@angular/router";
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

import {AppComponent} from './app.component';
import {
  NavbarComponent,
  FooterComponent,
} from './shared'

const rootRouting: ModuleWithProviders = RouterModule.forRoot([], {useHash: true});

describe('AppComponent', () => {
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        RouterModule,
        rootRouting,
        BrowserAnimationsModule,
      ],
      declarations: [
        AppComponent,
        NavbarComponent,
        FooterComponent,
      ],
    }).compileComponents();
  }));

  it('should create the app', async(() => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));


});
