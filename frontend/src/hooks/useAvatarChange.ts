import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";
import { useErrorDelay } from "./useErrorDelay";
import { formatBytesToBytes } from "bytes-transform";
import { UsersApi } from "@api/usersApi";
import { useUserStore } from "@stores/useUserStore";
import type { IUser } from "@appTypes/IUser";

interface IUseAvatarChange {
  avatarError: boolean;

  isLoading: boolean;

  fn: (file: File) => Promise<boolean>;
}

export const useAvatarChange = (): IUseAvatarChange => {
  const { actions: alertsActions } = useAlertsStore((state) => state);
  const { user, actions: userActions } = useUserStore((state) => state);

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

    const changeAvatarResponse = await UsersApi.changeAvatar(file);

    if (changeAvatarResponse.status === 200) {
      userActions.setUser({ ...user, ...changeAvatarResponse.body } as IUser);
      alertsActions.addSuccessAlert(changeAvatarResponse.message);
    } else {
      alertsActions.addErrorAlert(changeAvatarResponse.message);
      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { fn, isLoading, avatarError };
};
