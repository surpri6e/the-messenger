import type { IRoute } from "@appTypes/IRoute";
import type { FC } from "react";
import { Route, Routes } from "react-router-dom";

interface IRoutesList {
  routes: IRoute[];
}

const RoutesList: FC<IRoutesList> = ({ routes }) => {
  return (
    <Routes>
      {routes.map((route) => (
        <Route path={route.path} element={<route.element />} />
      ))}
    </Routes>
  );
};

export default RoutesList;
