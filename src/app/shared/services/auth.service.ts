import {Injectable} from '@angular/core';
import {BehaviorSubject} from "rxjs/BehaviorSubject";
import {ReplaySubject} from "rxjs/ReplaySubject";
import {Observable} from "rxjs/Observable";

import {ApiService} from './api.service';
import {User} from "../models/user";


@Injectable()
export class AuthService {
  private currentUserSubject: BehaviorSubject<User>;
  public currentUser: Observable<User>;

  private isAuthenticatedSubject: ReplaySubject<boolean>;
  public isAuthenticated: Observable<boolean>;


  constructor(private apiService: ApiService) {
    this.currentUserSubject = new BehaviorSubject<User>(new User());
    this.currentUser = this.currentUserSubject.asObservable().distinctUntilChanged();

    this.isAuthenticatedSubject = new ReplaySubject<boolean>(1);
    this.isAuthenticated = this.isAuthenticatedSubject.asObservable();
  }

  populate() {
    this.apiService.get(`/auth/verify`)
      .subscribe((res: Response) => {
        const resBody: any = res.json();
        this.setAuth(this.getValidUserFromJson(resBody.user));
      }, _ => {
        this.purgeAuth();
      })
  }

  register(body: Object = {}): Observable<any> {
    return this.apiService.post(`/register`, body)
      .map((res: Response) => {
        const resBody: any = res.json();
        this.setAuth(this.getValidUserFromJson(resBody.user));
        return res;
      })
      .catch(err => {
        return Observable.throw(err);
      })
  }

  login(body: Object = {}): Observable<any> {
    return this.apiService.post(`/login`, body)
      .map((res: Response) => {
        const resBody: any = res.json();
        this.setAuth(this.getValidUserFromJson(resBody.user));
        return res;
      })
      .catch(err => {
        return Observable.throw(err);
      })
  }

  logout() {
    this.apiService.post(`/auth/logout`, {})
      .subscribe((res: Response) => {
        this.purgeAuth();
      }, _ => {
        // keep logged in until session expires
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
