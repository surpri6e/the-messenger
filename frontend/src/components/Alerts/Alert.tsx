import { useAlertsStore, type IAlert } from "@stores/useAlertsStore";
import { useState, type FC } from "react";
import styles from "./Alerts.module.scss";
import errorAlert from "@images/alerts/error.png";
import successAlert from "@images/alerts/success.png";
import cross from "@images/alerts/cross.png";

const Alert: FC<IAlert> = ({ message, type, id }) => {
  const { actions } = useAlertsStore((state) => state);

  const [isRemoving, setIsRemoving] = useState(false);

  function deleteAlert(id: number) {
    setIsRemoving(true);

    setTimeout(() => {
      actions.deleteAlert(id);
    }, 350);
  }

  return (
    <div
      className={`${styles.alert} ${type === "error" ? styles.alertError : styles.alertSuccess} ${isRemoving ? styles.alertRemoving : ""}`}
    >
      <div className={styles.alertHeader}>
        <div className={styles.alertHeaderLeft}>
          {type === "error" ? (
            <>
              <img src={errorAlert} alt="Произошла ошибка" />
              <h4>Произошла ошибка</h4>
            </>
          ) : (
            <>
              <img src={successAlert} alt="Все выполнено успешно" />
              <h4>Все выполнено успешно</h4>
            </>
          )}
        </div>

        <button onClick={() => deleteAlert(id)}>
          <img src={cross} alt="Закрыть подсказку" />
        </button>
      </div>

      <p>{message}</p>
    </div>
  );
};

export default Alert;
