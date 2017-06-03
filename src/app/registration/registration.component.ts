import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators, FormControl } from '@angular/forms';

import { User } from '../shared/interface/user.interface';

import { RegistrationDirective } from './registration.directive';

// TODO write unit tests

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss']
})
export class RegistrationComponent implements OnInit {
  user: FormGroup;
  User: User;
  initialUsername: string; // either an empty initial username or get one from the session storage
  inititalEmail: string;
  initialPassword: string;
  initialConfirm: string;

  constructor(private fb: FormBuilder) { }

  ngOnInit() {

    if (sessionStorage.getItem("formUsername") === null) {
      this.initialUsername = '';
    } else {
      this.initialUsername = sessionStorage.getItem("formUsername");
    }
    if (sessionStorage.getItem("formEmail") === null) {
      this.inititalEmail = '';
    } else {
      this.inititalEmail = sessionStorage.getItem("formEmail");
    }
    if (sessionStorage.getItem("formPassword") === null) {
      this.initialPassword = '';
    } else {
      this.initialPassword = sessionStorage.getItem("formPassword");
    }
    if (sessionStorage.getItem("formConfirm") === null) {
      this.initialConfirm = '';
    } else {
      this.initialConfirm = sessionStorage.getItem("formConfirm");
    }


    this.user = this.fb.group({
      username: [this.initialUsername, [Validators.required, Validators.minLength(2), Validators.maxLength(30)]],
      email: [this.inititalEmail, [Validators.required, Validators.email]],
      account: this.fb.group({
        password: [this.initialPassword, [Validators.required, Validators.minLength(8)]],
        confirm: [this.initialConfirm, [Validators.required, this.validatePasswordConfirmation.bind(this)]]
      })

    });


    this.user.valueChanges.subscribe(value => {
      sessionStorage.setItem("formUsername", value.username);
      sessionStorage.setItem("formEmail", value.email);
      sessionStorage.setItem("formPassword", value.account.password);
      sessionStorage.setItem("formConfirm", value.account.confirm);
      this.onValueChanged(value);
    });

    this.onValueChanged();
  }

  validatePasswordConfirmation(control: FormControl): any {
  if(this.user) {
    return control.value === this.user.get('account').get('password').value ? null : { notsame: true}
  }
   }

  onSubmit({ value, valid }: { value: User, valid: boolean }) {
    console.log(value, valid);
  }

  onValueChanged(data?: any) {
    if (!this.user) { return; }
    const form = this.user;

    for (const field in this.formErrors) {
      // clear previous error message (if any)
      this.formErrors[field] = '';
      let control = form.get(field);
      if (field === "password" || field === "confirm") {
        control = form.get("account").get(field); // takes into account the nested fields
      }

      if (control && control.dirty && !control.valid) {
        const messages = this.validationMessages[field];
        for (const key in control.errors) {
          this.formErrors[field] += messages[key] + ' ';
        }
      }
    }
  }



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

  }

}
