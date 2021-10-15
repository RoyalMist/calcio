import React from "react";
import {Redirect, Route} from "react-router-dom";
import {useAuth} from "../stores/authentication";

interface LoggedInRouteProps {
    children: React.ReactNode;
    mustBeAdmin?: boolean;
    path?: string;
    exact?: boolean;
}

function LoggedInRoute({children, exact, path, mustBeAdmin = false}: LoggedInRouteProps) {
    const {isAdmin, isLoggedIn} = useAuth();
    if (mustBeAdmin && isLoggedIn()) {
        return <Route
            exact={exact}
            path={path}
            render={() =>
                isAdmin() ? (children) : (
                    <Redirect to={{pathname: "/"}}/>
                )
            }
        />
    } else {
        return (
            <Route
                exact={exact}
                path={path}
                render={() =>
                    isLoggedIn() ? (children) : (
                        <Redirect to={{pathname: "/login"}}/>
                    )
                }
            />
        );
    }
}

export default LoggedInRoute;
