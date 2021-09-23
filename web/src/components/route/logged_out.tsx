import React from "react";
import {Redirect, Route} from "react-router-dom";
import useAuthStore from "../../hooks/useAuthStore";

interface LoggedOutRouteProps {
    children: React.ReactNode;
    path?: string;
    exact?: boolean;
}

function LoggedOutRoute({exact, path, children}: LoggedOutRouteProps) {
    const auth = useAuthStore();

    return (
        <Route
            exact={exact}
            path={path}
            render={({location}) =>
                !auth.isLoggedIn() ? (children) : (
                    <Redirect
                        to={{
                            pathname: "/",
                            state: {from: location},
                        }}
                    />
                )
            }
        />
    );
}

export default LoggedOutRoute;
