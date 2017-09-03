import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-song',
  templateUrl: './song.component.html',
  styleUrls: ['./song.component.scss']
})
export class SongComponent implements OnInit {
  private images;

  constructor() { }

  ngOnInit() {
    this.images = [
      {
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
    ];
  }

}
