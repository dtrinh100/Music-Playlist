import {Injectable} from '@angular/core';
import {Headers, Http, Response, URLSearchParams} from '@angular/http';
import {Observable} from 'rxjs/Rx';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';


@Injectable()
export class ApiService {
  private apiUrl = '/api';

  constructor(private http: Http) {
  }

  private static setHeaders(): Headers {
    let headersConfig = {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    };

    return new Headers(headersConfig);
  }

  private static formatErrors(error: any) {
    return Observable.throw(error);
  }

  get(path: string, params: URLSearchParams = new URLSearchParams()): Observable<any> {
    return this.http.get(`${this.apiUrl}${path}`, {headers: ApiService.setHeaders(), search: params})
      .catch(ApiService.formatErrors)
      .map((res: Response) => res.json());
  }

  post(path: string, body: Object = {}): Observable<any> {
    return this.http.post(`${this.apiUrl}${path}`, JSON.stringify(body), {headers: ApiService.setHeaders()})
      .catch(ApiService.formatErrors)
      .map((res: Response) => res.json());
  }

  put(path: string, body: Object = {}): Observable<any> {
    return this.http.put(`${this.apiUrl}${path}`, JSON.stringify(body), {headers: ApiService.setHeaders()})
      .catch(ApiService.formatErrors)
      .map((res: Response) => res.json());
  }

  delete(path): Observable<any> {
    return this.http.delete(`${this.apiUrl}${path}`, {headers: ApiService.setHeaders()})
      .catch(ApiService.formatErrors)
      .map((res: Response) => res.json());
  }

}
