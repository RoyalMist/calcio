import React from "react";
import Loader from "react-loader-spinner";

interface SpinnerProps {
    loading: boolean;
}

function Spinner({loading}: SpinnerProps) {
    let spinner;

    if (loading) {
        spinner = (
            <div className="fixed top-0 left-0 z-50 block w-full h-full bg-gray-800 opacity-80">
        <span
            className="relative block w-0 h-0 mx-auto my-0 opacity-100"
            style={{top: "40%"}}
        >
          <Loader type="Watch" color="#fff" height={250} width={250}/>;
        </span>
            </div>
        );
    } else {
        spinner = <></>;
    }

    return spinner;
}

export default Spinner;
