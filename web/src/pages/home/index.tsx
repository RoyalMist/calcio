import React from "react";
import {Redirect} from "react-router-dom";
import {DASHBOARDS, GAMES} from "../../routes";
import {useAuth} from "../../stores/authentication";

function Home() {
    const {isAdmin} = useAuth();
    if (isAdmin()) {
        return <Redirect to={DASHBOARDS}/>;
    } else {
        return <Redirect to={GAMES}/>;
    }
}

export default Home;
