import { useGetRequest } from "@hooks/useGetRequest";
import styles from "./ChatMessages.module.scss";
import { LIST_ENDPOINT, MESSAGES_ENDPOINT } from "@constants/endpoints";
import { useCurrentChatStore } from "@stores/useCurrentChat";
import { GET_MESSAGES_LIST_ERROR_MESSAGES } from "@constants/apiErrors";
import type { IMessage } from "@appTypes/IMessage";
import { useUserStore } from "@stores/useUserStore";
import { useEffect, useRef } from "react";
import trash from "@images/trash.png";
import axios from "axios";

const ChatMessages = () => {
  const { currentChat } = useCurrentChatStore((state) => state);
  const { user } = useUserStore((state) => state);

  const messagesEndRef = useRef<HTMLDivElement | null>(null);

  const [messages, isLoading, isError] = useGetRequest<IMessage[]>(
    LIST_ENDPOINT + `/communication/chat/${currentChat?.id}`,
    GET_MESSAGES_LIST_ERROR_MESSAGES,
    [currentChat],
    350,
  );

  const prevMessagesLengthRef = useRef<number>(0);

  useEffect(() => {
    if (messages && messages.length > 0) {
      const hasNewMessage = messages.length > prevMessagesLengthRef.current;

      if (hasNewMessage) {
        messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
      }

      prevMessagesLengthRef.current = messages.length;
    }
  }, [messages]);

  const onClick = async (message_id: number) => {
    await axios.delete(MESSAGES_ENDPOINT + `/${message_id}`, {
      withCredentials: true,
    });
  };

  return (
    <div className={styles.messages} id="messages">
      {messages ? (
        messages.map((message) => (
          <div
            key={message.id}
            className={`${user!.id === message.user_id ? styles.myMessage : styles.message}`}
          >
            <div
              className={`${user!.id === message.user_id ? styles.myMessageBody : styles.messageBody}`}
            >
              <p>{message.text}</p>
              <section>
                <time>{message.created_at.slice(11, 16)}</time>
                {user!.id === message.user_id ? (
                  <button onClick={async () => onClick(message.id)}>
                    <img src={trash} alt="Удалить" />
                  </button>
                ) : (
                  <></>
                )}
              </section>
            </div>
          </div>
        ))
      ) : (
        <></>
      )}
      <div id="invize" ref={messagesEndRef}></div>
    </div>
  );
};

export default ChatMessages;
