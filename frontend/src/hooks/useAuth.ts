import { isEmailCorrectly } from "@functionals/isEmailCorrectly";
import { useErrorDelay } from "./useErrorDelay";
import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";
import { AuthApi } from "@api/authApi";

interface IUseAuth {
  emailError: boolean;
  passwordError: boolean;

  isLoading: boolean;

  fn: (email: string, password: string) => Promise<boolean>;
}

export const useAuth = (): IUseAuth => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [emailError, setEmailError] = useErrorDelay(1000);
  const [passwordError, setPasswordError] = useErrorDelay(1000);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(email: string, password: string): Promise<boolean> {
    let resultErrorMessage = "";

    if (!isEmailCorrectly(email)) {
      resultErrorMessage = "Почта введена неверно! ";
      setEmailError(true);
    }

    if (!(password.length >= 5 && password.length <= 20)) {
      resultErrorMessage += "Пароль должен быть от 5 до 20 символов! ";
      setPasswordError(true);
    }

    if (resultErrorMessage.length !== 0) {
      resultErrorMessage = resultErrorMessage.slice(0, -1);

      alertsActions.addErrorAlert(resultErrorMessage);
      return false;
    }

    setIsLoading(true);

    const setAuthResponse = await AuthApi.setAuthToken(email, password);

    if (setAuthResponse.status !== 200) {
      alertsActions.addErrorAlert(setAuthResponse.message);
      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { emailError, passwordError, fn, isLoading };
};
