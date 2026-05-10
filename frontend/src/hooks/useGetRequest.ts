import type { IResponse } from "@appTypes/IResponse";
import type { IAllErrorMessages } from "@constants/apiErrors";
import { catchError } from "@functionals/catchError";
import { useAlertsStore } from "@stores/useAlertsStore";
import axios from "axios";
import { useEffect, useState } from "react";

export function useGetRequest<T>(
  str: string,
  errorMessages: IAllErrorMessages,
  deps?: React.DependencyList,
  cond?: boolean,
): [T | null, boolean, boolean] {
  const { actions: alertActions } = useAlertsStore((state) => state);
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);

  useEffect(
    () => {
      if (cond || cond == undefined) {
        try {
          setIsLoading(true);

          (async () => {
            const response = await axios.get<IResponse<T>>(str, {
              withCredentials: true,
            });

            setData(response.data.body);
          })();
        } catch (error) {
          const errorResponse = catchError(error, errorMessages);

          alertActions.addErrorAlert(errorResponse.message);

          setIsError(true);
        } finally {
          setIsLoading(false);
        }
      }
    },
    deps ? deps : [],
  );

  return [data, isLoading, isError];
}
