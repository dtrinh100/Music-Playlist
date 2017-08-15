import {TestBed, inject, async, fakeAsync} from '@angular/core/testing';
import {
  BaseRequestOptions, Http, HttpModule, RequestMethod,
  Response, ResponseOptions, ResponseType
} from '@angular/http';
import {MockBackend} from '@angular/http/testing';

import {AuthService} from './auth.service';
import {ApiService} from './api.service';
import {Subscription} from "rxjs/Subscription";
import {User} from "../models/user";


describe('AuthService', () => {
  let userResponse_valid = {
    data: {
      user: {
        id: 1,
        username: "user_one",
        first_name: "user",
        last_name: "one",
        pic_url: "",
        contributions: [""],
        my_playlist: [""],
        email: "user@email.com",
        password: "",
        confirm: ""
      }
    }
  };

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpModule],
      providers: [
        ApiService,
        AuthService,
        BaseRequestOptions,
        MockBackend,
        {
          provide: Http,
          useFactory: (backend, options) => new Http(backend, options),
          deps: [MockBackend, BaseRequestOptions]
        }
      ]
    });
  });

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=- Constructors =-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  fit('should be created', inject([AuthService], (service: AuthService) => {
    expect(service).toBeTruthy();
  }));

  fit('should construct', async(inject([AuthService, MockBackend], (service, _) => {
    expect(service).toBeDefined();
  })));

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-= Helper Functions -=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  // This function helps test standard/error server-responses
  function mockBackendHelper(mockBackend, bodyStr = {}, statusNum = 200, responseType = ResponseType.Default): Subscription {
    let mockResponse = new Response(new ResponseOptions({
      body: JSON.stringify(bodyStr),
      status: statusNum,
      type: responseType
    }));

    return mockBackend.connections.subscribe(conn => {
      if (responseType === ResponseType.Default) {
        conn.mockRespond(mockResponse);
      } else if (responseType === ResponseType.Error) {
        conn.mockError(mockResponse);
      }
    });
  }

  // This function helps make tests easier to read
  function mockInjectAsync(expectTests: (service, backend) => void) {
    return inject([AuthService, MockBackend], fakeAsync((service: AuthService, backend: MockBackend) => {
      expectTests(service, backend);
    }));
  }

  // This function helps create a valid user from valid data: userResponse_valid
  function getValidUser(): User {
    let user = new User();
    (<any>Object).assign(user, userResponse_valid.data.user);
    return user;
  }

  // This function helps make user-validity easier to read
  function testUserValidity(service, validity) {
    service.currentUser.subscribe(userData => {
      expect(userData).toEqual(validity ? getValidUser() : new User());
    });

    service.isAuthenticated.subscribe(isAuthed => {
      validity ? expect(isAuthed).toBeTruthy() : expect(isAuthed).toBeFalsy();
    });
  }

  // This var-type should be used in params that need to be ignored
  let ignore = () => {};

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=- populate() =-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  fit('should try to make a GET request to /api/auth/verify', mockInjectAsync((service, backend) => {
      backend.connections.subscribe((connection) => {
        expect(connection.request.url).toBe('/api/auth/verify');
        expect(connection.request.method).toBe(RequestMethod.Get);
      });
      service.populate();
    })
  );

  fit('should have called setAuth', mockInjectAsync((service, backend) => {
      let setAuthSpy = spyOn(service, 'setAuth');

      mockBackendHelper(backend, userResponse_valid);
      service.populate();
      expect(setAuthSpy.calls.count()).toBe(1, 'setAuthSpy was not called');
    })
  );

  fit('should set authenticated user', mockInjectAsync((service, backend) => {
      mockBackendHelper(backend, userResponse_valid);
      service.populate();

      testUserValidity(service, true);
    })
  );

  fit('should have called purgeAuth', mockInjectAsync((service, backend) => {
      let purgeAuthSpy = spyOn(service, 'purgeAuth');
      mockBackendHelper(backend, {}, 404, ResponseType.Error);
      service.populate();
      expect(purgeAuthSpy.calls.count()).toBe(1, 'purgeAuth was not called');
    })
  );

  // NOTE: JWT is stored as a cookie so it's not demo'ed here
  fit('should de-authenticate an authenticated user with expired JWT', mockInjectAsync((service, backend) => {
      let mockConnectionSub = mockBackendHelper(backend, userResponse_valid);
      service.populate();

      // Ensure valid user
      let currentUserSub = service.currentUser.subscribe(userData => {
        expect(userData).toEqual(getValidUser());
      });
      // Ensure authentication
      let isAuthedSub = service.isAuthenticated.subscribe(isAuthed => {
        expect(isAuthed).toBeTruthy();
      });

      // Necessary to create new mockBackendHelper subscriptions
      [mockConnectionSub, currentUserSub, isAuthedSub].forEach(sub => {
        sub.unsubscribe();
      });

      // Attempt incorrect authentication
      mockBackendHelper(backend, {}, 404, ResponseType.Error);
      service.populate();

      // Ensure invalid user
      service.currentUser.subscribe(userData => {
        expect(userData).toEqual(new User());
      });
      // Ensure invalid authentication
      service.isAuthenticated.subscribe(isAuthed => {
        expect(isAuthed).toBeFalsy();
      });

    })
  );

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=- register() =-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  fit('should try to make a POST request to /api/auth/register', mockInjectAsync((service, backend) => {
      backend.connections.subscribe((connection) => {
        expect(connection.request.url).toBe('/api/auth/register');
        expect(connection.request.method).toBe(RequestMethod.Post);
      });
      service.register().subscribe();
    })
  );

  fit('should register & authenticate user', mockInjectAsync((service, backend) => {
      mockBackendHelper(backend, userResponse_valid);

      service.register().subscribe();
      testUserValidity(service, true);
    })
  );

  fit('should NOT register & NOT authenticate user upon server error', mockInjectAsync((service, backend) => {
      mockBackendHelper(backend, {}, 404, ResponseType.Error);

      service.register().subscribe(ignore, ignore);
      testUserValidity(service, false);
    })
  );

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-= login() =-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  fit('should try to make a POST request to /api/auth/login', mockInjectAsync((service, backend) => {
      backend.connections.subscribe((connection) => {
        expect(connection.request.url).toBe('/api/auth/login');
        expect(connection.request.method).toBe(RequestMethod.Post);
      });
      service.login().subscribe();
    })
  );

  fit('should login & authenticate user', mockInjectAsync((service, backend) => {
      mockBackendHelper(backend, userResponse_valid);

      service.login().subscribe();
      testUserValidity(service, true);
    })
  );

  fit('should NOT login & NOT authenticate user', mockInjectAsync((service, backend) => {
      mockBackendHelper(backend, {}, 404, ResponseType.Error);

      service.login().subscribe(ignore, ignore);
      testUserValidity(service, false);
    })
  );

  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=- logout() =-=-=-=-=-=-=-=-=-=-=-=
  // =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

  fit('should try to make a POST request to /api/auth/logout', mockInjectAsync((service, backend) => {
      backend.connections.subscribe((connection) => {
        expect(connection.request.url).toBe('/api/auth/logout');
        expect(connection.request.method).toBe(RequestMethod.Get);
      });
      service.logout();
    })
  );

  fit('should logout if successfully logged out of server', mockInjectAsync((service, backend) => {
      mockBackendHelper(backend, userResponse_valid);
      service.logout();

      testUserValidity(service, false);
    })
  );

  fit('should still logout if failed to logout of server', mockInjectAsync((service, backend) => {
      mockBackendHelper(backend, {}, 404, ResponseType.Error);
      service.login();

      testUserValidity(service, false);
    })
  );

});
