export interface User {
  username: string;
  email: string;
  account: {
    password: string;
    confirm: string;
  }
}
