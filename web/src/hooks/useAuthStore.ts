import create from "zustand";
import {Base64} from "js-base64";

type State = {
    token: string;
    setToken: (token: string) => void;
    init: () => void;
    clearStore: () => void;
    readToken: () => Paseto | null;
    isLoggedIn: () => boolean;
    isAdmin: () => boolean;
};

interface Paseto {
    exp: string;
    user_id: string;
    is_admin: boolean;
}

const useAuthStore = create<State>((set, get) => ({
    token: "",
    init: () => {
        const storedToken = localStorage.getItem("token");
        if (storedToken) {
            set({token: storedToken});
        }
    },
    setToken: (token) => {
        set({token});
        localStorage.setItem("token", token);
    },
    clearStore: () => {
        localStorage.removeItem("token");
        set({token: ""});
    },
    readToken: () => {
        let paseto: Paseto;
        try {
            const decoded = Base64.decode(get().token);
            paseto = JSON.parse(decoded.substring(decoded.indexOf("{"), decoded.indexOf("}") + 1));
        } catch (err) {
            return null;
        }

        const isValid = Date.parse(paseto.exp) > Date.now();
        if (!isValid) {
            localStorage.removeItem("token");
            set({token: ""});
            return null;
        }

        return paseto;
    },
    isLoggedIn: () => {
        return get().readToken() !== null
    },
    isAdmin: () => {
        let token = get().readToken();
        if (token === null) {
            return false;
        }

        return token.is_admin;
    },
}));

useAuthStore.getState().init();

export default useAuthStore;
