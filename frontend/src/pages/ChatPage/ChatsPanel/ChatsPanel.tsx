import styles from "./ChatsPanel.module.scss";
import ChatsListItem from "./ChatsListItem/ChatsListItem";
import HeaderText from "@components/HeaderText/HeaderText";
import Search from "@components/Search/Search";
import { useGetRequest } from "@hooks/useGetRequest";
import { LIST_ENDPOINT } from "@constants/endpoints";
import { useUserStore } from "@stores/useUserStore";
import { GET_CHATS_AND_GROUPS_LIST_ERROR_MESSAGES } from "@constants/apiErrors";
import type { IChat, IGroup } from "@appTypes/ICommunication";
import Loader from "@components/Loader/Loader";

const ChatsPanel = () => {
  const { user } = useUserStore((state) => state);

  // Если кто-то создаст со мной чат и напишет, то обновится ли список? deps?: [globalMessages]
  const [chats, isLoadingChats, isErrorChats] = useGetRequest<IChat[]>(
    LIST_ENDPOINT + `/${user!.id}/chats`,
    GET_CHATS_AND_GROUPS_LIST_ERROR_MESSAGES,
  );

  const [groups, isLoadingGroups, isErrorGroups] = useGetRequest<IGroup[]>(
    LIST_ENDPOINT + `/${user!.id}/groups`,
    GET_CHATS_AND_GROUPS_LIST_ERROR_MESSAGES,
  );

  console.log(chats, groups);

  return (
    <div className={styles.panel}>
      <div className={styles.top}>
        <HeaderText text="Чаты" />
        <Search
          placeholder="Поиск среди своих чатов"
          isLoading={false}
          fn={async () => true}
        />
      </div>

      <div className={styles.list}>
        {isLoadingChats || isLoadingGroups ? (
          <Loader />
        ) : isErrorChats || isErrorGroups ? (
          <p className={styles.error}>Произошла ошибка!</p>
        ) : chats && chats.length !== 0 && groups && groups.length !== 0 ? (
          [...chats!, ...groups!].map((communication) => (
            <ChatsListItem
              communication={communication}
              key={communication.id}
            />
          ))
        ) : (
          <p className={styles.notFound}>Ничего нет!</p>
        )}
      </div>
    </div>
  );
};

export default ChatsPanel;
