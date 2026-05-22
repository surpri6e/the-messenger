import { useUserStore } from "@stores/useUserStore";
import {
  ADMINS_ROUTES,
  PRIVATE_ROUTES,
  PUBLIC_ROUTES,
} from "./utils/constants/routes";
import RoutesList from "@components/RoutesList";
import { useGetUserData } from "@hooks/useGetUserData";
import { useEffect, useState } from "react";
import Loader from "@components/Loader/Loader";
import { setTheme } from "@functionals/setTheme";
import { PURPLE1_THEME, PURPLE2_THEME, PURPLE3_THEME } from "@constants/themes";
import axios from "axios";
import { USERS_ENDPOINT } from "@constants/endpoints";

// cd C:\Program Files\Google\Chrome\Application
// chrome.exe --ignore-certificate-errors

// Нужно будет куку из таблицы удалять
// Протестить как ведет себя api, когда кончается срок токена но ты в онлайне

function App() {
  const [isLoading, setIsLoading] = useState(true);
  const { fn: getUserData } = useGetUserData();

  const { user } = useUserStore((state) => state);

  useEffect(() => {
    (async () => {
      await getUserData(true);
      setIsLoading(false);
    })();
  }, []);

  useEffect(() => {
    const setIsOnline = async () => {
      try {
        await axios.put(
          USERS_ENDPOINT + "/online",
          {},
          { withCredentials: true },
        );
      } catch (error) {
        console.error("Ошибка:", error);
      }
    };

    if (user && !user.is_admin) {
      setIsOnline();

      const interval = setInterval(setIsOnline, 2222);

      return () => clearInterval(interval);
    }
  }, [user]);

  if (isLoading) {
    return <Loader />;
  }

  if (!user) {
    return <RoutesList routes={PUBLIC_ROUTES} />;
  }

  if (user && !user.is_admin) {
    // Выбор темы отрисовки

    switch (user.theme) {
      case "purple1": {
        setTheme(PURPLE1_THEME);
        break;
      }
      case "purple2": {
        setTheme(PURPLE2_THEME);
        break;
      }
      case "purple3": {
        setTheme(PURPLE3_THEME);
        break;
      }
    }

    return <RoutesList routes={PRIVATE_ROUTES} />;
  }

  if (user && user.is_admin) {
    return <RoutesList routes={ADMINS_ROUTES} />;
  }
}

export default App;
