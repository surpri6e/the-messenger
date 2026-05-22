import type { IResponse } from "@appTypes/IResponse";
import axios from "axios";
import { REGISTRATION_ENDPOINT } from "@constants/endpoints";
import { catchError } from "@functionals/catchError";
import { REGISTER_USER_ERROR_MESSAGES } from "@constants/apiErrors";

export class RegistrationApi {
  public static async registerUser(
    email: string,
    password: string,
    username: string,
  ): Promise<IResponse<null>> {
    try {
      const response = await axios.post<IResponse<null>>(
        REGISTRATION_ENDPOINT,
        { email, password, username },
        { withCredentials: true },
      );

      return response.data;
    } catch (error) {
      return catchError(error, REGISTER_USER_ERROR_MESSAGES);
    }
  }
}
