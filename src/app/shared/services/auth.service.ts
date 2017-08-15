import {Injectable} from '@angular/core';
import {BehaviorSubject} from "rxjs/BehaviorSubject";
import {ReplaySubject} from "rxjs/ReplaySubject";
import {Observable} from "rxjs/Observable";

import {ApiService} from './api.service';

import {User} from "../models/user";


@Injectable()
export class AuthService {
  private authUrl: string;

  private currentUserSubject: BehaviorSubject<User>;
  public currentUser: Observable<User>;

  private isAuthenticatedSubject: ReplaySubject<boolean>;
  public isAuthenticated: Observable<boolean>;


  constructor(private apiService: ApiService) {
    this.authUrl = '/auth';

    this.currentUserSubject = new BehaviorSubject<User>(new User());
    this.currentUser = this.currentUserSubject.asObservable().distinctUntilChanged();

    this.isAuthenticatedSubject = new ReplaySubject<boolean>(1);
    this.isAuthenticated = this.isAuthenticatedSubject.asObservable();
  }


  populate() {
    this.apiService.get(`${this.authUrl}/verify`)
      .subscribe((res: any) => {
        this.setAuth(this.getValidUserFromJson(res.data.user));
      }, err => {
        this.purgeAuth();
      })
  }

  register(body: Object = {}): Observable<any> {
    return this.apiService.post(`${this.authUrl}/register`, body)
      .map((res: any) => {
        this.setAuth(this.getValidUserFromJson(res.data.user));
      })
      .catch(err => {
        this.purgeAuth();
        return Observable.throw(err);
      })
  }

  login(body: Object = {}): Observable<any> {
    return this.apiService.post(`${this.authUrl}/login`, body)
      .map((res: any) => {
        this.setAuth(this.getValidUserFromJson(res.data.user));
      })
      .catch(err => {
        this.purgeAuth();
        return Observable.throw(err);
      })
  }

  logout() {
    this.apiService.get(`${this.authUrl}/logout`)
      .subscribe((res: any) => {
        this.purgeAuth();
      }, (err: any) => {
        this.purgeAuth();
      });
  }

  setAuth(user: User) {
    // Set current user data into observable
    this.currentUserSubject.next(user);
    // Set isAuthenticated to true
    this.isAuthenticatedSubject.next(true);
  }

  purgeAuth() {
    // Set current user to an empty object
    this.currentUserSubject.next(new User());
    // Set auth status to false
    this.isAuthenticatedSubject.next(false);
  }

  getValidUserFromJson(userJsonData): User {
    let user = new User();
    (<any>Object).assign(user, userJsonData);
    return user;
  }

}
