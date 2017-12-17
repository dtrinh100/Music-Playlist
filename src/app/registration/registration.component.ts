import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup, Validators, FormControl } from '@angular/forms';

import { UserInterface } from '../shared/interface/userinterface.interface';
import { User } from '../shared/models/user';

import { UserService } from '../shared/services';
import { AuthService } from '../shared/services/auth.service';

// TODO tests form submission, mock the service

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: [ './registration.component.scss' ],
  providers: [ UserService ]
})
export class RegistrationComponent implements OnInit {
  registrationForm: FormGroup;
  User: UserInterface; // The UserInterface is an interface used specifcally only for forms
  UserModel: User; // The User type is used as a model for REST API calls, not to be confused with the type UserInterface
  initialUsername: string; // either an empty initial username or get one from the session storage
  inititalEmail: string;
  initialPassword: string;
  initialConfirm: string;
  payload: object;

  formErrors = {
    'username': '',
    'email': '',
    'password': '',
    'confirm': ''

  };

  validationMessages = {
    'username': {
      'required': 'A username is required',
      'minlength': 'Must have a minimum of 2 characters',
      'maxlength': 'Username can only have a maximum of 30 characters'
    },
    'email': {
      'required': 'An email address is required',
      'email': 'Invalid email address'
    },

    'password': {
      'required': 'Password is required',
      'minlength': 'Must have a minimum of 8 characters'
    },
    'confirm': {
      'required': 'Please confirm your password',
      'notsame': 'Confirm does not match with password'
    }
  };

  constructor(private fb: FormBuilder,
              private authService: AuthService,
              private router: Router) {
  }

  ngOnInit() {

    if (sessionStorage.getItem('formUsername') === null) {
      this.initialUsername = '';
    } else {
      this.initialUsername = sessionStorage.getItem('formUsername');
    }
    if (sessionStorage.getItem('formEmail') === null) {
      this.inititalEmail = '';
    } else {
      this.inititalEmail = sessionStorage.getItem('formEmail');
    }
    if (sessionStorage.getItem('formPassword') === null) {
      this.initialPassword = '';
    } else {
      this.initialPassword = sessionStorage.getItem('formPassword');
    }
    if (sessionStorage.getItem('formConfirm') === null) {
      this.initialConfirm = '';
    } else {
      this.initialConfirm = sessionStorage.getItem('formConfirm');
    }


    this.registrationForm = this.fb.group({
      username: [ this.initialUsername, [ Validators.required, Validators.minLength(2), Validators.maxLength(30) ] ],
      email: [ this.inititalEmail, [ Validators.required, Validators.email ] ],
      account: this.fb.group({
        password: [ this.initialPassword, [ Validators.required, Validators.minLength(8) ] ],
        confirm: [ this.initialConfirm, [ Validators.required, this.validatePasswordConfirmation.bind(this) ] ]
      })

    });

    this.registrationForm.valueChanges.subscribe(value => {
      sessionStorage.setItem('formUsername', value.username);
      sessionStorage.setItem('formEmail', value.email);
      sessionStorage.setItem('formPassword', value.account.password);
      sessionStorage.setItem('formConfirm', value.account.confirm);
      this.onValueChanged(value);
    });

    this.onValueChanged();
  }

  validatePasswordConfirmation(control: FormControl): any {
    if (this.registrationForm) {
      return control.value === this.registrationForm.get('account').get('password').value ? null : {notsame: true}
    }
  }

  onValueChanged(data?: any) {
    if (!this.registrationForm) {
      return;
    }
    const form = this.registrationForm;

    for (const field in this.formErrors) {
      if (!this.formErrors.hasOwnProperty(field)) {
        continue;
      }
      // clear previous error message (if any)
      this.formErrors[ field ] = '';
      let control = form.get(field);
      if (field === 'password' || field === 'confirm') {
        control = form.get('account').get(field); // takes into account the nested fields
      }

      if (control && control.dirty && !control.valid) {
        const messages = this.validationMessages[ field ];
        for (const key in control.errors) {
          if (!control.errors.hasOwnProperty(key)) {
            continue;
          }
          this.formErrors[ field ] += messages[ key ] + ' ';
        }
      }
    }
  }

  onSubmit({value, valid}: {value: UserInterface, valid: boolean}) {
    this.payload = {
      username: value.username,
      email: value.email,
      password: value.account.password,
      confirm: value.account.confirm
    };
    this.authService.register(this.payload).subscribe((res: any) => {
      this.router.navigateByUrl('/');
    }, err => {
      const form = this.registrationForm;
      const ctrl = function (ctrlStr) {
        return form.get(ctrlStr);
      };
      const acctCtrl = function (ctrlStr) {
        return form.get('account').get(ctrlStr);
      };
      [ ctrl('username'),
        ctrl('email'),
        acctCtrl('password'),
        acctCtrl('confirm') ].forEach(control => {
        control.setValue('');
      });
      this.formErrors.username = err.json().errors.Register;
    });
  }

}
