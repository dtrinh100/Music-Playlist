import { Component, OnInit } from '@angular/core';


const SCROLL_DELIM: Number = 100;


@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss'],
  host: {
    '(window:scroll)': 'updateNavbarBasedOnScrollEvent($event)'
  }
})
export class NavbarComponent implements OnInit {
  isMenuOpen: boolean;
  isScrolledDown: boolean;


  constructor() {
  }

  ngOnInit() {
    this.isMenuOpen = false;
    this.isScrolledDown = false;
  }

  /**
   Helps toggle open the navbar-menu. This method is hooked into the navbar-button.
   **/
  toggleMenuButton() {
    this.isMenuOpen = (this.isMenuOpen === false);
  }

  /**
   This function is called When the user scrolls around. If the user is
   scrolled to the top, the navbar is gray. Otherwise, it's blue.
   **/
  updateNavbarBasedOnScrollEvent(evt) {
    let currPos = (window.pageYOffset || evt.target.scrollTop) - (evt.target.clientTop || 0);
    this.isScrolledDown = (currPos > SCROLL_DELIM);
  }

}
