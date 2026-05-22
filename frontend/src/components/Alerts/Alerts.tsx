import { useAlertsStore } from "@stores/useAlertsStore";
import styles from "./Alerts.module.scss";
import Alert from "./Alert";

const Alerts = () => {
  const { alerts } = useAlertsStore((state) => state);

  if (alerts.length === 0) {
    return <></>;
  }

  return (
    <div className={styles.alerts}>
      {alerts.map((alert) => (
        <Alert
          message={alert.message}
          id={alert.id}
          type={alert.type}
          key={alert.id}
        />
      ))}
    </div>
  );
};

export default Alerts;
