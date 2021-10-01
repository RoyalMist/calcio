import React from "react";
import SectionHeader from "../../components/section-header";
import {toast} from "react-hot-toast";

const Teams = () => {
    return (
        <>
            <SectionHeader action={() => toast.success("Hello")}>Teams</SectionHeader>
        </>

    );
};

export default Teams;
