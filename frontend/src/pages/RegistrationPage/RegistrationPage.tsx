import styles from "./RegistrationPage.module.scss";
import background from "@images/background_image_registration_page.png";
import MainSection from "./Sections/MainSection";
import Alerts from "@components/Alerts/Alerts";

const RegistrationPage = () => {
  return (
    <main className={styles.page}>
      <img src={background} alt="Фоновое изображение" />

      <MainSection />

      <Alerts />
    </main>
  );
};

export default RegistrationPage;
