import { useUserStore } from "@stores/useUserStore";
import styles from "./InformationSection.module.scss";
import { useState } from "react";
import { useExitFromAccount } from "@hooks/useExitFromAccount";
import { useInformationChange } from "@hooks/useInformationChange";
import { useCurrentChatStore } from "@stores/useCurrentChat";

const InformationSection = () => {
  const { user } = useUserStore((state) => state);
  const { actions: currentChatActions } = useCurrentChatStore((state) => state);

  const { fn: exitFromAccount } = useExitFromAccount();
  const {
    fn: changeInformation,
    informationError,
    isLoading,
    usernameError,
  } = useInformationChange();

  const [username, setUsername] = useState<string>(user!.username);
  const [info, setInfo] = useState<string>(user!.info);

  const onExitFromAccout = async () => {
    await exitFromAccount();
    currentChatActions.setCurrentChat(null);
    location.reload();
  };

  return (
    <section className={styles.informationSection}>
      <div className={styles.informationSectionBody}>
        <div className={styles.informationBlocks}>
          <div className={styles.informationBlock}>
            <p>Юзернейм:</p>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className={usernameError ? styles.error : ""}
            />
          </div>

          <div className={styles.informationBlock}>
            <p>Почта:</p>
            <input
              type="text"
              value={user?.email}
              disabled
              className={styles.disabledInput}
            />
          </div>

          <div className={styles.informationBlockTextarea}>
            <p>Информация:</p>
            <textarea
              value={info}
              onChange={(e) => setInfo(e.target.value)}
              className={informationError ? styles.error : ""}
            ></textarea>
          </div>
        </div>

        <div className={styles.buttons}>
          <button
            disabled={isLoading}
            className={`${styles.greenButton} ${isLoading ? styles.buttonDisabled : ""}`}
            onClick={async () => await changeInformation(username, info)}
          >
            Сохранить
          </button>
          <button className={styles.redButton} onClick={onExitFromAccout}>
            Выйти из аккаунта
          </button>
          <button className={styles.redButton}>Удалить аккаунт</button>
        </div>
      </div>
    </section>
  );
};

export default InformationSection;
