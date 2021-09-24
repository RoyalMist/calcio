import React from "react";
import {Redirect, Route} from "react-router-dom";
import useAuthStore from "../../hooks/useAuthStore";

interface LoggedInRouteProps {
    children: React.ReactNode;
    mustBeAdmin?: boolean;
    path?: string;
    exact?: boolean;
}

function LoggedInRoute({children, exact, path, mustBeAdmin = false}: LoggedInRouteProps) {
    const auth = useAuthStore();
    const authorized = mustBeAdmin ? auth.isAdmin() : auth.isLoggedIn();
    return (
        <Route
            exact={exact}
            path={path}
            render={() =>
                authorized ? (children) : (
                    <Redirect to={{pathname: "/auth/login"}}/>
                )
            }
        />
    );
}

export default LoggedInRoute;
