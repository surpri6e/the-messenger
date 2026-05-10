import styles from "./ChatPage.module.scss";
import Chat from "./Chat/Chat";
import ChatsPanel from "./ChatsPanel/ChatsPanel";
import PrivatePageLayout from "@components/PrivatePageLayout/PrivatePageLayout";

const ChatPage = () => {
  return (
    <PrivatePageLayout>
      <div className={styles.wrapper}>
        <ChatsPanel />
        <Chat />
      </div>
    </PrivatePageLayout>
  );
};

export default ChatPage;
