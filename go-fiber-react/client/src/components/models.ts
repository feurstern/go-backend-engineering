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

export interface UseResponse {
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
