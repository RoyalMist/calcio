import React from "react";
import {Link} from "react-router-dom";

interface GoProps {
    to: string;
    children: React.ReactNode;
    className?: string;
}

function Go({className, children, to}: GoProps) {
    return (
        <Link
            to={to}
            className={`font-medium text-blue-800 hover:text-blue-600 ${className}`}
        >
            {children}
        </Link>
    );
}

export default Go;
