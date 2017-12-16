// TODO at a later date if there is time, unit test to make sure that component.formError[field] has a value, not sure how to detect
// change yet

import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import {
    FormGroup,
    ReactiveFormsModule
} from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterTestingModule } from "@angular/router/testing";


import { RegistrationComponent } from './registration.component';

import { RegistrationDirective } from './registration.directive';

import { UserService } from '../shared/services';
import { ApiService } from '../shared/services';
import { Observable } from "rxjs/Observable";

describe('RegistrationComponent', () => {
  let component: RegistrationComponent;
  let fixture: ComponentFixture<RegistrationComponent>;
  let userServiceStub: object;
  let userService: any;
  let testResponse: any;
  let spy: any;


  beforeEach(async(() => {

    TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, HttpModule, RouterTestingModule],
      declarations: [ RegistrationComponent, RegistrationDirective ],
      providers: [ ApiService, UserService ]
    }).compileComponents();
 }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RegistrationComponent);
    component = fixture.componentInstance;
    userService = fixture.debugElement.injector.get(UserService);
    testResponse = {errors: "", success: false};
    spy = spyOn(userService, 'createUser').and.returnValue(Observable.of(testResponse));
    fixture.detectChanges();
  });

  it('should create the registration component', () => {
    expect(component).toBeTruthy();
  });

  it('should create a FormGroup', () => {
    component.ngOnInit();
    expect(component.registrationForm instanceof FormGroup).toBe(true);
  });

  it('should reject username that are less than 2 characters long', () => {
    component.ngOnInit();
    component.registrationForm.controls["username"].setValue("n");
    expect(component.registrationForm.get("username").status).toBe("INVALID");
  });

  it('should reject username that are more than 30 characters long', () => {
    component.ngOnInit();
    component.registrationForm.controls["username"].setValue("nijijijijijijijijijijijijijijij");
    expect(component.registrationForm.get("username").status).toBe("INVALID");
  });

  it('should accept username that are between 2 and 30 characters long', () => {
    component.ngOnInit();
    component.registrationForm.controls["username"].setValue("nijijijijijijijijijijijijijiji");
    expect(component.registrationForm.get("username").status).toBe("VALID");
  });

  it('should reject email that are not valid', () => {
    component.ngOnInit();
    component.registrationForm.controls["email"].setValue("fakerEmail");
    expect(component.registrationForm.get("email").status).toBe("INVALID");
  });

  it('should accept email that are valid', () => {
    component.ngOnInit();
    component.registrationForm.controls["email"].setValue("joe.doe@gmail.com");
    expect(component.registrationForm.get("email").status).toBe("VALID");
  });

  it('should reject password that are less than 8 characters in length', () => {
    component.ngOnInit();
    component.registrationForm.get("account").get("password").setValue("pas");
    expect(component.registrationForm.get("account").get("password").status).toBe("INVALID");
  });

  it('should accept password that are at least 8 characters in length', () => {
    component.ngOnInit();
    component.registrationForm.get("account").get("password").setValue("password");
    expect(component.registrationForm.get("account").get("password").status).toBe("VALID");
  });

  it('should reject password confirmation that are not the same as the original password', () => {
    component.ngOnInit();
    component.registrationForm.get("account").get("password").setValue("password");
    component.registrationForm.get("account").get("confirm").setValue("password2");
    expect(component.registrationForm.get("account").get("confirm").status).toBe("INVALID");
  });

  it('should accept password confirmation that are the same as the original password', () => {
    component.ngOnInit();
    component.registrationForm.get("account").get("password").setValue("password");
    component.registrationForm.get("account").get("confirm").setValue("password");
    expect(component.registrationForm.get("account").get("confirm").status).toBe("VALID");
  });

  it('should expect the userService to be called once upon onSubmit()', () => {
    component.ngOnInit();

    component.registrationForm.controls["username"].setValue("nijijijijijijijijijijijijijijij");
    component.registrationForm.controls["email"].setValue("joe.doe@gmail.com");
    component.registrationForm.get("account").get("password").setValue("password");
    component.registrationForm.get("account").get("confirm").setValue("password");

    component.onSubmit(component.registrationForm);
    expect(spy.calls.count()).toBe(1, 'stubbed method was called once');

  });

  it('should expect the userService to be called with specific parameters upon onSubmit()', () => {
    component.ngOnInit();

    component.registrationForm.controls["username"].setValue("nijijijijijijijijijijijijijijij");
    component.registrationForm.controls["email"].setValue("joe.doe@gmail.com");
    component.registrationForm.get("account").get("password").setValue("password");
    component.registrationForm.get("account").get("confirm").setValue("password");

    component.onSubmit(component.registrationForm);
    expect((spy)).toHaveBeenCalledWith("nijijijijijijijijijijijijijijij", "joe.doe@gmail.com", "password", "password");
  });


});
