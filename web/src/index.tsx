import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import {MutationCache, QueryCache, QueryClient, QueryClientProvider} from "react-query";
import App from "./pages/app";
import LoggedOutRoute from "./components/route/logged_out";
import {BrowserRouter as Router, Switch} from "react-router-dom";
import Login from "./pages/login";
import LoggedInRoute from "./components/route/logged_in";
import {DefaultRedirect} from "./pages/404";
import {toast, Toaster} from "react-hot-toast";

function Root() {
    const queryClient = new QueryClient({
        queryCache: new QueryCache({
            onError: (err) =>
                toast.error(`Something went wrong: ${err}`),
        }),
        mutationCache: new MutationCache({
            onError: (err) =>
                toast.error(`Something went wrong: ${err}`),
        })
    });

    return (<>
            <Toaster/>
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
        </>
    );
}

ReactDOM.render(
    <React.StrictMode>
        <Root/>
    </React.StrictMode>,
    document.getElementById('root')
)
