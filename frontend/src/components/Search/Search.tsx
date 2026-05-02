import type { FC } from "react";
import styles from "./Search.module.scss";
import search from "@images/search.png";

interface ISearch {
  placeholder: string;
}

const Search: FC<ISearch> = ({ placeholder }) => {
  return (
    <div className={styles.search}>
      <button className={styles.button}>
        <img src={search} alt={placeholder} />
      </button>
      <input type="text" className={styles.input} placeholder={placeholder} />
    </div>
  );
};

export default Search;
