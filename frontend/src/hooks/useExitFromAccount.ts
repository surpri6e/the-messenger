import { useAlertsStore } from "@stores/useAlertsStore";
import { UsersApi } from "@api/usersApi";

interface IUseExitFromAccount {
  fn: () => Promise<boolean>;
}

export const useExitFromAccount = (): IUseExitFromAccount => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  async function fn(): Promise<boolean> {
    const exitFromAccountResponse = await UsersApi.exitFromAccount();

    if (exitFromAccountResponse.status !== 200) {
      alertsActions.addErrorAlert(exitFromAccountResponse.message);
      return false;
    }

    return true;
  }

  return { fn };
};
