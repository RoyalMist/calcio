import {Menu, Popover, Transition} from "@headlessui/react";
import {MenuIcon, XIcon} from "@heroicons/react/outline";
import {ChartBarIcon, FireIcon, UserGroupIcon, UserIcon} from "@heroicons/react/solid";
import React, {Fragment} from "react";
import {Link, Switch, useLocation} from "react-router-dom";
import LoggedInRoute from "../components/LoggedInRoute";
import {DASHBOARDS, GAMES, HOME, LOGOUT, TEAMS, USERS} from "../routes";
import {classNames} from "../utils/classes";
import {DefaultRedirect} from "./404";
import Dashboards from "./dashboards";
import Home from "./home";
import Logout from "./logout";
import Users from "./users";
import logo from "../../images/logo.webp";
import {useAuth} from "../stores/authentication";
import Games from "./games";
import Teams from "./teams";

function App() {
    const path = useLocation().pathname;
    const {isAdmin, paseto} = useAuth();
    const ACCOUNT_NAVIGATION = [
        {name: "Sign out", to: LOGOUT},
    ];

    const MAIN_NAVIGATION = [
        {
            name: "Dashboard",
            to: DASHBOARDS,
            accessible: true,
            icon: ChartBarIcon,
            current: path.startsWith(DASHBOARDS),
        },
        {
            name: "Games",
            to: GAMES,
            accessible: true,
            icon: FireIcon,
            current: path.startsWith(GAMES),
        },
        {
            name: "Teams",
            to: TEAMS,
            accessible: true,
            icon: UserGroupIcon,
            current: path.startsWith(TEAMS),
        },
        {
            name: "Users",
            to: USERS,
            accessible: isAdmin(),
            icon: UserIcon,
            current: path.startsWith(USERS),
        },
    ];

    return (
        <div className="min-h-screen bg-gray-100">
            <Popover
                as="header"
                className={({open}) =>
                    classNames(
                        open ? "fixed inset-0 z-40 overflow-y-auto" : "",
                        "bg-white shadow-sm lg:static lg:overflow-y-visible"
                    )
                }
            >
                {({open}) => (
                    <>
                        <div className="px-4 mx-auto sm:px-6 lg:px-8 bg-blue-800">
                            <div className="relative flex justify-between xl:grid xl:grid-cols-12 lg:gap-8">
                                <div className="flex md:absolute md:left-0 md:inset-y-0 lg:static xl:col-span-2">
                                    <div className="flex items-center flex-shrink-0">
                                        <Link to={HOME}>
                                            <img
                                                className="block w-auto h-8"
                                                src={logo}
                                                alt="Calcio"
                                            />
                                        </Link>
                                    </div>
                                </div>
                                <div className="flex-1 min-w-0 col-span-6 md:px-8 lg:px-0 h-16"/>
                                <div className="flex items-center md:absolute md:right-0 md:inset-y-0 lg:hidden">
                                    <Popover.Button
                                        className="inline-flex items-center justify-center p-2 -mx-2 text-gray-400 rounded-md hover:bg-gray-100 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-blue-500">
                                        {open ? (
                                            <XIcon className="block w-6 h-6"/>
                                        ) : (
                                            <MenuIcon className="block w-6 h-6"/>
                                        )}
                                    </Popover.Button>
                                </div>
                                <div className="hidden lg:flex lg:items-center lg:justify-end xl:col-span-4">
                                    <Menu as="div" className="relative flex-shrink-0 ml-5">
                                        {({open}) => (
                                            <>
                                                <div>
                                                    <Menu.Button className="flex bg-white rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
                                                        <span className="inline-block h-10 w-10 rounded-full overflow-hidden bg-gray-100">
                                                            <svg className="h-full w-full text-gray-300" fill="currentColor" viewBox="0 0 24 24">
                                                                <path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z"/>
                                                            </svg>
                                                        </span>
                                                    </Menu.Button>
                                                </div>
                                                <Transition
                                                    show={open}
                                                    as={Fragment}
                                                    enter="transition ease-out duration-100"
                                                    enterFrom="transform opacity-0 scale-95"
                                                    enterTo="transform opacity-100 scale-100"
                                                    leave="transition ease-in duration-75"
                                                    leaveFrom="transform opacity-100 scale-100"
                                                    leaveTo="transform opacity-0 scale-95"
                                                >
                                                    <Menu.Items
                                                        static
                                                        className="absolute right-0 z-10 w-48 py-1 mt-2 origin-top-right bg-white rounded-md shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
                                                    >
                                                        {ACCOUNT_NAVIGATION.map((item) => (
                                                            <Menu.Item key={item.name}>
                                                                {({active}) => (
                                                                    <Link
                                                                        to={item.to}
                                                                        className={classNames(
                                                                            active
                                                                                ? "bg-blue-100 text-blue-800"
                                                                                : "",
                                                                            "block py-2 px-4 text-sm text-gray-700"
                                                                        )}
                                                                    >
                                                                        {item.name}
                                                                    </Link>
                                                                )}
                                                            </Menu.Item>
                                                        ))}
                                                    </Menu.Items>
                                                </Transition>
                                            </>
                                        )}
                                    </Menu>
                                </div>
                            </div>
                        </div>

                        <Popover.Panel as="nav" className="lg:hidden">
                            <div className="max-w-3xl px-2 pt-2 pb-3 mx-auto space-y-1 sm:px-4">
                                {MAIN_NAVIGATION.filter(n => n.accessible).map((item) => (
                                    <Link
                                        key={item.name}
                                        to={item.to}
                                        className={classNames(
                                            item.current
                                                ? "bg-blue-100 text-blue-800"
                                                : "hover:bg-blue-50",
                                            "block rounded-md py-2 px-3 text-base font-medium"
                                        )}
                                    >
                                        {item.name}
                                    </Link>
                                ))}
                            </div>
                            <div className="pt-4 pb-3 border-t border-gray-200">
                                <div className="flex items-center max-w-3xl px-4 mx-auto sm:px-6">
                                    <div className="flex-shrink-0">
                                        <span className="inline-block h-10 w-10 rounded-full overflow-hidden bg-gray-100">
                                            <svg className="h-full w-full text-gray-300" fill="currentColor" viewBox="0 0 24 24">
                                                <path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z"/>
                                            </svg>
                                        </span>
                                    </div>
                                    <div className="ml-3">
                                        <div className="text-base font-medium text-gray-800">
                                            {paseto?.user_name}
                                        </div>
                                    </div>
                                </div>
                                <div className="max-w-3xl px-2 mx-auto mt-3 space-y-1 sm:px-4">
                                    {ACCOUNT_NAVIGATION.map((item) => (
                                        <Link
                                            key={item.name}
                                            to={item.to}
                                            className="block px-3 py-2 text-base font-medium text-gray-500 rounded-md hover:bg-blue-50 hover:text-gray-900"
                                        >
                                            {item.name}
                                        </Link>
                                    ))}
                                </div>
                            </div>
                        </Popover.Panel>
                    </>
                )}
            </Popover>

            <div className="py-5">
                <div className="w-screen mx-auto sm:px-6 lg:px-8 lg:grid lg:grid-cols-12 lg:gap-8">
                    <div className="hidden lg:block lg:col-span-3 xl:col-span-2">
                        <nav className="sticky divide-y divide-gray-300 top-4">
                            <div className="pb-8 space-y-1">
                                {MAIN_NAVIGATION.filter(n => n.accessible).map((item) => (
                                    <Link
                                        key={item.name}
                                        to={item.to}
                                        className={classNames(
                                            item.current
                                                ? "bg-blue-200 text-blue-800"
                                                : "text-gray-600 hover:bg-blue-100",
                                            "group flex items-center px-3 py-2 text-sm font-medium rounded-md"
                                        )}
                                    >
                                        <item.icon
                                            className={classNames(
                                                item.current
                                                    ? "text-blue-800"
                                                    : "text-gray-400 group-hover:text-gray-600",
                                                "flex-shrink-0 -ml-1 mr-3 h-6 w-6"
                                            )}
                                        />
                                        <span className="truncate">{item.name}</span>
                                    </Link>
                                ))}
                            </div>
                        </nav>
                    </div>
                    <main className="w-full md:col-span-6 lg:col-span-9 xl:col-span-10">
                        <div className="px-4 sm:px-0">
                            <div className="p-5 bg-white rounded-lg shadow-2xl">
                                <Switch>
                                    <LoggedInRoute exact path={HOME}>
                                        <Home/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={LOGOUT}>
                                        <Logout/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={DASHBOARDS}>
                                        <Dashboards/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={GAMES}>
                                        <Games/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={TEAMS}>
                                        <Teams/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={USERS} mustBeAdmin={true}>
                                        <Users/>
                                    </LoggedInRoute>
                                    <DefaultRedirect/>
                                </Switch>
                            </div>
                        </div>
                    </main>
                </div>
            </div>
        </div>
    );
}

export default App;
