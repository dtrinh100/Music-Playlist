import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-songs',
  templateUrl: './songs.component.html',
  styleUrls: ['./songs.component.scss']
})
export class SongsComponent implements OnInit {
  private images;

  constructor() { }

  ngOnInit() {
    this.images = [
      {
        "id": 1,
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "id": 2,
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "id": 3,
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "id": 4,
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "id": 5,
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
      {
        "id": 6,
        "url": "../../assets/img/album_art.png",
        "alt": "Album Art Picture"
      },
    ];
  }

}
