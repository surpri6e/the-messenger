import { useRegistrationSectionStore } from "@stores/useRegistrationSectionStore";
import AuthSection from "./AuthSection";
import styles from "./MainSection.module.scss";
import RegistrationSection from "./RegistrationSection";
import RestoreSection from "./RestoreSection";

const MainSection = () => {
  const { sectionType } = useRegistrationSectionStore((state) => state);

  return (
    <section className={styles.body}>
      {sectionType === "registration" ? (
        <RegistrationSection />
      ) : sectionType === "auth" ? (
        <AuthSection />
      ) : (
        <RestoreSection />
      )}
    </section>
  );
};

export default MainSection;
