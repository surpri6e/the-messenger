import Sidebar from "@components/Sidebar/Sidebar";
import styles from "./PrivatePageLayout.module.scss";
import Alerts from "@components/Alerts/Alerts";
import type { FC } from "react";

interface IPrivatePageLayout {
  children: React.ReactNode;
}

const PrivatePageLayout: FC<IPrivatePageLayout> = ({ children }) => {
  return (
    <main className={styles.page}>
      <Sidebar />
      <div className={styles.body}>{children}</div>

      <Alerts />
    </main>
  );
};

export default PrivatePageLayout;
