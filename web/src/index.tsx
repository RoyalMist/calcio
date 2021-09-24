import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import {QueryClient, QueryClientProvider} from "react-query";
import App from "./pages/app";
import LoggedOutRoute from "./components/route/logged_out";
import {BrowserRouter as Router, Switch} from "react-router-dom";
import Login from "./pages/login";
import LoggedInRoute from "./components/route/logged_in";
import {DefaultRedirect} from "./pages/404";

function Root() {
    const queryClient = new QueryClient({
        defaultOptions: {
            queries: {
                onError: err => console.error(err)
            },
            mutations: {
                onError: err => console.error(err)
            }
        },
    })

    return (
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
    );
}

ReactDOM.render(
    <React.StrictMode>
        <Root/>
    </React.StrictMode>,
    document.getElementById('root')
)
