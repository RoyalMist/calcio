import React from "react";
import Split from "../../components/split";

interface BaseProps {
  title: string;
  children: React.ReactNode;
}

function Base(props: BaseProps) {
  return (
      <div className="flex items-center justify-center min-h-screen px-4 py-12 sm:px-6 lg:px-8">
        <div className="w-full max-w-md space-y-8">
          <div>
            <img
                className="w-auto h-12 mx-auto"
                src="/images/logo.svg"
                alt="Calcio"
            />
            <Split>{props.title}</Split>
          </div>
          {props.children}
        </div>
      </div>
  );
}

export default Base;
