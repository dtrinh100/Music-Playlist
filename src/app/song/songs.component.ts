import { Component, OnInit } from '@angular/core';

import 'rxjs/add/operator/mergeMap';

import {ApiService} from '../shared/services';

@Component({
  selector: 'app-songs',
  templateUrl: './songs.component.html',
  styleUrls: ['./songs.component.scss']
})
export class SongsComponent implements OnInit {
  private images;
  private sub: any;

  constructor() { }

  ngOnInit() {
    this.sub = this.apiService.get("/songs/" + this.id).subscribe(data => {
      this.images = data;
    });
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
