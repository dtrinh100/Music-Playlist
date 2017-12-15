import { Injectable } from '@angular/core';

import { Observable } from 'rxjs/Observable';

import { ApiService } from './api.service';
import { User } from '../models'

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Injectable()
export class UserService {
  private usersUrl = '/auth/users';

  constructor(private apiService: ApiService) {
  }

  getUsers(): Observable<User[]> {
    return this.apiService.get(`${this.usersUrl}`)
      .map((res: any) => {
        return res.users as User[];
      });
  }

  getUser(userName: string): Observable<User> {
    return this.apiService.get(`${this.usersUrl}/${userName}`)
      .map((res: any) => res.user as User);
  }

  deleteUser(userName: string): Observable<void> {
    return this.apiService.delete(`${this.usersUrl}/${userName}`);
  }

  updateUser(user: User): Observable<User> {
    return this.apiService.put(`${this.usersUrl}/${user.username}`, user)
      .map((res: any) => res.user as User);
  }

}
