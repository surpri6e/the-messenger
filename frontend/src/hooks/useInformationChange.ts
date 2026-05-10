import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";
import { useErrorDelay } from "./useErrorDelay";
import { useUserStore } from "@stores/useUserStore";
import { UsersApi } from "@api/usersApi";
import type { IUser } from "@appTypes/IUser";

interface IUseInformationChange {
  usernameError: boolean;
  informationError: boolean;

  isLoading: boolean;

  fn: (username: string, information: string) => Promise<boolean>;
}

export const useInformationChange = (): IUseInformationChange => {
  const { actions: alertsActions } = useAlertsStore((state) => state);
  const { user, actions: userActions } = useUserStore((state) => state);

  const [usernameError, setUsernameError] = useErrorDelay(1000);
  const [informationError, setInformationError] = useErrorDelay(1000);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(username: string, information: string): Promise<boolean> {
    let resultErrorMessage = "";

    if (username.length === 0) {
      resultErrorMessage += "Ваш юзернейм не может быть пустым! ";
      setUsernameError(true);
    }
    if (information.length >= 250) {
      resultErrorMessage +=
        "Информация о вас не должна превышать 250 символов! ";
      setInformationError(true);
    }

    if (resultErrorMessage.length !== 0) {
      resultErrorMessage = resultErrorMessage.slice(0, -1);

      alertsActions.addErrorAlert(resultErrorMessage);
      return false;
    }

    setIsLoading(true);

    const changeInformationResponse = await UsersApi.changeInformation(
      username,
      information,
      user!.theme,
    );

    if (changeInformationResponse.status === 200) {
      userActions.setUser({ ...user, username, info: information } as IUser);
      alertsActions.addSuccessAlert(changeInformationResponse.message);
    } else {
      alertsActions.addErrorAlert(changeInformationResponse.message);
      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { fn, isLoading, informationError, usernameError };
};
