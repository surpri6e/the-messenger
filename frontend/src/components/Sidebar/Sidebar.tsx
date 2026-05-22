import styles from "./Sidebar.module.scss";
import SidebarAvatar from "./SidebarAvatar/SidebarAvatar";
import SidebarButton from "./SidebarButton/SidebarButton";

import chats from "@images/sidebar/chats.png";
import settings from "@images/sidebar/settings.png";
import contacts from "@images/user.png";
import {
  PRIVATE_ROUTE_CHATS,
  PRIVATE_ROUTE_CONTACTS,
  PRIVATE_ROUTE_SETTINGS,
} from "@constants/routes";

const Sidebar = () => {
  return (
    <div className={styles.sidebar}>
      <div className={styles.top}>
        <SidebarAvatar />
        <SidebarButton
          href={PRIVATE_ROUTE_CHATS}
          iconSrc={chats}
          isActive={
            location.pathname !== PRIVATE_ROUTE_CONTACTS &&
            location.pathname !== PRIVATE_ROUTE_SETTINGS
          }
        />
        <SidebarButton
          href={PRIVATE_ROUTE_CONTACTS}
          iconSrc={contacts}
          isActive={location.pathname === PRIVATE_ROUTE_CONTACTS}
        />
      </div>
      <div className={styles.bottom}>
        <SidebarButton
          href={PRIVATE_ROUTE_SETTINGS}
          iconSrc={settings}
          isActive={location.pathname === PRIVATE_ROUTE_SETTINGS}
        />
      </div>
    </div>
  );
};

export default Sidebar;
