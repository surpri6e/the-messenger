import type { IResponse } from "@appTypes/IResponse";

import { SEND_MESSAGE_ERROR_MESSAGES } from "@constants/apiErrors";
import { MESSAGES_ENDPOINT } from "@constants/endpoints";
import { catchError } from "@functionals/catchError";
import axios from "axios";

export class MessagesApi {
  public static async sendMessage(
    message: string,
    where_id: number,
  ): Promise<IResponse<null>> {
    try {
      const response = await axios.post<IResponse<null>>(
        MESSAGES_ENDPOINT,
        {
          text: message,
          status: "sended",
          type: "text",
          where_id,
          file_link: "",
        },
        { withCredentials: true },
      );

      return response.data;
    } catch (error) {
      return catchError(error, SEND_MESSAGE_ERROR_MESSAGES);
    }
  }
}
