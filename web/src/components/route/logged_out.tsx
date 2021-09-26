import React from "react";
import {Redirect, Route} from "react-router-dom";
import {useAuth} from "../../stores/authentication";

interface LoggedOutRouteProps {
    children: React.ReactNode;
    path?: string;
    exact?: boolean;
}

function LoggedOutRoute({exact, path, children}: LoggedOutRouteProps) {
    const {isLoggedIn} = useAuth();

    return (
        <Route
            exact={exact}
            path={path}
            render={({location}) =>
                !isLoggedIn() ? (children) : (
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
