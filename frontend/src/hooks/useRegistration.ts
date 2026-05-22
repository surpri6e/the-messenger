import { useAlertsStore } from "@stores/useAlertsStore";
import { useErrorDelay } from "./useErrorDelay";
import { isEmailCorrectly } from "@functionals/isEmailCorrectly";
import { useState } from "react";
import { RegistrationApi } from "@api/registrationApi";

interface IUseRegistration {
  emailError: boolean;
  passwordError: boolean;
  usernameError: boolean;

  isLoading: boolean;

  fn: (email: string, password: string, username: string) => Promise<boolean>;
}

export const useRegistration = (): IUseRegistration => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [emailError, setEmailError] = useErrorDelay(1000);
  const [passwordError, setPasswordError] = useErrorDelay(1000);
  const [usernameError, setUsernameError] = useErrorDelay(1000);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(
    email: string,
    password: string,
    username: string,
  ): Promise<boolean> {
    let resultErrorMessage = "";

    if (!isEmailCorrectly(email)) {
      resultErrorMessage = "Почта введена неверно! ";
      setEmailError(true);
    }

    if (!(password.length >= 5 && password.length <= 20)) {
      resultErrorMessage += "Пароль должен быть от 5 до 20 символов! ";
      setPasswordError(true);
    }

    if (username.length === 0) {
      resultErrorMessage += "Ваш юзернейм не может быть пустым! ";
      setUsernameError(true);
    }

    if (resultErrorMessage.length !== 0) {
      resultErrorMessage = resultErrorMessage.slice(0, -1);

      alertsActions.addErrorAlert(resultErrorMessage);
      return false;
    }

    setIsLoading(true);

    const registerUserResponse = await RegistrationApi.registerUser(
      email,
      password,
      username,
    );

    if (registerUserResponse.status === 200) {
      alertsActions.addSuccessAlert(registerUserResponse.message);
    } else {
      alertsActions.addErrorAlert(registerUserResponse.message);
      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { emailError, passwordError, usernameError, isLoading, fn };
};
