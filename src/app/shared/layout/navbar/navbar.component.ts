import { Component, HostListener, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services';


const SCROLL_DELIM: Number = 100;


@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: [ './navbar.component.scss' ],
})
export class NavbarComponent implements OnInit {
  isMenuOpen: boolean;
  isScrolledDown: boolean;

  constructor(private authService: AuthService,
              private router: Router) {
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
  @HostListener('window:scroll', ['$event']) updateNavbarBasedOnScrollEvent(evt) {
    const currPos = (window.pageYOffset || evt.target.scrollTop) - (evt.target.clientTop || 0);
    this.isScrolledDown = (currPos > SCROLL_DELIM);
  }

  /**
   logOut logs the user off of the API.
   **/
  logOut() {
    this.authService.logout().subscribe(_ => {
      this.router.navigateByUrl('/');
    })
  }

}
