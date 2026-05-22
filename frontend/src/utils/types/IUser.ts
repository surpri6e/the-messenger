export interface IUser {
  id: number;
  email: string;
  password: string;
  username: string;
  theme: "purple1" | "purple2" | "purple3";
  info: string;
  avatar_link: string;
  created_at: string;
  is_admin: boolean;
  last_seen: string;
  is_email_accepted: boolean;
  is_muted_chats_id: number[];
  is_pinned_chats_id: number[];
}

export interface ISmallUser {
  id: number;
  username: string;
  info: string;
  avatar_link: string;
  created_at: string;
  is_online: boolean;
  is_email_accepted: boolean;
  last_seen: string;
}
