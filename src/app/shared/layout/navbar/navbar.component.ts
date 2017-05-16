import { Component, OnInit } from '@angular/core';


const SCROLL_DELIM: Number = 100;

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss'],
  host: {
    '(window:scroll)': 'updateHeader($event)'
  }
})
export class NavbarComponent implements OnInit {
  isMenuOpen: boolean;
  isScrolled: boolean;


  constructor() {
    this.isMenuOpen = false;
    this.isScrolled = false;
  }

  ngOnInit() { }

  toggleState() {
    this.isMenuOpen = (this.isMenuOpen === false);
  }

  updateHeader(evt) {
    let currPos = (window.pageYOffset || evt.target.scrollTop) - (evt.target.clientTop || 0);
    this.isScrolled = (currPos > SCROLL_DELIM);
  }

}
