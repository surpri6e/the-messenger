import { useUserStore } from "@stores/useUserStore";
import styles from "./InformationSection.module.scss";

const InformationSection = () => {
  const { user } = useUserStore((state) => state);
  return (
    <section className={styles.informationSection}>
      <div className={styles.informationSectionBody}>
        <div className={styles.informationBlock}>
          <p>Юзернейм</p>
          <input type="text" value={user?.username} />
        </div>

        <div className={styles.informationBlock}>
          <p>Почта</p>
          <input type="text" value={user?.email} />
        </div>

        <div className={styles.informationBlock}>
          <p>Информация</p>
          <textarea name="" id="" value={user?.info}></textarea>
        </div>

        <div className={styles.buttons}>
          <button>Сохранить</button>
          <button>Выйти из аккаунта</button>
          <button>Удалить аккаунт</button>
        </div>
      </div>
    </section>
  );
};

export default InformationSection;
