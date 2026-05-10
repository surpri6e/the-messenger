import styles from "./ContactsPage.module.scss";
import PrivatePageLayout from "@components/PrivatePageLayout/PrivatePageLayout";
import PeopleSection from "./Sections/PeopleSection/PeopleSection";
import SearchSection from "./Sections/SearchSection/SearchSection";

const ContactsPage = () => {
  return (
    <PrivatePageLayout>
      <div className={styles.wrapper}>
        <PeopleSection />
        <SearchSection />
      </div>
    </PrivatePageLayout>
  );
};

export default ContactsPage;
