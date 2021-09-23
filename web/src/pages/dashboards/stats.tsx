import React from "react";
import {Bar} from "react-chartjs-2";

interface StatsProps {
}

function Stats(props: StatsProps) {
    const data = {
        labels: [
            "January",
            "February",
            "March",
            "April",
            "May",
            "June",
            "July",
            "August",
            "September",
            "October",
            "November",
            "December",
        ],
        datasets: [
            {
                label: "Total goals",
                data: [120, 190, 450, 500, 400, 238, 690, 800, 150, 300, 678, 890],
                backgroundColor: ["rgba(255, 159, 64, 0.2)"],
                borderColor: ["rgba(255, 159, 64, 1)"],
                borderWidth: 1,
            },
            {
                label: "Total players",
                data: [60, 75, 200, 300, 235, 120, 420, 535, 70, 127, 430, 600],
                backgroundColor: "rgba(54, 162, 235, 0.2)",
                borderColor: ["rgba(54, 162, 235, 1)"],
                borderWidth: 1,
            },
            {
                label: "Total games",
                data: [30, 50, 125, 175, 120, 45, 235, 356, 25, 68, 345, 450],
                backgroundColor: "rgba(153, 102, 255, 0.2)",
                borderColor: ["rgba(153, 102, 255, 1)"],
                borderWidth: 1,
            },
        ],
    };

    return (
        <div>
            <Bar data={data}/>
        </div>
    );
}

export default Stats;
