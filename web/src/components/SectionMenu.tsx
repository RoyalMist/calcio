import React from "react";
import {Link} from "react-router-dom";
import {classNames} from "../utils/classes";

interface MenuItem {
    name: string;
    to: string;
    current: boolean;
}

interface SectionMenuProps {
    tabs: MenuItem[];
}

function SectionMenu({tabs}: SectionMenuProps) {
    return (
        <div className="mb-5">
            <div className="sm:hidden">
                <select
                    id="question-tabs"
                    className="block w-full text-base font-medium text-gray-900 border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring-blue-500"
                    defaultValue={tabs.find((tab) => tab.current)?.name}
                >
                    {tabs.map((tab) => (
                        <option key={tab.name}>{tab.name}</option>
                    ))}
                </select>
            </div>
            <div className="hidden sm:block">
                <nav className="relative z-0 flex divide-x divide-gray-200 rounded-lg shadow">
                    {tabs.map((tab, tabIdx) => (
                        <Link
                            key={tab.name}
                            to={tab.to}
                            className={classNames(
                                tab.current
                                    ? "text-gray-900"
                                    : "text-gray-500 hover:text-gray-700",
                                tabIdx === 0 ? "rounded-l-lg" : "",
                                tabIdx === tabs.length - 1 ? "rounded-r-lg" : "",
                                "group relative min-w-0 flex-1 overflow-hidden bg-gray-50 py-4 px-6 text-sm font-medium text-center hover:bg-blue-50 focus:z-10 shadow-md"
                            )}
                        >
                            <span>{tab.name}</span>
                            <span
                                className={classNames(
                                    tab.current ? "bg-blue-500" : "bg-transparent",
                                    "absolute inset-x-0 bottom-0 h-0.5"
                                )}
                            />
                        </Link>
                    ))}
                </nav>
            </div>
        </div>
    );
}

export default SectionMenu;
