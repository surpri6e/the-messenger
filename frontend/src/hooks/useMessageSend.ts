import { useAlertsStore } from "@stores/useAlertsStore";
import { useState } from "react";
import { MessagesApi } from "@api/messagesApi";

interface IUseMessageSend {
  isLoading: boolean;

  fn: (message: string, where_id: number) => Promise<boolean>;
}

export const useMessageSend = (): IUseMessageSend => {
  const { actions: alertsActions } = useAlertsStore((state) => state);

  const [isLoading, setIsLoading] = useState(false);

  async function fn(message: string, where_id: number): Promise<boolean> {
    let resultErrorMessage = "";

    if (message.length === 0) {
      resultErrorMessage += "Сообщение не может быть пустым! ";
    }

    if (resultErrorMessage.length !== 0) {
      resultErrorMessage = resultErrorMessage.slice(0, -1);

      alertsActions.addErrorAlert(resultErrorMessage);
      return false;
    }

    setIsLoading(true);

    const sendMessageResponse = await MessagesApi.sendMessage(
      message,
      where_id,
    );

    if (sendMessageResponse.status !== 200) {
      alertsActions.addErrorAlert(sendMessageResponse.message);
      setIsLoading(false);
      return false;
    }

    setIsLoading(false);

    return true;
  }

  return { fn, isLoading };
};
