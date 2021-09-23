import React from "react";

interface SplitProps {
    children: React.ReactNode;
}

function Split({children}: SplitProps) {
    return (
        <div className="relative mt-6">
            <div className="absolute inset-0 flex items-center" aria-hidden="true">
                <div className="w-full border-t border-gray-300"/>
            </div>
            <div className="relative flex justify-center text-sm">
                <span className="px-2 text-gray-500 bg-white">{children}</span>
            </div>
        </div>
    );
}

export default Split;
