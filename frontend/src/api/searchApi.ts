import type { IResponse } from "@appTypes/IResponse";
import type { ISmallUser } from "@appTypes/IUser";
import { SEARCH_USER_ERROR_MESSAGES } from "@constants/apiErrors";
import { SEARCH_ENDPOINT } from "@constants/endpoints";
import { catchError } from "@functionals/catchError";
import axios from "axios";

export class SearchApi {
  public static async searchUser(
    username: string,
  ): Promise<IResponse<ISmallUser[] | null>> {
    try {
      const response = await axios.get<IResponse<ISmallUser[]>>(
        SEARCH_ENDPOINT + `/${username}`,
        { withCredentials: true },
      );

      return response.data;
    } catch (error) {
      return catchError(error, SEARCH_USER_ERROR_MESSAGES);
    }
  }
}
