import type { IResponse } from "@appTypes/IResponse";
import type { IUser } from "@appTypes/IUser";
import {
  CHANGE_AVATAR_ERROR_MESSAGES,
  CHANGE_INFORMATION_ERROR_MESSAGES,
  EXIT_FROM_ACCOUNT_ERROR_MESSAGES,
} from "@constants/apiErrors";
import { USERS_ENDPOINT } from "@constants/endpoints";
import { catchError } from "@functionals/catchError";
import axios from "axios";

export class UsersApi {
  public static async exitFromAccount(): Promise<IResponse<null>> {
    try {
      const response = await axios.post<IResponse<null>>(
        USERS_ENDPOINT + "/exit",
        {},
        { withCredentials: true },
      );

      return response.data;
    } catch (error) {
      return catchError(error, EXIT_FROM_ACCOUNT_ERROR_MESSAGES);
    }
  }

  public static async changeAvatar(
    avatar: File,
  ): Promise<IResponse<{ avatar_link: IUser["avatar_link"] } | null>> {
    try {
      const formData = new FormData();

      formData.append("avatar", avatar);

      const response = await axios.post<
        IResponse<{ avatar_link: IUser["avatar_link"] }>
      >(USERS_ENDPOINT + "/avatarlink", formData, {
        withCredentials: true,
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      return response.data;
    } catch (error) {
      return catchError(error, CHANGE_AVATAR_ERROR_MESSAGES);
    }
  }

  public static async changeInformation(
    username: string,
    info: string,
    theme: string,
  ): Promise<IResponse<null>> {
    try {
      const response = await axios.put<IResponse<null>>(
        USERS_ENDPOINT,
        { username, info, theme },
        { withCredentials: true },
      );

      return response.data;
    } catch (error) {
      return catchError(error, CHANGE_INFORMATION_ERROR_MESSAGES);
    }
  }
}
