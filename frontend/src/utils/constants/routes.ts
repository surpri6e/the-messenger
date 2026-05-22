import type { IRoute } from "@appTypes/IRoute";
import ChatPage from "@pages/ChatPage/ChatPage";
import ContactsPage from "@pages/ContactsPage/ContactsPage";
import RegistrationPage from "@pages/RegistrationPage/RegistrationPage";
import SettingsPage from "@pages/SettingsPage/SettingsPage";

export const PUBLIC_ROUTES: IRoute[] = [
  { path: "*", element: RegistrationPage },
];

export const PRIVATE_ROUTE_CONTACTS = "/contacts";
export const PRIVATE_ROUTE_SETTINGS = "/settings";
export const PRIVATE_ROUTE_CHATS = "/chats";

export const PRIVATE_ROUTES: IRoute[] = [
  { path: "*", element: ChatPage },
  { path: PRIVATE_ROUTE_CONTACTS, element: ContactsPage },
  { path: PRIVATE_ROUTE_SETTINGS, element: SettingsPage },
];

export const ADMINS_ROUTES: IRoute[] = [];
