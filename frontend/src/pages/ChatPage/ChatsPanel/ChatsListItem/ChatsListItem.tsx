import type { IChat } from "@appTypes/IChat";
import styles from "./ChatsListItem.module.scss";
import { useGetRequest } from "@hooks/useGetRequest";
import type { ISmallUser } from "@appTypes/IUser";
import { USERS_ENDPOINT } from "@constants/endpoints";
import { GET_ONE_USER_ERROR_MESSAGES } from "@constants/apiErrors";
import { useUserStore } from "@stores/useUserStore";
import basicAvatar from "@images/user.png";
import Loader from "@components/Loader/Loader";
import { useCurrentChatStore } from "@stores/useCurrentChat";

const ChatsListItem: React.FC<{ chat: IChat }> = ({ chat }) => {
  const { user } = useUserStore((state) => state);
  const { actions: currentChatActions } = useCurrentChatStore((state) => state);

  const [data, isLoading, isError] = useGetRequest<ISmallUser>(
    USERS_ENDPOINT +
      `/${chat.first_person_id === user?.id ? chat.second_person_id : chat.first_person_id}`,
    GET_ONE_USER_ERROR_MESSAGES,
    [],
    700,
  );

  const onClick = () => {
    currentChatActions.setCurrentChat(chat);
  };

  return (
    <button
      className={styles.item}
      onClick={onClick}
      disabled={isLoading || isError}
    >
      {isLoading ? (
        <div className={styles.loading}>
          <Loader />
        </div>
      ) : isError ? (
        <p></p>
      ) : (
        <>
          <div className={styles.avatarBody}>
            {data?.avatar_link.length === 0 ? (
              <img
                src={basicAvatar}
                alt="Базовый аватар"
                className={styles.basicAvatar}
              />
            ) : (
              <img
                src={data?.avatar_link}
                alt="Аватар пользователя"
                className={styles.avatar}
              />
            )}

            <span
              className={`${styles.online} ${data?.is_online ? styles.isOnline : ""}`}
            ></span>
          </div>

          <p>{data?.username}</p>
        </>
      )}
    </button>
  );
};

export default ChatsListItem;
