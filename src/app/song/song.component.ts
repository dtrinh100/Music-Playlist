import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import 'rxjs/add/operator/mergeMap';

import {ApiService} from '../shared/services';

@Component({
  selector: 'app-song',
  templateUrl: './song.component.html',
  styleUrls: ['./song.component.scss']
})
export class SongComponent implements OnInit {

  private song;
  private id: number;
  private sub: any;
  private audio: any;
  private isPlaying: boolean;
  private playMessage: string;

  constructor(private apiService: ApiService, private route: ActivatedRoute) { }

  ngOnInit() {
    this.playMessage = "Play";
    this.isPlaying = false;
    this.sub = this.route.params.subscribe(params => {
      this.id = +params["id"];
      this.apiService.get("/songs/" + this.id).subscribe(data => {
        this.audio = new Audio(data.audioPath);
        this.song = {
          id: this.id,
          name: data.name,
          imgURL: data.imgURL,
          alt: data.alt,
          description: data.description,
          credit: data.credit
        }
      });
    });
  }

  // Plays the selected song
  play() {

    if (this.isPlaying === true) {
      this.audio.play();
      this.playMessage = "Stop";
    } else {
      this.audio.pause();
      this.audio.currentTime = 0;
      this.playMessage = "Play";
    }
    this.isPlaying = !this.isPlaying;

  }

  ngOnDestroy() {
    this.sub.unsubscribe();
  }

}
