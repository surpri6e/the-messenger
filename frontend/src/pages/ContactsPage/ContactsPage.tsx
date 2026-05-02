import HeaderText from "@components/HeaderText/HeaderText";
import styles from "./ContactsPage.module.scss";
import PrivatePageLayout from "@components/PrivatePageLayout/PrivatePageLayout";
import Search from "@components/Search/Search";

const ContactsPage = () => {
  return (
    <PrivatePageLayout>
      <div className={styles.wrapper}>
        <section className={styles.searchSection}>
          <HeaderText text="Глобальный поиск людей" />
          <Search placeholder="Глобальный поиск людей" />
        </section>
        <div className={styles.people}>
          <section className={styles.enemiesSection}></section>
          <section className={styles.contactsSection}></section>
        </div>
      </div>
    </PrivatePageLayout>
  );
};

export default ContactsPage;
