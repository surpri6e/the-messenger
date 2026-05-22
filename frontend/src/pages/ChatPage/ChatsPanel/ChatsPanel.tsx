import styles from "./ChatsPanel.module.scss";
import ChatsListItem from "./ChatsListItem/ChatsListItem";
import HeaderText from "@components/HeaderText/HeaderText";
import { useGetRequest } from "@hooks/useGetRequest";
import { LIST_ENDPOINT } from "@constants/endpoints";
import { GET_ONE_USER_ERROR_MESSAGES } from "@constants/apiErrors";
import type { IChat } from "@appTypes/IChat";
import Loader from "@components/Loader/Loader";
import { useUserStore } from "@stores/useUserStore";

const ChatsPanel = () => {
  const { user } = useUserStore((state) => state);

  // Если кто-то создаст со мной чат и напишет, то обновится ли список? deps?: [globalMessages]
  const [chats, isLoadingChats, isErrorChats] = useGetRequest<IChat[]>(
    LIST_ENDPOINT + `/chats/user/${user?.id}`,
    GET_ONE_USER_ERROR_MESSAGES, // ДРУГОЙ ТИП ОШИБКИ
    [],
    700,
  );

  return (
    <div className={styles.panel}>
      <div className={styles.top}>
        <HeaderText text="Чаты" />
        {/* <Search
          placeholder="Поиск среди своих чатов"
          isLoading={false}
          fn={async () => true}
        /> */}
      </div>

      <div className={styles.list}>
        {isLoadingChats ? (
          <div className={styles.loading}>
            <Loader />
          </div>
        ) : isErrorChats ? (
          <p className={styles.error}>Произошла ошибка!</p>
        ) : chats && chats.length !== 0 ? (
          [...chats!].map((chat) => <ChatsListItem chat={chat} key={chat.id} />)
        ) : (
          <p className={styles.notFound}>Ничего нет!</p>
        )}
      </div>
    </div>
  );
};

export default ChatsPanel;
