import React, {useEffect} from "react";
import {AuthActionKind, useAuth} from "../../stores/authentication";

function Logout() {
    const {dispatcher} = useAuth();

    useEffect(() => {
        dispatcher({type: AuthActionKind.CLEAR});
    }, [dispatcher]);

    return <></>;
}

export default Logout;
