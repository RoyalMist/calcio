import React, {useEffect} from "react";
import useAuthStore from "../../hooks/useAuthStore";

function Logout() {
    const authStore = useAuthStore();

    useEffect(() => {
        authStore.clearStore();
    }, [authStore]);

    return <></>;
}

export default Logout;
