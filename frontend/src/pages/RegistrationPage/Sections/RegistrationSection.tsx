import { useRegistrationSectionStore } from "@stores/useRegistrationSectionStore";
import { useRegistrationDataStore } from "@stores/useRegistrationDataStore";
import styles from "./MainSection.module.scss";
import { useEffect } from "react";
import { useRegistration } from "@hooks/useRegistration";
import { useAuth } from "@hooks/useAuth";
import { useGetUserData } from "@hooks/useGetUserData";
import Loader from "@components/Loader/Loader";

const RegistrationSection = () => {
  const { actions: sectionActions } = useRegistrationSectionStore(
    (state) => state,
  );

  const {
    email,
    username,
    password,
    actions: dataActions,
  } = useRegistrationDataStore((state) => state);

  const {
    fn: registration,
    emailError,
    passwordError,
    usernameError,
    isLoading: isLoadingRegistration,
  } = useRegistration();
  const { fn: auth, isLoading: isLoadingAuth } = useAuth();
  const { fn: getUserData, isLoading: isLoadingUserData } = useGetUserData();

  useEffect(() => {
    return () => {
      dataActions.setDefault();
    };
  }, []);

  const onClick = async () => {
    const registrationStatus = await registration(email, password, username);
    if (registrationStatus) {
      const authStatus = await auth(email, password);
      if (authStatus) {
        await getUserData();
        location.reload();
      }
    }
  };

  if (isLoadingRegistration || isLoadingAuth || isLoadingUserData) {
    return (
      <div className={styles.isLoading}>
        <Loader />
      </div>
    );
  }

  return (
    <>
      <h1>Регистрация</h1>

      <div>
        <div className={styles.inputs}>
          <section className={styles.inputBody}>
            <p>Почта:</p>
            <input
              type="email"
              placeholder="Введите почту..."
              value={email}
              onChange={(e) => dataActions.setNewEmail(e.target.value)}
              className={`${styles.input} ${emailError ? styles.errorInput : ""}`}
            />
          </section>

          <section className={styles.inputBody}>
            <p>Пароль:</p>
            <input
              type="password"
              placeholder="Введите пароль..."
              value={password}
              onChange={(e) => dataActions.setNewPassword(e.target.value)}
              className={`${styles.input} ${passwordError ? styles.errorInput : ""}`}
            />
          </section>

          <section className={styles.inputBody}>
            <p>Юзернейм:</p>
            <input
              type="text"
              placeholder="Введите юзернейм..."
              value={username}
              onChange={(e) => dataActions.setNewUsername(e.target.value)}
              className={`${styles.input} ${usernameError ? styles.errorInput : ""}`}
            />
          </section>
        </div>

        <div className={styles.buttons}>
          <button onClick={onClick}>Регистрация</button>

          <p className={styles.miniButtonBody}>
            <button onClick={sectionActions.setAuthType}>
              Войти в аккаунт
            </button>
          </p>
        </div>
      </div>
    </>
  );
};

export default RegistrationSection;
