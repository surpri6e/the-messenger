import { AuthApi } from "@api/authApi";
import { useAlertsStore } from "@stores/useAlertsStore";
import { useUserStore } from "@stores/useUserStore";
import { useState } from "react";

interface IUseGetUserData {
  fn: (isWithoutAlerts?: boolean) => Promise<boolean>;
  isLoading: boolean;
}

export const useGetUserData = (): IUseGetUserData => {
  const { actions: userActions } = useUserStore((state) => state);
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(isWithoutAlerts?: boolean): Promise<boolean> {
    setIsLoading(true);

    const getUserDataResponse = await AuthApi.getUserData();

    if (getUserDataResponse.status === 200) {
      if (!isWithoutAlerts) {
        alertsActions.addSuccessAlert(getUserDataResponse.message);
      }

      userActions.setUser(getUserDataResponse.body!);
    } else {
      if (!isWithoutAlerts) {
        alertsActions.addErrorAlert(getUserDataResponse.message);
      }

      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { fn, isLoading };
};
