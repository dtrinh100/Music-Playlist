import { Component, OnInit } from '@angular/core';
import { AuthService } from './shared/services';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {
  constructor(private authService: AuthService) {
  }
  ngOnInit(): void {
    this.authService.populate();
  }
}
