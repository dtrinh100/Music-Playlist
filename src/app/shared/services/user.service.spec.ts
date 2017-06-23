// TODO: Possibly come back to this test and see what is causing the second test to keep return undefined

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
import { User } from '../models/user'

describe('UserService', () => {
  let mockBackend: MockBackend;

  beforeEach(() => {

    TestBed.configureTestingModule({
      imports: [HttpModule],
      providers: [
        UserService,
        MockBackend,
        BaseRequestOptions,
        {
          provide: Http,
          deps: [MockBackend, BaseRequestOptions],
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
      data: [
        { id: 1, username: "dtrinh100", first_name: "David", last_name: "Trinh", pic_url: "www.example.com/pic", contributions: [], my_playlist: [] },
        { id: 2, username: "hlovo", first_name: "Hector", last_name: "Lovo", pic_url: "www.example.com/pic", contributions: [], my_playlist: [] },
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
      expect(userService).toBeDefined();




      userService.getUsers().then((users) => {
       expect(users.length).toBe(2);
        expect(users[0].username).toEqual('dtrinh100');
        expect(users[1].username).toEqual('hlovo');

        expect(users[0]["first_name"]).toEqual('David');
        expect(users[1]["first_name"]).toEqual('Hector');

        expect(users[0]["last_name"]).toEqual('Trinh');
        expect(users[1]["last_name"]).toEqual('Lovo');

        expect(users[0]["pic_url"]).toEqual('www.example.com/pic');
        expect(users[1]["pic_url"]).toEqual('www.example.com/pic');

        expect(users[0]["contributions"].length).toEqual(0);
        expect(users[1]["contributions"].length).toEqual(0);

        expect(users[0]["my_playlist"].length).toEqual(0);
        expect(users[1]["my_playlist"].length).toEqual(0);
        done();

      });

    });


  });


it('should fetch a single user by an id key', done => {
      let userService: UserService;
      const mockResponse = {

     id: 1, username: "dtrinh100", first_name: "David", last_name: "Trinh", pic_url: "www.example.com/pic", contributions: [], my_playlist: []


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

      userService.getUser(1).then(function(user) {
        expect(user.id).toBe(1);
        done();
      });

        });






  });


/*it('should insert a new user',
  async(inject([UserService], (userService) => {

     getTestBed().compileComponents().then(() => {

    mockBackend.connections.subscribe((connection: MockConnection) => {
      console.log(connection.request.method);
      expect(connection.request.method).toBe(RequestMethod.Post);

      connection.mockRespond(new Response(new ResponseOptions({status: 201})));
    });

  });


    userService.create("dtrinh100", "email@email.com", "password", "password").then(
      (data) => {
        console.log(data);
        expect(data).toBeDefined();
        expect(data.status).toBe(202);
      });
})));

it('should update the user information',
  async(inject([UserService], (userService) => {

     getTestBed().compileComponents().then(() => {
    mockBackend.connections.subscribe((connection: MockConnection) => {
      console.log(connection.request.method);
      expect(connection.request.method).toBe(RequestMethod.Put);

      connection.mockRespond(new Response(new ResponseOptions({status: 204})));
    });

  });

    const userObj = {
      username: "dtrinh100",
      email: "test@email.com",
      password: "password",
      confirm: "password"
    }

    userService.update(userObj).then(
      (data) => {
        console.log(data);
        expect(data).toBeDefined();
        expect(data.status).toBe(205);
      });
}))); */

});
