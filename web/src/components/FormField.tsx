import React from "react";
import {ExclamationCircleIcon} from "@heroicons/react/solid";
import {FieldError, UseFormRegisterReturn} from "react-hook-form";

interface FormFieldProps {
    children: React.ReactNode;
    type?: string;
    error?: FieldError;
    errorMessage?: string;
    placeholder?: string;
    defaultValue?: string;
    form: UseFormRegisterReturn;
}

const FormField = ({children, type = "text", error, errorMessage = "Error !", placeholder = "", defaultValue = "", form}: FormFieldProps) => {
    return (
        <>
            <div className={`relative border rounded-md px-3 py-3 shadow-sm ${error != undefined ? "border-red-600" : "border-blue-600 focus-within:ring-1"}`}>
                <label
                    className={`absolute -top-2 left-2 -mt-px inline-block px-1 bg-white text-xs font-medium ${error != undefined ? "focus-within:ring-red-600 focus-within:border-red-600 text-red-600" : "focus-within:ring-blue-600 focus-within:border-blue-600 text-gray-900"}`}>
                    {children}
                </label>
                <input
                    type={type}
                    placeholder={placeholder}
                    defaultValue={defaultValue}
                    className={`block w-full border-0 p-0 focus:ring-0 sm:text-sm ${error != undefined ? "text-red-600 placeholder-red-300" : "text-gray-900 placeholder-gray-300"}`}
                    {...form}
                />
                {!!error && <div className="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
                    <ExclamationCircleIcon className="h-5 w-5 text-red-500" aria-hidden="true"/>
                </div>}
            </div>
            {!!error && <p className="text-sm text-red-600">
                {errorMessage}
            </p>}
        </>
    );
};

export default FormField;
