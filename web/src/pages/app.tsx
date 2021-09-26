import {Menu, Popover, Transition} from "@headlessui/react";
import {MenuIcon, XIcon} from "@heroicons/react/outline";
import {ChartBarIcon, SearchIcon, UsersIcon} from "@heroicons/react/solid";
import React, {Fragment} from "react";
import {Link, Switch, useLocation} from "react-router-dom";
import LoggedInRoute from "../components/route/logged_in";
import {DASHBOARDS, HOME, LOGOUT, PLAYERS, PROFILE} from "../routes";
import {classNames} from "../utils/classes";
import {DefaultRedirect} from "./404";
import Dashboards from "./dashboards";
import Home from "./home";
import Logout from "./logout";
import Profile from "./profile";
import Players from "./users";
import avatar from "../../images/avatar.webp";
import logo from "../../images/logo.webp";
import {useAuth} from "../stores/authentication";

const user = {
    name: "Thibault Fouache",
    email: "thibault@Calcio.ch",
    imageUrl: avatar,
};

function App() {
    const path = useLocation().pathname;
    const {isAdmin} = useAuth();
    const ACCOUNT_NAVIGATION = [
        {name: "Your Profile", to: PROFILE},
        {name: "Sign out", to: LOGOUT},
    ];

    const MAIN_NAVIGATION = [
        {
            name: "Dashboard",
            to: DASHBOARDS,
            restricted: isAdmin(),
            icon: ChartBarIcon,
            current: path.startsWith(DASHBOARDS),
        },
        {
            name: "Players",
            to: PLAYERS,
            restricted: isAdmin(),
            icon: UsersIcon,
            current: path.startsWith(PLAYERS),
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
                        <div className="px-4 mx-auto sm:px-6 lg:px-8">
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
                                <div className="flex-1 min-w-0 col-span-6 md:px-8 lg:px-0">
                                    <div className="flex items-center px-6 py-4 md:max-w-lg md:mx-auto lg:max-w-4xl lg:mx-0 xl:px-0">
                                        <div className="w-full">
                                            <div className="relative">
                                                <div className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                                                    <SearchIcon className="w-5 h-5 text-gray-400"/>
                                                </div>
                                                <input
                                                    id="search"
                                                    name="search"
                                                    className="block w-full py-2 pl-10 pr-3 text-sm placeholder-gray-500 bg-white border border-gray-300 rounded-md focus:outline-none focus:text-gray-900 focus:placeholder-gray-400 focus:ring-1 focus:ring-blue-800 focus:border-blue-800 sm:text-sm"
                                                    placeholder="Search"
                                                    type="search"
                                                />
                                            </div>
                                        </div>
                                    </div>
                                </div>
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
                                                        <img
                                                            className="w-8 h-8 rounded-full"
                                                            src={user.imageUrl}
                                                            alt=""
                                                        />
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
                                {MAIN_NAVIGATION.filter(n => n.restricted).map((item) => (
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
                                        <img
                                            className="w-10 h-10 rounded-full"
                                            src={user.imageUrl}
                                            alt=""
                                        />
                                    </div>
                                    <div className="ml-3">
                                        <div className="text-base font-medium text-gray-800">
                                            {user.name}
                                        </div>
                                        <div className="text-sm font-medium text-gray-500">
                                            {user.email}
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
                                {MAIN_NAVIGATION.filter(n => n.restricted).map((item) => (
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
                                    <LoggedInRoute path={PROFILE}>
                                        <Profile/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={LOGOUT}>
                                        <Logout/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={DASHBOARDS} mustBeAdmin={true}>
                                        <Dashboards/>
                                    </LoggedInRoute>
                                    <LoggedInRoute path={PLAYERS} mustBeAdmin={true}>
                                        <Players/>
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
