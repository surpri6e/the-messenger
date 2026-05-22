import { AxiosError } from "axios";
import { type IAllErrorMessages } from "@constants/apiErrors";
import type { IResponse } from "@appTypes/IResponse";

export const catchError = (
  error: unknown,
  errorMessages: IAllErrorMessages,
): IResponse<null> => {
  if (error instanceof AxiosError) {
    if (error.response) {
      if (errorMessages[400] && error.response.status === 400)
        return { body: null, status: 400, message: errorMessages[400] };

      if (errorMessages[404] && error.response.status === 404)
        return { body: null, status: 404, message: errorMessages[404] };

      if (errorMessages[429] && error.response.status === 429)
        return { body: null, status: 429, message: errorMessages[429] };

      if (errorMessages[430] && error.response.status === 430)
        return { body: null, status: 430, message: errorMessages[430] };

      if (errorMessages[500] && error.response.status === 500)
        return { body: null, status: 500, message: errorMessages[500] };
    }
  }

  return { body: null, status: 700, message: "Неизвестная ошибка" };
};
