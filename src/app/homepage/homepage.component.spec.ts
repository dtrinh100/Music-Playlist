import {async, ComponentFixture, TestBed} from '@angular/core/testing';

import {HomepageComponent} from './homepage.component';
import {
  SharedModule
} from '../shared'

describe('HomepageComponent', () => {
  let component: HomepageComponent;
  let fixture: ComponentFixture<HomepageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [SharedModule],
      declarations: [HomepageComponent]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomepageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  // Detects when a scroll event happens and detects if updateElementsBasedOnScrollEvent
  // gets called.
  it('should determine if function is called when the user scrolls', () => {
    spyOn(component, 'updateElementsBasedOnScrollEvent');

    let scrollEvent = document.createEvent('CustomEvent');
    scrollEvent.initEvent('scroll', false, false);
    window.dispatchEvent(scrollEvent);

    expect(component.updateElementsBasedOnScrollEvent).toHaveBeenCalled();
  });

  // The bottom point of an element is 99 units from the window's top-point
  // Note that the window is 100 units tall.
  it('should test isVisible', () => {
    expect(component.isVisible(99, 100)).toBeTruthy();
  });

  it('should test NOT isVisible', () => {
    // The window is 100 units tall & the element's bottom point is
    // 200 units from the window's current top point.
    expect(component.isVisible(300, 100)).toBeTruthy();
  });
});
