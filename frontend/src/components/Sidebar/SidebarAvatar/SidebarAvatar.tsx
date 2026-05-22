import { Link } from "react-router-dom";
import styles from "./SidebarAvatar.module.scss";
import { PRIVATE_ROUTE_SETTINGS } from "@constants/routes";
import { useUserStore } from "@stores/useUserStore";
import basicAvatar from "@images/user.png";

const SidebarAvatar = () => {
  const { user } = useUserStore((state) => state);

  return (
    <Link to={PRIVATE_ROUTE_SETTINGS} className={styles.avatarBody}>
      {user?.avatar_link.length === 0 ? (
        <img
          src={basicAvatar}
          alt="Ваш базовый аватар"
          className={styles.basicAvatar}
        />
      ) : (
        <img
          src={user?.avatar_link}
          alt="Ваш аватар"
          className={styles.avatar}
        />
      )}
    </Link>
  );
};

export default SidebarAvatar;
