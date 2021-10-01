export interface Paseto {
    exp: string;
    user_id: string;
    user_name: string,
    is_admin: boolean;
}

const parse = (token: string | null): Paseto | null => {
    if (token == null) {
        return null;
    }

    try {
        const decoded = window.atob(token.replaceAll(/[^A-Za-z0-9]|v[0-9]+\..+\./g, ""));
        return JSON.parse(decoded.substring(decoded.indexOf("{"), decoded.indexOf("}") + 1));
    } catch (err) {
        console.error(err);
        return null;
    }
}

export {parse};
