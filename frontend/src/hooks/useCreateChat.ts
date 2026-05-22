import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";
import { ChatsApi } from "@api/chatsApi";
import type { IChat } from "@appTypes/IChat";

interface IUseCreateChat {
  isLoading: boolean;

  fn: (
    first_person_id: number,
    second_person_id: number,
  ) => Promise<IChat | null>;
}

export const useCreateChat = (): IUseCreateChat => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(
    first_person_id: number,
    second_person_id: number,
  ): Promise<IChat | null> {
    setIsLoading(true);

    const createChatResponse = await ChatsApi.createChat(
      first_person_id,
      second_person_id,
    );

    if (createChatResponse.status !== 200) {
      alertsActions.addErrorAlert(createChatResponse.message);
      setIsLoading(false);
      return null;
    }

    setIsLoading(false);

    return createChatResponse.body;
  }

  return { fn, isLoading };
};
