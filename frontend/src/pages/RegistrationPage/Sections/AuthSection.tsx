import { useEffect } from "react";
import { useAuthDataStore } from "@stores/useAuthDataStore";
import { useRegistrationSectionStore } from "@stores/useRegistrationSectionStore";
import styles from "./MainSection.module.scss";
import { useAuth } from "@hooks/useAuth";
import { useGetUserData } from "@hooks/useGetUserData";
import Loader from "@components/Loader/Loader";

const AuthSection = () => {
  const { actions: sectionActions } = useRegistrationSectionStore(
    (state) => state,
  );

  const {
    email,
    password,
    actions: dataActions,
  } = useAuthDataStore((state) => state);

  const {
    fn: auth,
    emailError,
    passwordError,
    isLoading: isLoadingAuth,
  } = useAuth();
  const { fn: getUserData, isLoading: isLoadingUserData } = useGetUserData();

  useEffect(() => {
    return () => {
      dataActions.setDefault();
    };
  }, []);

  const onClick = async () => {
    const authStatus = await auth(email, password);
    if (authStatus) {
      await getUserData();
      location.reload();
    }
  };

  if (isLoadingAuth || isLoadingUserData) {
    return (
      <div className={styles.isLoading}>
        <Loader />
      </div>
    );
  }

  return (
    <>
      <h1>Войти в аккаунт</h1>

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
        </div>

        <div className={styles.buttons}>
          <button onClick={onClick}>Войти</button>

          <p className={styles.miniButtonBody}>
            <button onClick={sectionActions.setRegistrationType}>
              Зарегистрироваться
            </button>

            <span> | </span>

            <button onClick={sectionActions.setRestoreType}>
              Восстановить
            </button>
          </p>
        </div>
      </div>
    </>
  );
};

export default AuthSection;
