import type { ITheme } from "@constants/themes";

export const setTheme = (theme: ITheme) => {
  Object.entries(theme).forEach(([key, value]) => {
    document.documentElement.style.setProperty(key, value);
  });
};
