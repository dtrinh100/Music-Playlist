import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';

import { User } from '../shared/interface/user.interface';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss']
})
export class RegistrationComponent implements OnInit {
  user: FormGroup;
  User: User;

  constructor() { }

  ngOnInit() {
    this.user = new FormGroup({
     username: new FormControl(''),
     email: new FormControl(''),
     account: new FormGroup({
       password: new FormControl(''),
       confirm: new FormControl('')
     })
   });
  }

  onSubmit({ value, valid }: { value: User, valid: boolean }) {
    console.log(value, valid);
  } 

}
