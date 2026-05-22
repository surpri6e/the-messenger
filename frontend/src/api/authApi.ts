import type { IResponse } from "@appTypes/IResponse";
import axios from "axios";
import { AUTH_ENDPOINT } from "@constants/endpoints";
import type { IUser } from "@appTypes/IUser";
import { catchError } from "@functionals/catchError";
import {
  GET_USER_DATA_ERROR_MESSAGES,
  SET_AUTH_TOKEN_ERROR_MESSAGES,
} from "../utils/constants/apiErrors";

export class AuthApi {
  public static async setAuthToken(
    email: string,
    password: string,
  ): Promise<IResponse<null>> {
    try {
      const response = await axios.post<IResponse<null>>(
        AUTH_ENDPOINT,
        { email, password },
        { withCredentials: true },
      );

      return response.data;
    } catch (error) {
      return catchError(error, SET_AUTH_TOKEN_ERROR_MESSAGES);
    }
  }

  public static async getUserData(): Promise<IResponse<IUser | null>> {
    try {
      const response = await axios.get<IResponse<IUser>>(AUTH_ENDPOINT, {
        withCredentials: true,
      });

      return response.data;
    } catch (error) {
      return catchError(error, GET_USER_DATA_ERROR_MESSAGES);
    }
  }
}
