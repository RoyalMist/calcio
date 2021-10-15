import {Dialog, Transition} from '@headlessui/react';
import {XIcon} from '@heroicons/react/outline';
import React, {Fragment} from 'react';

interface SlideOverProps {
    open: boolean;
    close: () => void;
    title: string;
    subTitle?: string;
    children?: React.ReactNode;
}

function SlideOver({open, close, title, subTitle, children}: SlideOverProps) {
    return (
        <Transition.Root show={open} as={Fragment}>
            <Dialog as="div" static className="fixed inset-0 overflow-hidden" open={open} onClose={close}>
                <div className="absolute inset-0 overflow-hidden">
                    <Dialog.Overlay className="absolute inset-0"/>
                    <div className="fixed inset-y-0 right-0 flex max-w-full pl-10">
                        <Transition.Child
                            as={Fragment}
                            enter="transform transition ease-in-out duration-500 sm:duration-700"
                            enterFrom="translate-x-full"
                            enterTo="translate-x-0"
                            leave="transform transition ease-in-out duration-500 sm:duration-700"
                            leaveFrom="translate-x-0"
                            leaveTo="translate-x-full"
                        >
                            <div className="w-screen max-w-md">
                                <div className="flex flex-col h-full overflow-y-scroll bg-white shadow-xl">
                                    <div className="px-4 py-6 bg-blue-700 sm:px-6">
                                        <div className="flex items-center justify-between">
                                            <Dialog.Title className="text-lg font-medium text-white">{title}</Dialog.Title>
                                            <div className="flex items-center ml-3 h-7">
                                                <button
                                                    className="text-blue-200 bg-blue-700 rounded-md hover:text-white focus:outline-none focus:ring-2 focus:ring-white"
                                                    onClick={close}
                                                >
                                                    <XIcon className="w-6 h-6" aria-hidden="true"/>
                                                </button>
                                            </div>
                                        </div>
                                        {subTitle !== undefined &&
                                        <div className="mt-1">
                                            <p className="text-sm text-blue-300">
                                                {subTitle}
                                            </p>
                                        </div>
                                        }
                                    </div>
                                    <div className="relative flex flex-col items-center justify-start h-full px-4 py-10 sm:px-6">
                                        {children}
                                    </div>
                                </div>
                            </div>
                        </Transition.Child>
                    </div>
                </div>
            </Dialog>
        </Transition.Root>
    )
}

export default SlideOver;
