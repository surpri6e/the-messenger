import { useEffect, useState } from "react";

export const useErrorDelay = (
  delay: number,
): [boolean, React.Dispatch<React.SetStateAction<boolean>>] => {
  const [error, setError] = useState(false);

  useEffect(() => {
    if (error) {
      setTimeout(() => setError(false), delay);
    }
  }, [error]);

  return [error, setError];
};
