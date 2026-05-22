import { type FC } from "react";
import styles from "./HeaderText.module.scss";

interface IHeaderText {
  text: string;
  isDarken?: boolean;
}

const HeaderText: FC<IHeaderText> = ({ text, isDarken }) => {
  return <h2 className={isDarken ? styles.darkenText : styles.text}>{text}</h2>;
};

export default HeaderText;
