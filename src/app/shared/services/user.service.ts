import { Injectable }    from '@angular/core';
import { Headers, Http } from '@angular/http';

import { Observable } from "rxjs/Observable";

import { ApiService } from './api.service';
import { User } from '../models/user'

import 'rxjs/add/operator/map';

@Injectable()
export class UserService {
  private usersUrl = '/users';

  constructor(private apiService: ApiService) {
  }

  // TODO: Take special note: All of the following funcs have been adapted to use Observable

  getUsers(): Observable<User[]> {
    return this.apiService.get(`${this.usersUrl}`)
      .map((res: any) => res.data as User[]);
  }

  getUser(id: number): Observable<User> {
    return this.apiService.get(`${this.usersUrl}/${id}`)
      .map((res: any) => res.data as User);
  }

  createUser(username: string, email: string, password: string, confirm: string): Observable<Response> {
    return this.apiService.post(
      this.usersUrl,
      JSON.stringify({
        username: username,
        email: email,
        password: password,
        confirm: confirm
      })
    ).map((res: any) => res.data as Response);
  }

  deleteUser(id: number): Observable<void> {
    return this.apiService.delete(`${this.usersUrl}\${id}`)
      .map(_ => null);
  }

  updateUser(user: User): Observable<Response> {
    return this.apiService.put(`${this.usersUrl}/${user.id}`, JSON.stringify(user))
      .map((res: any) => res.data as Response);
  }

 //  private handleError(error: any): Promise<any> {
 //    console.error('An error occurred', error); // we won't be really handling errors in this app
 //    return Promise.reject(error.message || error);
 //  }
}
