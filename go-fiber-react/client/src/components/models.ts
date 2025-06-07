export interface Todo {
  id: number;
  content: string;
}

export interface Meta {
  totalCount: number;
}

export interface Skills {
  id: number;
  language: string;
  level: string;
}

export interface CarouselInterface {
  id: number;
  text: string;
  iconName: string;
}

export interface UserResponse {
  data: Users[];
  success: boolean;
  message: string;
}
export interface Users {
  Id: number;
  Username: string;
  Email: string;
  RoleId: number;
  Role: Roles;
}

export interface Roles {
  Id: number;
  Name: string;
  CreatedAt: string;
}

export interface userRegistrationPayload {
  username: string;
  email: string;
  password: string;
}

export interface UserRegistrationRespsonse {
  success: boolean;
  message: string;
  data: Users;
}

export interface UserListRows {
  no: number;
  Username: string;
  Email: string;
}
