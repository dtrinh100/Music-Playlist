import { Component, OnInit, OnDestroy } from '@angular/core';

import 'rxjs/add/operator/mergeMap';
import { Song } from '../shared/models/song';

import { SongService } from '../shared/services';

@Component({
  selector: 'app-songs',
  templateUrl: './songs.component.html',
  styleUrls: ['./songs.component.scss']
})
export class SongsComponent implements OnInit, OnDestroy {
  private songs;
  private getSongsSubscription: any;

  constructor(private songService: SongService) { }

  ngOnInit() {
    this.getSongsSubscription = this.songService.getSongs().subscribe((songs: Song[]) => {
      this.songs = songs;
    });
  }

  ngOnDestroy() {
    this.getSongsSubscription.unsubscribe();
  }

}
