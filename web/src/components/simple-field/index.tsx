import {Field} from "formik";
import React from "react";

interface SimpleFieldProps {
    children?: React.ReactNode;
    type: string;
    name: string;
    placeholder?: string;
    className?: string;
}

function SimpleField({children, name, className, placeholder, type}: SimpleFieldProps) {
    return (
        <div className={`space-y-1 ${className}`}>
            {!!children && (
                <label className="block mb-1 text-sm font-medium text-gray-700">
                    {children}
                </label>
            )}
            <Field
                className="block w-full px-3 py-3 placeholder-gray-400 border border-gray-300 rounded-md shadow-sm appearance-none focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                type={type}
                name={name}
                placeholder={placeholder}
            />
        </div>
    );
}

export default SimpleField;
