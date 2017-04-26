import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  isIn = false;

  constructor() { }

  ngOnInit() { }

  toggleState() {
    this.isIn = (this.isIn === false);
  }
}
