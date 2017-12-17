import { Injectable }    from '@angular/core';

import { Observable } from "rxjs/Observable";

import { ApiService } from './api.service';
import { User } from '../models/song'

import 'rxjs/add/operator/map';

@Injectable()
export class SongService {
  private songsUrl = '/songs';

  constructor(private apiService: ApiService) {
  }

  getSongs(): Observable<Song[]> {
    return this.apiService.get(`${this.songsUrl}`)
      .map((res: any) => res.data as Song[]);
  }

  getSong(id: number): Observable<Song> {
    return this.apiService.get(`${this.songsUrl}/${id}`)
      .map((res: any) => res.data as Song);
  }


}
