import React from "react";
import {Redirect} from "react-router-dom";
import useAuthStore from "../../hooks/useAuthStore";
import {DASHBOARDS, PROFILE} from "../../routes";

function Home() {
    const auth = useAuthStore();
    if (auth.isAdmin()) {
        return <Redirect to={DASHBOARDS}/>;
    } else {
        return <Redirect to={PROFILE}/>;
    }
}

export default Home;
