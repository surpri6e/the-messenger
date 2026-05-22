import { useCurrentChatStore } from "@stores/useCurrentChat";
import styles from "./Chat.module.scss";
import ChatHeader from "./ChatHeader/ChatHeader";
import ChatSender from "./ChatSender/ChatSender";
import ChatMessages from "./ChatMessages/ChatMessages";

const Chat = () => {
  const { currentChat } = useCurrentChatStore((state) => state);
  return (
    <div className={styles.chat}>
      <div className={styles.chatBody}>
        {currentChat === null ? (
          <></>
        ) : (
          <>
            <ChatHeader />

            {currentChat.id === -1 ? <></> : <ChatMessages />}

            <ChatSender />
          </>
        )}
      </div>
    </div>
  );
};

export default Chat;
