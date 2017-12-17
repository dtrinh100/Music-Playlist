import { BaseRequestOptions, HttpModule, Http, Response, ResponseOptions, ResponseType } from '@angular/http';
import { async, ComponentFixture, fakeAsync, inject, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { MockBackend } from '@angular/http/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { Subscription } from 'rxjs/Subscription';

import { UserService } from '../shared/services';
import { AuthService } from '../shared/services/auth.service';
import { ApiService } from '../shared/services';
import { LoginComponent } from './login.component';


describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  const errorResponse = {
    data: {errors: {credentials: 'Invalid Credentials'}}
  };
  const successResponse = {
    data: {errors: {credentials: 'Success'}}
  };

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LoginComponent ],
      imports: [ HttpModule, ReactiveFormsModule, RouterTestingModule ],
      providers: [
        ApiService,
        UserService,
        AuthService,
        BaseRequestOptions,
        MockBackend,
        {
          provide: Http,
          useFactory: (backend, options) => new Http(backend, options),
          deps: [ MockBackend, BaseRequestOptions ]
        }
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=- Constructors =-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('should expect hasSubmittedOnce to be false', () => {
    expect(component.hasSubmittedOnce).toBeFalsy();
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-= Form Group -=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  it('should expect form to have an email & password field', () => {
    const form = component.loginForm;
    expect(form.contains('email')).toBeTruthy();
    expect(form.contains('password')).toBeTruthy();
  });

  it('should accept valid email address', () => {
    const emailControl = component.loginForm.get('email');
    emailControl.setValue('user@email.com');
    expect(emailControl.status).toEqual('VALID')
  });

  it('should reject invalid email addresses', () => {
    const emailCtrl = component.loginForm.get('email');
    emailCtrl.setValue('user@email.');
    expect(emailCtrl.status).toEqual('INVALID');
    emailCtrl.setValue('');
    expect(emailCtrl.status).toEqual('INVALID');
  });

  it('should reject invalid password', () => {
    const passCtrl = component.loginForm.get('password');
    passCtrl.setValue('');
    expect(passCtrl.status).toEqual('INVALID');
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=- isValidEmail() =-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  it('should return true on isExclamation = true & false if isExclamation = false; on invalid data', () => {
    const emailCtrl = component.loginForm.get('email');
    emailCtrl.setValue('');
    emailCtrl.markAsDirty();
    expect(component.isValidEmail(true)).toBeTruthy();
    expect(component.isValidEmail(false)).toBeFalsy();
  });

  it('should return false on isExclamation = true & true if isExclamation = false; on valid data ', () => {
    const emailCtrl = component.loginForm.get('email');
    emailCtrl.setValue('user@email.com');
    emailCtrl.markAsDirty();
    expect(component.isValidEmail(true)).toBeFalsy();
    expect(component.isValidEmail(false)).toBeTruthy();
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-= isValidPassword() =-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  it('should return true on isExclamation = true & false if isExclamation = false; on invalid data', () => {
    const passwordCtrl = component.loginForm.get('password');
    passwordCtrl.setValue('');
    passwordCtrl.markAsDirty();
    expect(component.isValidPassword(true)).toBeTruthy();
    expect(component.isValidPassword(false)).toBeFalsy();
  });

  it('should return false on isExclamation = true & true if isExclamation = false; on valid data ', () => {
    const passwordCtrl = component.loginForm.get('password');
    passwordCtrl.setValue('pass');
    passwordCtrl.markAsDirty();
    expect(component.isValidPassword(true)).toBeFalsy();
    expect(component.isValidPassword(false)).toBeTruthy();
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=- onValueChanged() =-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  it('should expect empty-input errors', () => {
    component.onValueChanged();
    expect(component.formErrors).toEqual({
      email: 'Email is required. Incorrect email format. ',
      password: 'Password is required. '
    })
  });

  it('should expect \'incorrect email format\' error', () => {
    const emailCtrl = component.loginForm.get('email');
    const passwordCtrl = component.loginForm.get('password');

    emailCtrl.setValue('user@email.');
    passwordCtrl.setValue('password');

    component.onValueChanged();
    expect(component.formErrors).toEqual({
      email: 'Incorrect email format. ',
      password: ''
    })
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=- isInvalidForm() -=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  it('should expect no errors, so, isInvalidForm returns false', () => {
    const emailCtrl = component.loginForm.get('email');
    const passwordCtrl = component.loginForm.get('password');
    emailCtrl.setValue('user@email.com');
    passwordCtrl.setValue('password');

    component.onValueChanged();
    expect(component.isInvalidForm()).toBeFalsy();
  });

  it('should expect email error, so, isInvalidForm returns true', () => {
    const emailCtrl = component.loginForm.get('email');
    const passwordCtrl = component.loginForm.get('password');
    emailCtrl.setValue('user@email.');
    passwordCtrl.setValue('password');

    component.onValueChanged();
    expect(component.isInvalidForm()).toBeTruthy();
  });

  it('should expect password error, so, isInvalidForm returns true', () => {
    const emailCtrl = component.loginForm.get('email');
    const passwordCtrl = component.loginForm.get('password');
    emailCtrl.setValue('user@email.com');
    passwordCtrl.setValue('');

    component.onValueChanged();
    expect(component.isInvalidForm()).toBeTruthy();
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-= submitForm() -=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  it('should expect hasSubmittedOnce to be -- initially false -- but true after call to submitForm', () => {
    expect(component.hasSubmittedOnce).toBeFalsy();
    component.submitForm();
    expect(component.hasSubmittedOnce).toBeTruthy();
  });

  it('should expect onValueChanged to have been called', () => {
    const s = spyOn(component, 'onValueChanged');
    component.submitForm();
    expect(s.calls.count()).toEqual(1);
  });

  it('should expect isInvalidForm to have been called', () => {
    const s = spyOn(component, 'isInvalidForm');
    component.submitForm();
    expect(s.calls.count()).toEqual(1);
  });

  // TODO: Make DRY
  // This function helps make tests easier to read
  function mockInjectAsync(expectTests: (service, backend) => void) {
    return inject([ AuthService, MockBackend ], fakeAsync((service: AuthService, backend: MockBackend) => {
      expectTests(service, backend);
    }));
  }

  // TODO: Make DRY
  // This function helps test standard/error server-responses
  function mockBackendHelper(mockBackend, bodyStr = {}, statusNum = 200, responseType = ResponseType.Default): Subscription {
    const mockResponse = new Response(new ResponseOptions({
      body: JSON.stringify(bodyStr),
      status: statusNum,
      type: responseType
    }));

    return mockBackend.connections.subscribe(conn => {
      if (responseType === ResponseType.Default) {
        conn.mockRespond(mockResponse);
      } else if (responseType === ResponseType.Error) {
        conn.mockError(mockResponse);
      }
    });
  }

  it('should expect password input-field to clear & \'Invalid Credentials\' password error w/ wrong credentials',
    mockInjectAsync((_, mockBackend) => {
      const emailCtrl = component.loginForm.get('email');
      const passwordCtrl = component.loginForm.get('password');
      emailCtrl.setValue('user@email.com');
      passwordCtrl.setValue('wrong_password');

      mockBackendHelper(mockBackend, errorResponse, 401, ResponseType.Error);
      component.submitForm();
      expect(passwordCtrl.value).toEqual('');
      expect(component.formErrors).toEqual({
        email: '',
        password: 'Invalid credentials'
      })
    }));

  it('should expect response \'Success\' & a redirect to / w/ correct credentials', mockInjectAsync((s, b) => {
      const navigateSpy = spyOn((<any>component).router, 'navigateByUrl');
      const emailCtrl = component.loginForm.get('email');
      const passwordCtrl = component.loginForm.get('password');
      emailCtrl.setValue('user@email.com');
      passwordCtrl.setValue('password');

      mockBackendHelper(b, successResponse);
      component.submitForm();
      expect(navigateSpy).toHaveBeenCalledWith('/');
    })
  );

});
