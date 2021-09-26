import React from "react";
import {Redirect} from "react-router-dom";
import {DASHBOARDS, PROFILE} from "../../routes";
import {useAuth} from "../../stores/authentication";

function Home() {
    const {isAdmin} = useAuth();
    if (isAdmin()) {
        return <Redirect to={DASHBOARDS}/>;
    } else {
        return <Redirect to={PROFILE}/>;
    }
}

export default Home;
