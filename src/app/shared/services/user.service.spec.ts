import { TestBed, getTestBed, async, inject } from '@angular/core/testing';
import {
  HttpModule,
  Http,
  Response,
  ResponseOptions,
  XHRBackend,
  BaseRequestOptions,
  RequestMethod
} from '@angular/http';
import { MockBackend, MockConnection } from '@angular/http/testing';
import { UserService } from './user.service';
import { ApiService } from './api.service';
import { User } from '../models/user'

describe('UserService', () => {
  let mockBackend: MockBackend;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ HttpModule ],
      providers: [
        UserService,
        ApiService,
        MockBackend,
        BaseRequestOptions,
        {
          provide: Http,
          deps: [ MockBackend, BaseRequestOptions ],
          useFactory:
            (backend: XHRBackend, defaultOptions: BaseRequestOptions) => {
              return new Http(backend, defaultOptions);
            }
        }
      ]
    });
    mockBackend = getTestBed().get(MockBackend);
  });

  it('should get users', (done: DoneFn) => {
    let userService: UserService;
    const mockResponse = {
      users: [
        {
          id: 1,
          username: 'dtrinh100',
          firstname: 'David',
          lastname: 'Trinh',
          picurl: 'www.example.com/pic',
          contributions: [],
          my_playlist: []
        },
        {
          id: 2,
          username: 'hlovo',
          firstname: 'Hector',
          lastname: 'Lovo',
          picurl: 'www.example.com/pic',
          contributions: [],
          my_playlist: []
        },
      ]
    };

    getTestBed().compileComponents().then(() => {
      mockBackend.connections.subscribe(
        (connection: MockConnection) => {
          connection.mockRespond(new Response(
            new ResponseOptions({
                body: JSON.stringify(mockResponse)
              }
            )));
        });

      userService = getTestBed().get(UserService);
      userService.getUsers().subscribe((users) => {
        expect(users.length).toBe(2);
        expect(users[ 0 ].username).toEqual('dtrinh100');
        expect(users[ 1 ].username).toEqual('hlovo');

        expect(users[ 0 ][ 'firstname' ]).toEqual('David');
        expect(users[ 1 ][ 'firstname' ]).toEqual('Hector');

        expect(users[ 0 ][ 'lastname' ]).toEqual('Trinh');
        expect(users[ 1 ][ 'lastname' ]).toEqual('Lovo');

        expect(users[ 0 ][ 'picurl' ]).toEqual('www.example.com/pic');
        expect(users[ 1 ][ 'picurl' ]).toEqual('www.example.com/pic');

        expect(users[ 0 ][ 'contributions' ].length).toEqual(0);
        expect(users[ 1 ][ 'contributions' ].length).toEqual(0);

        expect(users[ 0 ][ 'my_playlist' ].length).toEqual(0);
        expect(users[ 1 ][ 'my_playlist' ].length).toEqual(0);
        done();
      });
    });
  });

  it('should fetch a single user by a username key', done => {
    let userService: UserService;
    const mockResponse = {
      user: {
        id: 1,
        username: 'dtrinh100',
        firstname: 'David',
        lastname: 'Trinh',
        picurl: 'www.example.com/pic',
        contributions: [],
        my_playlist: []
      }
    };

    getTestBed().compileComponents().then(() => {
      mockBackend.connections.subscribe(
        (connection: MockConnection) => {
          expect(connection.request.url).toBe('/api/auth/users/dtrinh100');
          connection.mockRespond(new Response(
            new ResponseOptions({
                body: JSON.stringify(mockResponse)
              }
            )));
        });

      userService = getTestBed().get(UserService);
      userService.getUser('dtrinh100').subscribe(function (user) {
        expect(user.id).toBe(1);
        done();
      });

    });

  });

  it('should update the user information', async(inject([ UserService ], (userService) => {
    const mockResponse = {
      user: {
        id: 1,
        username: 'dtrinh100',
        firstname: 'David',
        lastname: 'Trinh',
        picurl: 'www.example.com/pic',
        contributions: [],
        my_playlist: []
      }
    };

    getTestBed().compileComponents().then(() => {
      mockBackend.connections.subscribe((connection: MockConnection) => {
        expect(connection.request.method).toBe(RequestMethod.Put);
        connection.mockRespond(new Response(new ResponseOptions({body: mockResponse})));
      });

      const userObj = {
        id: 1,
        username: 'dtrinh100',
        firstname: 'David',
        lastname: 'Trinh',
        picurl: 'www.example.com/pic',
        contributions: [],
        my_playlist: []
      };

      userService.updateUser(userObj).subscribe(
        (data) => {
          expect(data).toEqual(userObj);
        });
    });
  })));

  it('should delete an existing user', async(inject([ UserService ], (userService) => {
    getTestBed().compileComponents().then(() => {
      mockBackend.connections.subscribe(connection => {
        expect(connection.request.method).toBe(RequestMethod.Delete);
        connection.mockRespond(new Response(new ResponseOptions({body: {}, status: 204})));
      });

      userService.deleteUser('dtrinh100').subscribe(
        (res: Response) => {
          expect(res.status).toEqual(204);
        });
    });
  })));

});
