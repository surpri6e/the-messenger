import { UsersApi } from "@api/usersApi";
import type { IUser } from "@appTypes/IUser";
import { useAlertsStore } from "@stores/useAlertsStore";
import { useUserStore } from "@stores/useUserStore";
import { useState } from "react";

interface IUseThemeChange {
  isLoading: boolean;

  fn: (theme: string) => Promise<boolean>;
}

export const useThemeChange = (): IUseThemeChange => {
  const { actions: alertsActions } = useAlertsStore((state) => state);
  const { user, actions: userActions } = useUserStore((state) => state);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(theme: string): Promise<boolean> {
    setIsLoading(true);

    const changeInformationResponse = await UsersApi.changeInformation(
      user!.username,
      user!.info,
      theme,
    );

    if (changeInformationResponse.status === 200) {
      userActions.setUser({ ...user, theme } as IUser);
      alertsActions.addSuccessAlert(changeInformationResponse.message);
    } else {
      alertsActions.addErrorAlert(changeInformationResponse.message);
      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { fn, isLoading };
};
