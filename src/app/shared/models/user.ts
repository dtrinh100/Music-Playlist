// The User type is used as a model for REST API calls, not to be confused with the type UserInterface
export class User {
  id: number;
  username: string;
  firstname: string;
  lastname: string;
  picurl: string;
  contributions: string[];
  playlist: string[];
  email: string;
  password: string;
  confirm: string;
}
