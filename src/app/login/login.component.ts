import { FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";
import { Component, OnInit } from '@angular/core';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;

  constructor(
    private fb: FormBuilder
  ) { }

  ngOnInit() {
    this.loginForm = this.fb.group({
      'email':    ['', [Validators.required, Validators.email]],
      'password': ['', [Validators.required]]
    })
  }

}
