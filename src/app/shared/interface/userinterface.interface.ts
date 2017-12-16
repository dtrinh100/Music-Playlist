// The UserInterface is an interface used specifcally only for forms
export interface UserInterface {
  username: string;
  email: string;
  account: {
    password: string;
    confirm: string;
  }
}
