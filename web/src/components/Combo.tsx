import {Listbox, Transition} from "@headlessui/react";
import {ChevronDownIcon} from "@heroicons/react/solid";
import React, {Fragment, useState} from "react";
import {classNames} from "../utils/classes";

interface ComboListOption {
    label: string;
    value: string;
    current: boolean;
}

interface ComboListProps {
    options: ComboListOption[];
}

function ComboList({options}: ComboListProps) {
    const [selected, setSelected] = useState(options[0]);

    return (
        <Listbox value={selected} onChange={setSelected}>
            {({open}) => (
                <>
                    <div className="relative">
                        <div className="inline-flex divide-x divide-blue-800 rounded-md shadow-sm">
                            <div className="relative z-0 inline-flex divide-x divide-blue-800 rounded-md shadow-sm">
                                <div className="relative inline-flex items-center py-2 pl-3 pr-4 text-white bg-blue-800 border border-transparent shadow-sm rounded-l-md">
                                    <p className="ml-2.5 text-sm font-medium">{selected.label}</p>
                                </div>
                                <Listbox.Button
                                    className="relative inline-flex items-center p-2 text-sm font-medium text-white bg-blue-800 rounded-l-none rounded-r-md hover:bg-blue-600 focus:outline-none focus:z-10 focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-50 focus:ring-blue-500">
                                    <ChevronDownIcon
                                        className="w-5 h-5 text-white"
                                        aria-hidden="true"
                                    />
                                </Listbox.Button>
                            </div>
                        </div>

                        <Transition
                            show={open}
                            as={Fragment}
                            leave="transition ease-in duration-100"
                            leaveFrom="opacity-100"
                            leaveTo="opacity-0"
                        >
                            <Listbox.Options
                                static
                                className="absolute right-0 mt-2 overflow-hidden origin-top-right bg-white divide-y divide-gray-200 rounded-md shadow-lg w-72 ring-1 ring-black ring-opacity-5 focus:outline-none"
                            >
                                {options.map((option) => (
                                    <Listbox.Option
                                        key={option.label}
                                        className={({active}) =>
                                            classNames(
                                                active ? "text-white bg-blue-800" : "text-gray-900",
                                                "cursor-default select-none relative p-4 text-sm"
                                            )
                                        }
                                        value={option}
                                    >
                                        {({selected}) => (
                                            <div className="flex flex-col">
                                                <div className="flex justify-between">
                                                    <p
                                                        className={
                                                            selected ? "font-semibold" : "font-normal"
                                                        }
                                                    >
                                                        {option.label}
                                                    </p>
                                                </div>
                                            </div>
                                        )}
                                    </Listbox.Option>
                                ))}
                            </Listbox.Options>
                        </Transition>
                    </div>
                </>
            )}
        </Listbox>
    );
}

export default ComboList;
