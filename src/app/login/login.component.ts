import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {Component, OnInit} from '@angular/core';
import {Router} from "@angular/router";
import {Subscription} from "rxjs/Subscription";

import {AuthService} from "../shared/services/auth.service";


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  hasSubmittedOnce: boolean;
  formValueChangesSub: Subscription;
  loginForm: FormGroup;

  formErrors = {
    'email': '',
    'password': '',
  };

  validationMessages = {
    'email': {
      'required': 'Email is required.',
      'email': 'Incorrect email format.',
    },
    'password': {
      'required': 'Password is required.'
    }
  };

  constructor(private fb: FormBuilder,
              private authService: AuthService,
              private router: Router) {
  }

  ngOnInit() {
    this.hasSubmittedOnce = false;
    this.loginForm = this.fb.group({
      'email': ['', [Validators.required, Validators.email]],
      'password': ['', [Validators.required]]
    });
  }

  isValidEmail(isExclamation: boolean): boolean {
    let emailForm = this.loginForm.get('email');
    if (isExclamation) {
      return emailForm.status === 'INVALID' && emailForm.dirty
    }
    return emailForm.status === 'VALID' && emailForm.dirty;
  }

  isValidPassword(isExclamation: boolean): boolean {
    let passwordForm = this.loginForm.get('password');
    if (isExclamation) {
      return passwordForm.status === 'INVALID' && passwordForm.dirty
    }
    return passwordForm.status === 'VALID' && passwordForm.dirty;
  }

  onValueChanged() {
    if (this.loginForm == undefined) {
      return;
    }

    for (const field in this.loginForm.controls) {
      if (!this.loginForm.controls.hasOwnProperty(field)) {
        continue;
      }
      const control = this.loginForm.get(field);

      // Clear previous error messages
      this.formErrors[field] = '';

      if (control && !control.valid) {
        let messages = this.validationMessages[field];

        for (const key in control.errors) {
          if (!control.errors.hasOwnProperty(key)) {
            continue;
          }
          this.formErrors[ field ] += messages[ key ] + ' ';
        }
      }
    }
  }

  isInvalidForm(): boolean {
    for (const field in this.formErrors) {
      if (this.formErrors[field].length > 0) {
        return true;
      }
    }

    return false;
  }

  submitForm() {
    if (!this.hasSubmittedOnce) {
      this.formValueChangesSub = this.loginForm.valueChanges.subscribe(_ => this.onValueChanged());
      this.hasSubmittedOnce = true;
    }

    this.onValueChanged();
    if (this.isInvalidForm()) {
      return;
    }

    const credentials = this.loginForm.value;
    this.authService.login(credentials).subscribe(_ => {
      this.router.navigateByUrl('/');
    }, (errResponse: Response) => {
      if (errResponse.status === 401) {
        this.loginForm.get('password').setValue('');
        this.formErrors.password = 'Invalid credentials';
      }
    });
  }

}
