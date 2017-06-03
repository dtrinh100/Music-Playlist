export interface UserInterface {
  username: string;
  email: string;
  account: {
    password: string;
    confirm: string;
  }
}
