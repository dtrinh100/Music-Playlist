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
        this.audio = new Audio(data.musicURL);
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
    /*  this.audio = new Audio("../../assets/audio/audio1.mp3");
      this.song = {
        "id": 1,
        "name": "Going Higher",
        "imgURL": "../../assets/img/album_art.png",
        "alt": "Album Art Picture",
        "description": `Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium.
  Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, `,
        "credit": "https://www.bensound.com/royalty-free-music"
      }; */
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
