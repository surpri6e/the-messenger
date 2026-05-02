import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";
import { useErrorDelay } from "./useErrorDelay";
import { formatBytesToBytes } from "bytes-transform";

interface IUseAvatarChange {
  avatarError: boolean;

  isLoading: boolean;

  fn: (file: File) => Promise<boolean>;
}

export const useAvatarChange = (): IUseAvatarChange => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [avatarError, setAvatarError] = useErrorDelay(1000);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(file: File): Promise<boolean> {
    let resultErrorMessage = "";

    if (
      file.type !== "image/png" &&
      file.type !== "image/jpeg" &&
      file.type !== "image/jpg"
    ) {
      resultErrorMessage = "Неверный тип аватарки! ";
      setAvatarError(true);
    }

    if (formatBytesToBytes(4, "MB") < file.size) {
      resultErrorMessage += "Файл не должен превышать 4 МБ! ";
      setAvatarError(true);
    }

    if (resultErrorMessage.length !== 0) {
      resultErrorMessage = resultErrorMessage.slice(0, -1);

      alertsActions.addErrorAlert(resultErrorMessage);
      return false;
    }

    setIsLoading(true);

    // const setAuthResponse = await AuthApi.setAuthToken(email, password);

    // if (setAuthResponse.status !== 200) {
    //   alertsActions.addErrorAlert(setAuthResponse.message);
    //   setIsLoading(false);
    //   return false;
    // }

    const delay = (ms: number) =>
      new Promise((resolve) => setTimeout(resolve, ms));

    await delay(2000);

    setIsLoading(false);

    return true;
  }

  return { fn, isLoading, avatarError };
};
