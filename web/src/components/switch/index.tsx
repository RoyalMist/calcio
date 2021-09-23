import {Switch as HdSwitch} from "@headlessui/react";
import React from "react";
import {classNames} from "../../utils/classes";

interface SwitchProps {
    title: string;
    value: boolean;
    update: (value: boolean) => void;
    children?: React.ReactNode;
}

function Switch({title, value, update, children}: SwitchProps) {
    return (
        <HdSwitch.Group as="div" className="flex items-center justify-between py-4">
            <div className="flex flex-col">
                <HdSwitch.Label
                    as="p"
                    className="text-sm font-medium text-gray-900"
                    passive
                >
                    {title}
                </HdSwitch.Label>
                <HdSwitch.Description className="text-sm text-gray-500">
                    {children}
                </HdSwitch.Description>
            </div>
            <HdSwitch
                checked={value}
                onChange={update}
                className={classNames(value ? "bg-blue-800" : "bg-gray-200",
                    "ml-4 relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none"
                )}
            >
        <span
            className={classNames(
                value ? "translate-x-5" : "translate-x-0",
                "inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200"
            )}
        />
            </HdSwitch>
        </HdSwitch.Group>
    );
}

export default Switch;
