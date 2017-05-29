import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

import { User } from '../shared/interface/user.interface';

// TODO trigger only 1 error at a time

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss']
})
export class RegistrationComponent implements OnInit {
  user: FormGroup;
  User: User;

  constructor(private fb: FormBuilder) { }

  ngOnInit() {
    this.user = this.fb.group({
     username: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(30)]],
     email: ['', [Validators.required, Validators.email]],
     account: this.fb.group({
       password: ['',  [Validators.required, Validators.minLength(8)]],
       confirm: ['',  [Validators.required]]
     })
   });
  }

  onSubmit({ value, valid }: { value: User, valid: boolean }) {
    console.log(value, valid);
  }

}
