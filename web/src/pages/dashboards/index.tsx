import React from "react";
import {Switch, useLocation} from "react-router-dom";
import LoggedInRoute from "../../components/route/logged_in";
import SectionMenu from "../../components/section-menu";
import {DefaultRedirect} from "../404";
import Versus from "./versus";
import Stats from "./stats";

const STATS = "/dashboards";
const VERSUS = "/dashboards/versus";

function Dashboards() {
    const path = useLocation().pathname;

    const tabs = [
        {name: "Stats", to: STATS, current: path === STATS},
        {
            name: "Versus",
            to: VERSUS,
            current: path === VERSUS,
        },
    ];

    return (
        <>
            <SectionMenu tabs={tabs}/>
            <Switch>
                <LoggedInRoute path={VERSUS}>
                    <Versus/>
                </LoggedInRoute>
                <LoggedInRoute path={STATS}>
                    <Stats/>
                </LoggedInRoute>
                <DefaultRedirect/>
            </Switch>
        </>
    );
}

export default Dashboards;
