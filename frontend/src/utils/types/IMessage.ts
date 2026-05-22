export interface IMessage {
  id: number;
  user_id: number;
  where_id: number;
  text: string;
  status: string;
  is_pinned: boolean;
  created_at: string;
  is_changed: boolean;
  is_forwarded: boolean;
  type: string;
  file_link: string;
}
