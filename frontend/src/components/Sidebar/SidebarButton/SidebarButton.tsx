import { type FC } from "react";
import styles from "./SidebarButton.module.scss";
import { Link } from "react-router-dom";

interface ISidebarButton {
  href: string;
  iconSrc: string;
  isActive: boolean;
}

const SidebarButton: FC<ISidebarButton> = ({ href, iconSrc, isActive }) => {
  return (
    <Link
      to={href}
      className={`${styles.button} ${isActive ? styles.activeButton : ""}`}
    >
      <img src={iconSrc} className={styles.img} alt="Иконка боковой панели" />
    </Link>
  );
};

export default SidebarButton;
