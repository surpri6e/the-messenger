import { useState, type FC } from "react";
import styles from "./Search.module.scss";
import search from "@images/search.png";

interface ISearch {
  placeholder: string;

  isLoading: boolean;
  fn: (searchString: string) => Promise<boolean>;
}

const Search: FC<ISearch> = ({ placeholder, fn, isLoading }) => {
  const [searchString, setSearchString] = useState("");

  return (
    <div className={styles.search}>
      <button
        className={`${styles.button} ${isLoading ? styles.disabledButton : ""}`}
        onClick={async () => await fn(searchString)}
        disabled={isLoading}
      >
        <img src={search} alt={placeholder} />
      </button>
      <input
        type="text"
        className={styles.input}
        placeholder={placeholder}
        value={searchString}
        onChange={(e) => setSearchString(e.target.value)}
      />
    </div>
  );
};

export default Search;
