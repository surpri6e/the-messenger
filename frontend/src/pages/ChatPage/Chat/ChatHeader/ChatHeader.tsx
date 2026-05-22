import type { ISmallUser } from "@appTypes/IUser";
import { GET_ONE_USER_ERROR_MESSAGES } from "@constants/apiErrors";
import { USERS_ENDPOINT } from "@constants/endpoints";
import { useGetRequest } from "@hooks/useGetRequest";
import { useCurrentChatStore } from "@stores/useCurrentChat";
import { useUserStore } from "@stores/useUserStore";
import styles from "./ChatHeader.module.scss";
import Loader from "@components/Loader/Loader";
import phone from "@images/phone.png";
import settings from "@images/settigns.png";
import basicAvatar from "@images/user.png";
import HeaderText from "@components/HeaderText/HeaderText";

const ChatHeader = () => {
  const { user } = useUserStore((state) => state);
  const { currentChat } = useCurrentChatStore((state) => state);

  const [userInfo, isLoadingUserInfo, isErrorUserInfo] =
    useGetRequest<ISmallUser>(
      USERS_ENDPOINT +
        `/${currentChat?.first_person_id === user?.id ? currentChat?.second_person_id : currentChat?.first_person_id}`,
      GET_ONE_USER_ERROR_MESSAGES,
      [currentChat],
      700,
    );

  return (
    <div className={styles.header}>
      {isLoadingUserInfo ? (
        <div className={styles.center}>
          <Loader />
        </div>
      ) : isErrorUserInfo ? (
        <div className={styles.center}>
          <p className={styles.error}>Не удалось загрузить данные</p>
        </div>
      ) : (
        <>
          <div className={styles.headerLeft}>
            <div className={styles.avatarBody}>
              {userInfo?.avatar_link.length === 0 ? (
                <img
                  src={basicAvatar}
                  alt="Базовый аватар пользователя"
                  className={styles.basicAvatar}
                />
              ) : (
                <img
                  src={userInfo?.avatar_link}
                  alt="Аватар пользователя"
                  className={styles.avatar}
                />
              )}

              <span
                className={`${styles.online} ${userInfo?.is_online ? styles.isOnline : ""}`}
              ></span>
            </div>

            <div className={styles.text}>
              <HeaderText
                text={userInfo?.username ? userInfo.username : ""}
                isDarken
              />
              <p
                className={`${styles.onlineText} ${userInfo?.is_online ? styles.isOnlineText : ""}`}
              >
                {userInfo?.is_online ? "На связи" : "Спит"}
              </p>
            </div>
          </div>

          <div className={styles.headerRight}>
            {/* <button>
              <img src={phone} alt="Позвонить" />
            </button> */}

            <button>
              <img src={settings} alt="Настройки" />
            </button>
          </div>
        </>
      )}
    </div>
  );
};

export default ChatHeader;
