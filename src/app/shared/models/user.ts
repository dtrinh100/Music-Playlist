// The User type is used as a model for REST API calls, not to be confused with the type UserInterface
export class User {
  id: number;
  username: string;
  first_name: string;
  last_name: string;
  pic_url: string;
  contributions: string[];
  my_playlist: string[];
  email: string;
  password: string;
  confirm: string;

}
