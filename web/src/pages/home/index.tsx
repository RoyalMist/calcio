import React from "react";
import {Redirect} from "react-router-dom";
import {DASHBOARDS, PROFILE} from "../../routes";
import {useAuth} from "../../stores/authentication";

function Home() {
    const {authState} = useAuth();
    if (authState.paseto?.user_id) {
        return <Redirect to={DASHBOARDS}/>;
    } else {
        return <Redirect to={PROFILE}/>;
    }
}

export default Home;
