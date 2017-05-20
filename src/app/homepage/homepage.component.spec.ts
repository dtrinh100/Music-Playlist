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

  // TODO: create test
  it('should test updateElementsBasedOnScrollEvent', () => {

  });

  // TODO: create test
  it('should test isVisible', () => {

  });
});
