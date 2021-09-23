import React from "react";
import {Route} from "react-router-dom";

function NotFound() {
    return (
        <div className="flex flex-col pt-16 pb-12 bg-white">
            <main className="flex flex-col justify-center flex-grow w-full px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
                <div className="flex justify-center flex-shrink-0">
                    <a href="/" className="inline-flex">
                        <span className="sr-only">Calcio</span>
                        <img
                            className="w-auto h-12"
                            src="/images/logo.svg"
                            alt=""
                        />
                    </a>
                </div>
                <div className="py-16">
                    <div className="text-center">
                        <h1 className="mt-2 text-4xl font-extrabold tracking-tight text-gray-900 sm:text-5xl">Not found</h1>
                        <p className="mt-2 text-base text-gray-500">Sorry</p>
                        <div className="mt-6">
                            <a href="/" className="text-base font-medium text-blue-600 hover:text-blue-500">
                                Home<span aria-hidden="true"> &rarr;</span>
                            </a>
                        </div>
                    </div>
                </div>
            </main>
        </div>
    );
}

export function DefaultRedirect() {
    return <Route path="*" component={NotFound}/>;
}

export default NotFound;
