import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import {MutationCache, QueryCache, QueryClient, QueryClientProvider} from "react-query";
import App from "./pages/app";
import LoggedOutRoute from "./components/LoggedOutRoute";
import {BrowserRouter as Router, Switch} from "react-router-dom";
import Login from "./pages/login";
import LoggedInRoute from "./components/LoggedInRoute";
import {DefaultRedirect} from "./pages/404";
import {toast, Toaster} from "react-hot-toast";
import {AuthProvider} from "./stores/authentication";

function Root() {
    const queryClient = new QueryClient({
        queryCache: new QueryCache({
            onError: (err) =>
                toast.error(`${err}`),
        }),
        mutationCache: new MutationCache({
            onError: (err) =>
                toast.error(`${err}`),
        })
    });

    return (<>
            <Toaster/>
            <AuthProvider>
                <QueryClientProvider client={queryClient}>
                    <Router>
                        <Switch>
                            <LoggedOutRoute path="/login">
                                <Login/>
                            </LoggedOutRoute>
                            <LoggedInRoute path="/">
                                <App/>
                            </LoggedInRoute>
                            <DefaultRedirect/>
                        </Switch>
                    </Router>
                </QueryClientProvider>
            </AuthProvider>
        </>
    );
}

ReactDOM.render(
    <React.StrictMode>
        <Root/>
    </React.StrictMode>,
    document.getElementById('root')
)
