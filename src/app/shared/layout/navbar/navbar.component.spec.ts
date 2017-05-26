import {async, ComponentFixture, TestBed} from '@angular/core/testing';
import {ModuleWithProviders} from "@angular/core";
import {RouterModule} from "@angular/router";

import {NavbarComponent} from './navbar.component';


const rootRouting: ModuleWithProviders = RouterModule.forRoot([], {useHash: true});

describe('NavbarComponent', () => {
  let component: NavbarComponent;
  let fixture: ComponentFixture<NavbarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [RouterModule, rootRouting],
      declarations: [NavbarComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should toggle menu-button open', () => {
    let isInitiallyClosed = !component.isMenuOpen;
    component.toggleMenuButton();   // Open menu-bar

    let expected = isInitiallyClosed && component.isMenuOpen;
    expect(expected).toBeTruthy()
  });

  it('should indicate \'not scrolled-down\'', () => {
    let event: any = {
      target: {
        scrollTop: 0
      }
    };

    component.updateNavbarBasedOnScrollEvent(event);
    expect(component.isScrolledDown).toBeFalsy();
  });

  it('should indicate \'scrolled-down\'', () => {
    let event: any = {
      target: {
        scrollTop: 500
      }
    };

    component.updateNavbarBasedOnScrollEvent(event);
    expect(component.isScrolledDown).toBeTruthy();
  })
});
