import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import 'rxjs/add/operator/mergeMap';
import { Song } from '../shared/models/song';

import { SongService } from '../shared/services';

@Component({
  selector: 'app-song',
  templateUrl: './song.component.html',
  styleUrls: ['./song.component.scss']
})
export class SongComponent implements OnInit, OnDestroy {
  private song;
  private id: number;
  private sub: any;
  private audio: any;
  private isPlaying: boolean;
  private playMessage: string;
  private isSongAvailable: boolean;

  constructor(private songService: SongService, private route: ActivatedRoute) {
  }

  ngOnInit() {
    this.playMessage = 'Play';
    this.isPlaying = false;
    this.isSongAvailable = false;

    this.sub = this.route.params.subscribe(params => {
      this.id = +params['id'];
      this.songService.getSong(this.id).subscribe((song: Song) => {
        this.isSongAvailable = true;
        this.audio = new Audio(song.audiopath);
        this.song = song;
      });
    });
  }

  // Plays the selected song
  play() {
    if (this.isPlaying === true) {
      this.audio.play();
      this.playMessage = 'Stop';
    } else {
      this.audio.pause();
      this.audio.currentTime = 0;
      this.playMessage = 'Play';
    }
    this.isPlaying = !this.isPlaying;

  }

  ngOnDestroy() {
    this.sub.unsubscribe();
    if (this.audio) {
      this.audio.pause();
      this.audio = null;
    }
  }

}
