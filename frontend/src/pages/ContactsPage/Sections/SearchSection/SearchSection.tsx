import HeaderText from "@components/HeaderText/HeaderText";
import Search from "@components/Search/Search";
import styles from "./SearchSection.module.scss";
import { useSearchPeople } from "@hooks/useSearchPeople";
import Loader from "@components/Loader/Loader";
import basicAvatar from "@images/user.png";
import { PRIVATE_ROUTE_CHATS } from "@constants/routes";
import { useCurrentCommunicationStore } from "@stores/useCurrentCommunication";
import { useUserStore } from "@stores/useUserStore";

const SearchSection = () => {
  const { fn: searchUser, isLoading, searchedUsers } = useSearchPeople();
  const { actions: currentCommunicationActions } = useCurrentCommunicationStore(
    (state) => state,
  );
  const { user } = useUserStore((state) => state);

  const onOpenUserPage = () => {};

  const onAddContact = (e: React.MouseEvent<HTMLSpanElement, MouseEvent>) => {
    e.stopPropagation();
  };

  const onChat = (
    e: React.MouseEvent<HTMLParagraphElement, MouseEvent>,
    secondPersonId: number,
  ) => {
    e.stopPropagation();

    currentCommunicationActions.setCurrentChat({
      created_at: "",
      first_person_id: user!.id,
      second_person_id: secondPersonId,
      id: -1,
    });

    window.open(`${window.location.origin}${PRIVATE_ROUTE_CHATS}`, "_blank");
  };

  return (
    <section className={styles.searchSection}>
      <HeaderText text="Глобальный поиск людей" />
      <Search
        placeholder="Глобальный поиск людей"
        fn={searchUser}
        isLoading={isLoading}
      />
      <div className={`${styles.list} ${isLoading ? styles.loadingList : ""}`}>
        {isLoading ? (
          <Loader />
        ) : searchedUsers.length === 0 ? (
          <p className={styles.notFound}>
            Ничего не найдено или не было введено!
          </p>
        ) : (
          searchedUsers.map((user) => (
            <button
              className={styles.listElement}
              key={user.id}
              onClick={onOpenUserPage}
            >
              <div className={styles.listElementLeft}>
                <div className={styles.avatarBody}>
                  {user?.avatar_link.length === 0 ? (
                    <img
                      src={basicAvatar}
                      alt="Базовый аватар"
                      className={styles.basicAvatar}
                    />
                  ) : (
                    <img
                      src={user?.avatar_link}
                      alt="Аватар пользователя"
                      className={styles.avatar}
                    />
                  )}
                </div>

                <p>{user.username}</p>
              </div>

              <div className={styles.listElementRight}>
                <span
                  className={styles.button}
                  onClick={(e) => onAddContact(e)}
                >
                  <img src={basicAvatar} alt="Добавить в контакты" />
                </span>

                <p onClick={(e) => onChat(e, user.id)}>Сообщение...</p>
              </div>
            </button>
          ))
        )}
      </div>
    </section>
  );
};

export default SearchSection;
