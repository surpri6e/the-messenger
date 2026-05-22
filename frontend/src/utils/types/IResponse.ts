export type TResponseStatus = 200 | 400 | 404 | 429 | 430 | 500 | 700;

export interface IResponse<T> {
  status: TResponseStatus;
  message: string;
  body: null | T;
}
