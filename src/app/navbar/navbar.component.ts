import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  isMenuOpen = false;

  constructor() { }

  ngOnInit() { }

  toggleState() {
    this.isMenuOpen = (this.isMenuOpen === false);
  }
}
