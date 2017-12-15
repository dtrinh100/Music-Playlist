import {Injectable} from '@angular/core';
import {Headers, Http, URLSearchParams} from '@angular/http';
import {Observable} from 'rxjs/Rx';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';


@Injectable()
export class ApiService {
  private apiUrl = '/api';

  private static formatErrors(error: any) {
    return Observable.throw(error);
  }

  private static setHeaders(): Headers {
    let headersConfig = {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    };

    return new Headers(headersConfig);
  }

  constructor(private http: Http) {
  }

  get(path: string, params: URLSearchParams = new URLSearchParams()): Observable<any> {
    return this.http.get(`${this.apiUrl}${path}`, {headers: ApiService.setHeaders(), search: params})
      .catch(ApiService.formatErrors)
  }

  post(path: string, body: Object = {}): Observable<any> {
    return this.http.post(`${this.apiUrl}${path}`, body, {headers: ApiService.setHeaders()})
      .catch(ApiService.formatErrors)
  }

  put(path: string, body: Object = {}): Observable<any> {
    return this.http.put(`${this.apiUrl}${path}`, body, {headers: ApiService.setHeaders()})
      .catch(ApiService.formatErrors)
  }

  delete(path): Observable<any> {
    return this.http.delete(`${this.apiUrl}${path}`, {headers: ApiService.setHeaders()})
      .catch(ApiService.formatErrors)
  }

  patch(path: string, body: Object = {}): Observable<any> {
    return this.http.patch(`${this.apiUrl}${path}`, body, {headers: ApiService.setHeaders()})
  }

}
