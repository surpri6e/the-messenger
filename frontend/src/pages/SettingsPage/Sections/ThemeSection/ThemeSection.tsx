import HeaderText from "@components/HeaderText/HeaderText";
import styles from "./ThemeSection.module.scss";
import { PURPLE1_THEME, PURPLE2_THEME, PURPLE3_THEME } from "@constants/themes";
import { setTheme } from "@functionals/setTheme";
import type { IUser } from "@appTypes/IUser";
import { useEffect, useState } from "react";
import { useThemeChange } from "@hooks/useThemeChange";
import Loader from "@components/Loader/Loader";

const ThemeSection = () => {
  const [themeName, setThemeName] = useState<IUser["theme"] | "">("");

  const { isLoading, fn: changeTheme } = useThemeChange();

  const onClick = async () => {
    if (themeName.length !== 0) {
      await changeTheme(themeName);
    }
  };

  useEffect(() => {
    onClick();
  }, [themeName]);

  return (
    <section className={styles.themeSection}>
      <div className={styles.themeSectionBody}>
        <HeaderText text="Стиль" isDarken />

        <div className={styles.themes}>
          {isLoading ? (
            <div className={styles.themesLoaderBody}>
              <Loader />
            </div>
          ) : (
            <>
              <span
                className={styles.theme}
                style={{ backgroundColor: PURPLE1_THEME["--second-color"] }}
                onClick={() => {
                  setThemeName("purple1");
                  setTheme(PURPLE1_THEME);
                }}
              ></span>
              <span
                className={styles.theme}
                style={{ backgroundColor: PURPLE2_THEME["--second-color"] }}
                onClick={() => {
                  setThemeName("purple2");
                  setTheme(PURPLE2_THEME);
                }}
              ></span>
              <span
                className={styles.theme}
                style={{ backgroundColor: PURPLE3_THEME["--second-color"] }}
                onClick={() => {
                  setThemeName("purple3");
                  setTheme(PURPLE3_THEME);
                }}
              ></span>
            </>
          )}
        </div>
      </div>
    </section>
  );
};

export default ThemeSection;
