import * as React from "react";
import {useReducer} from "react";
import {parse, Paseto} from "../utils/paseto";

enum AuthActionKind {
    SET,
    CLEAR,
}

interface AuthAction {
    type: AuthActionKind;
    token?: string;
}

interface AuthState {
    token: string | null;
    paseto: Paseto | null;
}

interface AuthProps {
    children: React.ReactNode;
}

interface AuthContext {
    token: string;
    paseto: Paseto | null;
    isLoggedIn: () => boolean;
    isAdmin: () => boolean;
    dispatcher: React.Dispatch<AuthAction>
}

const EMPTY: AuthState = {token: null, paseto: null}
const AuthContext = React.createContext<AuthContext>({token: "", paseto: null, isLoggedIn: () => false, isAdmin: () => false, dispatcher: () => EMPTY});
const stateFromToken = (token: string | null): AuthState => {
    const paseto = parse(token);
    if (token != null && paseto != null && Date.parse(paseto.exp) > Date.now()) {
        localStorage.setItem("token", token);
        return {token, paseto};
    } else {
        localStorage.removeItem("token");
        return EMPTY;
    }
}

const authReducer = (state: AuthState, action: AuthAction): AuthState => {
    const {type, token} = action;
    switch (type) {
        case AuthActionKind.SET: {
            return stateFromToken(token === undefined ? null : token);
        }

        case AuthActionKind.CLEAR: {
            return stateFromToken(null);
        }

        default: {
            throw new Error(`Unknown action type: ${action.type}`);
        }
    }
}

function AuthProvider({children}: AuthProps) {
    const storedToken = localStorage.getItem("token");
    const [state, dispatch] = useReducer(authReducer, stateFromToken(storedToken));
    const isAdmin = () => !!stateFromToken(state.token).paseto?.is_admin;
    const isLoggedIn = () => !!stateFromToken(state.token).paseto;
    const value: AuthContext = {token: `Bearer ${state.token}`, paseto: state.paseto, dispatcher: dispatch, isAdmin, isLoggedIn};
    return (
        <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
    );
}

function useAuth() {
    const context = React.useContext(AuthContext);
    if (context === undefined) {
        throw new Error("useAuth must be used within a AuthProvider");
    }

    return context;
}

export {AuthActionKind, AuthProvider, useAuth};
