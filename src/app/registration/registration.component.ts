import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { User } from '../shared/interface/user.interface';

import { RegistrationDirective } from './registration.directive';

// TODO trigger only 1 error at a time, implment a custom validator for confirms field

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss']
})
export class RegistrationComponent implements OnInit {
  user: FormGroup;
  User: User;
  initialUsername: string;
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
       password: [this.initialPassword,  [Validators.required, Validators.minLength(8)]],
       confirm: [this.initialConfirm,  [Validators.required]]
     })

  });

  this.user.valueChanges.subscribe(function(value) {
    sessionStorage.setItem("formUsername", value.username);
    sessionStorage.setItem("formEmail", value.email);
    sessionStorage.setItem("formPassword", value.account.password);
    sessionStorage.setItem("formConfirm", value.account.confirm);
  });


  }

  onSubmit({ value, valid }: { value: User, valid: boolean }) {
    console.log(value, valid);
  }

}
