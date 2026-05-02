import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";

interface IUseThemeChange {
  isLoading: boolean;

  fn: (themeName: string) => Promise<boolean>;
}

export const useThemeChange = (): IUseThemeChange => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(themeName: string): Promise<boolean> {
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

  return { fn, isLoading };
};
