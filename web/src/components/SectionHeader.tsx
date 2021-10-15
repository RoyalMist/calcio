import {PlusCircleIcon} from '@heroicons/react/outline';
import React from "react";

interface SectionHeaderProps {
    children: React.ReactNode;
    action?: () => void;
}

function SectionHeader({children, action}: SectionHeaderProps) {
    return (
        <div className="flex items-center justify-between mb-10">
            <div className="flex-1 min-w-0">
                <h2 className="text-2xl font-bold leading-7 text-gray-900 sm:text-3xl sm:truncate">{children}</h2>
            </div>
            {action !== undefined && <div className="flex mt-4 md:mt-0 md:ml-4">
                <button
                    type="button"
                    onClick={action}
                    className="btn btn-primary"
                >
                    <PlusCircleIcon className="w-6 h-6"/>
                </button>
            </div>}
        </div>
    )
}

export default SectionHeader;
