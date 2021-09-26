import React, {useEffect} from "react";
import {AuthActionKind, useAuth} from "../../stores/authentication";

function Logout() {
    const {authDispatch} = useAuth();

    useEffect(() => {
        authDispatch({type: AuthActionKind.CLEAR, token: null});
    }, [authDispatch]);

    return <></>;
}

export default Logout;
