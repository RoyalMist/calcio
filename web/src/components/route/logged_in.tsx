import React from "react";
import {Redirect, Route} from "react-router-dom";
import useAuthStore from "../../hooks/useAuthStore";

interface LoggedInRouteProps {
    children: React.ReactNode;
    withRoles?: string[];
    path?: string;
    exact?: boolean;
}

function LoggedInRoute({children, exact, path, withRoles}: LoggedInRouteProps) {
    const auth = useAuthStore();
    const authorized = withRoles !== undefined ? withRoles.some(r => auth.hasRole(r)) : auth.isLoggedIn()

    return (
        <Route
            exact={exact}
            path={path}
            render={({location}) =>
                authorized ? (children) : (
                    <Redirect
                        to={{
                            pathname: "/auth/login",
                            state: {from: location},
                        }}
                    />
                )
            }
        />
    );
}

export default LoggedInRoute;
