import {decode} from "js-base64";

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
        const decoded = decode(token);
        return JSON.parse(decoded.substring(decoded.indexOf("{"), decoded.indexOf("}") + 1));
    } catch (err) {
        console.error(err);
        return null;
    }
}

export {parse};
