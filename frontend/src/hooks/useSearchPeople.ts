import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";
import { SearchApi } from "@api/searchApi";
import type { ISmallUser } from "@appTypes/IUser";

interface IUseSearchPeople {
  searchedUsers: ISmallUser[];

  isLoading: boolean;

  fn: (username: string) => Promise<boolean>;
}

export const useSearchPeople = (): IUseSearchPeople => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [searchedUsers, setSearchedUsers] = useState<ISmallUser[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  async function fn(username: string): Promise<boolean> {
    let resultErrorMessage = "";

    if (username.length === 0) {
      resultErrorMessage += "Поиск не может быть пустым! ";
    }

    if (resultErrorMessage.length !== 0) {
      resultErrorMessage = resultErrorMessage.slice(0, -1);

      alertsActions.addErrorAlert(resultErrorMessage);
      return false;
    }

    setIsLoading(true);

    const searchUserResponse = await SearchApi.searchUser(username);

    if (searchUserResponse.status === 200) {
      setSearchedUsers(searchUserResponse.body!);
    } else {
      alertsActions.addErrorAlert(searchUserResponse.message);
      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { fn, isLoading, searchedUsers };
};
