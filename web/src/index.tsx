import React from "react";
import ReactDOM from "react-dom";
import { toast, Toaster } from "react-hot-toast";
import {
  MutationCache,
  QueryCache,
  QueryClient,
  QueryClientProvider,
} from "react-query";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import LoggedInRoute from "./components/LoggedInRoute";
import LoggedOutRoute from "./components/LoggedOutRoute";
import "./index.css";
import { DefaultRedirect } from "./pages/404";
import App from "./pages/app";
import Login from "./pages/login";
import { AuthProvider } from "./stores/authentication";

const queryClient = new QueryClient({
  queryCache: new QueryCache({
    onError: (err) => toast.error(`${err}`),
  }),
  mutationCache: new MutationCache({
    onError: (err) => toast.error(`${err}`),
  }),
});

function Root() {
  return (
    <>
      <Toaster />
      <AuthProvider>
        <QueryClientProvider client={queryClient}>
          <Router>
            <Switch>
              <LoggedOutRoute path="/login">
                <Login />
              </LoggedOutRoute>
              <LoggedInRoute path="/">
                <App />
              </LoggedInRoute>
              <DefaultRedirect />
            </Switch>
          </Router>
        </QueryClientProvider>
      </AuthProvider>
    </>
  );
}

ReactDOM.render(
  <React.StrictMode>
    <Root />
  </React.StrictMode>,
  document.getElementById("root")
);
