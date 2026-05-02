import AvatarSection from "./Sections/AvatarSection/AvatarSection";
import InformationSection from "./Sections/InformationSection/InformationSection";
import ThemeSection from "./Sections/ThemeSection/ThemeSection";
import styles from "./SettingsPage.module.scss";
import PrivatePageLayout from "@components/PrivatePageLayout/PrivatePageLayout";

const SettingsPage = () => {
  return (
    <PrivatePageLayout>
      <div className={styles.wrapper}>
        <div className={styles.wrapperBody}>
          <AvatarSection />
          <InformationSection />
          <ThemeSection />
        </div>
      </div>
    </PrivatePageLayout>
  );
};

export default SettingsPage;
