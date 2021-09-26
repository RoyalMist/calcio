import React from "react";
import {Redirect, Route} from "react-router-dom";
import {useAuth} from "../../stores/authentication";

interface LoggedInRouteProps {
    children: React.ReactNode;
    mustBeAdmin?: boolean;
    path?: string;
    exact?: boolean;
}

function LoggedInRoute({children, exact, path, mustBeAdmin = false}: LoggedInRouteProps) {
    const {authState} = useAuth();
    const authorized = mustBeAdmin ? authState.paseto?.is_admin : !!authState.token;
    return (
        <Route
            exact={exact}
            path={path}
            render={() =>
                authorized ? (children) : (
                    <Redirect to={{pathname: "/login"}}/>
                )
            }
        />
    );
}

export default LoggedInRoute;
