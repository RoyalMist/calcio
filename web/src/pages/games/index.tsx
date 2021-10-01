import React from "react";
import {toast} from "react-hot-toast";
import SectionHeader from "../../components/section-header";

const Games = () => {
    return (
        <>
            <SectionHeader action={() => toast.success("Hello")}>Games</SectionHeader>
        </>
    );
};

export default Games;
