import HeaderText from "@components/HeaderText/HeaderText";
import styles from "./PeopleSection.module.scss";

const PeopleSection = () => {
  return (
    <div className={styles.people}>
      <section className={styles.contactsSection}>
        <HeaderText text="Контакты" />
        <div className={styles.list}></div>
      </section>
      <section className={styles.enemiesSection}>
        <HeaderText text="Забаненные" />
        <div className={styles.list}></div>
      </section>
    </div>
  );
};

export default PeopleSection;
