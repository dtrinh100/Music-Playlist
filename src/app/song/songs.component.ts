import { Component, OnInit, OnDestroy } from '@angular/core';

import 'rxjs/add/operator/mergeMap';

import { SongService } from '../shared/services/song.service';

@Component({
  selector: 'app-songs',
  templateUrl: './songs.component.html',
  styleUrls: ['./songs.component.scss']
})
export class SongsComponent implements OnInit, OnDestroy {
  private songs;
  private sub: any;

  constructor(private songService: SongService) { }

  ngOnInit() {
    this.sub = this.songService.getSongs().subscribe(data => {
      this.songs = data;
    });
  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
