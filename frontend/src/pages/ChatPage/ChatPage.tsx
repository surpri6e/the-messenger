import Sidebar from "@components/Sidebar/Sidebar";
import Chat from "./Chat/Chat";
import styles from "./ChatPage.module.scss";
import ChatPanel from "./ChatPanel/ChatPanel";
import Alerts from "@components/Alerts/Alerts";

const ChatPage = () => {
  return (
    <main className={styles.page}>
      <Sidebar />
      <div className={styles.body}>
        <ChatPanel />
        <Chat />
      </div>

      <Alerts />
    </main>
  );
};

export default ChatPage;
