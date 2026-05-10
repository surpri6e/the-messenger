import type { IChat, IGroup } from "@appTypes/ICommunication";
import styles from "./ChatsListItem.module.scss";
import { useGetRequest } from "@hooks/useGetRequest";
import type { ISmallUser } from "@appTypes/IUser";
import { USERS_ENDPOINT } from "@constants/endpoints";
import { GET_ONE_USER_ERROR_MESSAGES } from "@constants/apiErrors";
import { useUserStore } from "@stores/useUserStore";

interface IChatsListItem {
  communication: IChat | IGroup;
}

const ChatsListItem: React.FC<IChatsListItem> = ({ communication }) => {
  const { user } = useUserStore((state) => state);

  let chat: IChat | undefined = undefined;
  let group: IGroup | undefined = undefined;

  if (communication.id % 2 !== 0) {
    chat = communication as IChat;
  } else {
    group = communication as IGroup;
  }

  const [data, isLoading, isError] = useGetRequest<ISmallUser>(
    USERS_ENDPOINT +
      `/${chat?.first_person_id === user?.id ? chat?.second_person_id : chat?.first_person_id}`,
    GET_ONE_USER_ERROR_MESSAGES,
    [chat],
    chat ? true : false,
  );

  return (
    <div className={styles.item}>
      {!isLoading && !isError ? (
        <div className={styles.avatarBody}>
          {/* <img src={avatarSrc} alt="avatar" />
          <div
            className={
              isOnline
                ? styles.online
                : `${styles.online + " " + styles.yesOnline}`
            }
          ></div> */}

          <p>{data?.username}</p>
        </div>
      ) : (
        <></>
      )}
    </div>
  );
};

export default ChatsListItem;
