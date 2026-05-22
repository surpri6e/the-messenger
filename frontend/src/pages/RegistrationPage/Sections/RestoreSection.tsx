import { useEffect, useState } from "react";
import { useRegistrationSectionStore } from "@stores/useRegistrationSectionStore";
import styles from "./MainSection.module.scss";
import { useRestoreDataStore } from "@stores/useRestoreDataStore";

const RestoreSection = () => {
  const { actions: sectionActions } = useRegistrationSectionStore(
    (state) => state,
  );

  const {
    email,
    password,
    code,
    actions: dataActions,
  } = useRestoreDataStore((state) => state);

  const [isFirst, setIsFirst] = useState(true);

  // const [emailError, setEmailError] = useErrorDelay(1000);
  //   const [passwordError, setPasswordError] = useErrorDelay(1000);
  //   const [usernameError, setUsernameError] = useErrorDelay(1000);

  useEffect(() => {
    return () => {
      dataActions.setDefault();
    };
  }, []);

  // function sendCode(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
  //   e.preventDefault();
  //   setIsFirst(false);
  // }

  // async function sendData(email: string, password: string) {
  //   let resultErrorMessage = "";

  //   if (!isEmailCorrectly(email)) {
  //     resultErrorMessage = "Почта введена неверно! ";
  //     setEmailError(true);
  //   }

  //   if (!(password.length >= 5 && password.length <= 20)) {
  //     resultErrorMessage += "Пароль должен быть от 5 до 20 символов! ";
  //     setPasswordError(true);
  //   }

  //   if (resultErrorMessage.length !== 0) {
  //     resultErrorMessage = resultErrorMessage.slice(0, -1);

  //     alertsActions.addErrorAlert(resultErrorMessage);
  //     return;
  //   }
  // }

  return (
    <>
      <h1>Восстановить пароль</h1>

      <div>
        <div className={styles.inputs}>
          <section className={styles.inputBody}>
            <p>Почта:</p>
            <input
              type="email"
              placeholder="Введите почту..."
              value={email}
              onChange={(e) => dataActions.setNewEmail(e.target.value)}
              // className={`${styles.input} ${emailError ? styles.errorInput : ""}`}
              className={`${styles.input}`}
            />
          </section>

          <section className={styles.inputBody}>
            <p>Новый пароль:</p>
            <input
              type="password"
              placeholder="Введите новый пароль..."
              value={password}
              onChange={(e) => dataActions.setNewPassword(e.target.value)}
              // className={`${styles.input} ${emailError ? styles.errorInput : ""}`}
              className={`${styles.input}`}
            />
          </section>

          <section className={styles.inputBody}>
            <p>Код из сообщения:</p>
            <input
              type="password"
              placeholder="Введите код из сообщения..."
              value={code}
              onChange={(e) => dataActions.setNewCode(e.target.value)}
              // className={`${styles.input} ${emailError ? styles.errorInput : ""}`}
              className={`${styles.input}`}
            />

            <p className={styles.miniButtonBody}>
              {/* <button onClick={(e) => onSendCode(e)}> */}
              <button onClick={() => setIsFirst(false)}>
                {isFirst ? "Отправить код на почту" : "Отправить код повторно"}
              </button>
            </p>
          </section>
        </div>

        <div className={styles.buttons}>
          <button onSubmit={() => true}>Изменить пароль</button>

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

export default RestoreSection;
