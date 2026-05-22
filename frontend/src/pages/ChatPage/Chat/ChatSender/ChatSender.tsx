import styles from "./ChatSender.module.scss";
import clip from "@images/clip.png";
import arrow from "@images/arrow.png";
import microphone from "@images/microphone.png";
import { useState } from "react";
import { useCurrentChatStore } from "@stores/useCurrentChat";
import { useCreateChat } from "@hooks/useCreateChat";
import { useMessageSend } from "@hooks/useMessageSend";

const ChatSender = () => {
  const { currentChat, actions: currentCommunicationActions } =
    useCurrentChatStore((state) => state);

  const { fn: createChat } = useCreateChat();
  const { fn: sendMessage } = useMessageSend();
  const [message, setMessage] = useState("");

  const onClick = async () => {
    if (currentChat!.id === -1) {
      const newChat = await createChat(
        currentChat!.first_person_id,
        currentChat!.second_person_id,
      );

      if (newChat) {
        currentCommunicationActions.setCurrentChat(newChat);
        await sendMessage(message, newChat?.id);
      }
    }

    if (currentChat?.id) {
      await sendMessage(message, currentChat.id);
    }

    setMessage("");
  };

  return (
    <div className={styles.sender}>
      <div className={styles.left}>
        <button className={styles.button}>
          <img src={clip} alt="Прикрепить файл" className={styles.invert} />
        </button>

        <input
          type="text"
          className={styles.input}
          placeholder="Сообщение..."
          value={message}
          onChange={(e) => setMessage(e.target.value)}
        />
      </div>

      <div className={styles.right}>
        <button className={styles.button}>
          <img
            src={microphone}
            alt="Записать голосовое сообщение"
            className={styles.invert}
          />
        </button>

        <button className={styles.button} onClick={onClick}>
          <img
            src={arrow}
            alt="Отправить сообщение"
            className={styles.invert}
          />
        </button>
      </div>
    </div>
  );
};

export default ChatSender;
