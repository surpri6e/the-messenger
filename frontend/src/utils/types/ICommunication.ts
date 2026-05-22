export interface IChat {
  id: number;
  first_person_id: number;
  second_person_id: number;
  created_at: string;
}

export interface IGroup {
  id: number;
  owner_id: number;
  name: string;
  info: string;
  users_id: number[];
  admins_id: number[];
  enemies_id: number[];
  created_at: string;
}

export interface ICommunication {
  type: "chat" | "group" | null;
  currentChat: IChat | null;
  currentGroup: IGroup | null;
}
