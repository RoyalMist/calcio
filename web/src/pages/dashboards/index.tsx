import React from "react";
import {Switch, useLocation} from "react-router-dom";
import LoggedInRoute from "../../components/route/logged_in";
import SectionMenu from "../../components/section-menu";
import {DefaultRedirect} from "../404";
import Players from "./players";
import Teams from "./teams";
import Stats from "./stats";

const STATS = "/dashboards";
const BEST_TEAMS = "/dashboards/teams";
const BEST_PLAYERS = "/dashboards/players";

function Dashboards() {
    const path = useLocation().pathname;

    const tabs = [
        {name: "Stats", to: STATS, current: path === STATS},
        {
            name: "Best Teams",
            to: BEST_TEAMS,
            current: path === BEST_TEAMS,
        },
        {
            name: "Best Players",
            to: BEST_PLAYERS,
            current: path === BEST_PLAYERS,
        },
    ];

    return (
        <>
            <SectionMenu tabs={tabs}/>
            <Switch>
                <LoggedInRoute path={BEST_TEAMS}>
                    <Teams/>
                </LoggedInRoute>
                <LoggedInRoute path={BEST_PLAYERS}>
                    <Players/>
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
