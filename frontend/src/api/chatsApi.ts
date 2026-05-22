import type { IChat } from "@appTypes/IChat";
import type { IResponse } from "@appTypes/IResponse";
import { CREATE_CHAT_ERROR_MESSAGES } from "@constants/apiErrors";
import { CHATS_ENDPOINT } from "@constants/endpoints";
import { catchError } from "@functionals/catchError";
import axios from "axios";

export class ChatsApi {
  public static async createChat(
    first_person_id: number,
    second_person_id: number,
  ): Promise<IResponse<IChat | null>> {
    try {
      const response = await axios.post<IResponse<IChat>>(
        CHATS_ENDPOINT,
        { first_person_id, second_person_id },
        { withCredentials: true },
      );

      return response.data;
    } catch (error) {
      return catchError(error, CREATE_CHAT_ERROR_MESSAGES);
    }
  }
}
