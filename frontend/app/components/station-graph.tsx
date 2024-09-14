"use client";

import { GraphData, LineGraph } from "@/app/components/graph";
import { getInitialData, incrementTime } from "@/app/lib/sensor-data";
import React, { useEffect } from "react";

export default function StationGraph({
  name,
  period,
}: {
  name: string;
  period: string;
}) {
  // const name: string = "Station " + params.station;
  // const [period, setPeriod] = React.useState<string>("day");
  const [data, setData] = React.useState<GraphData[]>(
    getInitialData(50, period)
  );
  const [description, setDescription] = React.useState<string>("Sensor Data");

  useEffect(() => {
    let time = 100000000;
    switch (period) {
      case "second":
        time = 5000;
        setDescription("Sensor Data - Every 5 Seconds");
        break;
      case "minute":
        time = 60000;
        setDescription("Sensor Data - Every Minute");
        break;
      case "hour":
        time = 3600000;
        setDescription("Sensor Data - Hourly");
        break;
      case "day":
        time = 86400000;
        setDescription("Sensor Data - Daily");
        break;
      case "week":
        time = 604800000;
        setDescription("Sensor Data - Weekly");
        break;
    }
    const interval = setInterval(() => {
      setData((prevState) => {
        const newState = [...prevState];
        newState.shift();
        const lastEntry = { ...newState[newState.length - 1] };
        lastEntry.temperature += Math.random() - 0.5;
        lastEntry.ph = Math.abs(lastEntry.ph + (Math.random() - 0.5) * 0.2);
        lastEntry.nitrite = Math.abs(
          lastEntry.nitrite + (Math.random() - 0.5) * 0.2
        );
        lastEntry.chlorine = Math.abs(
          lastEntry.chlorine + (Math.random() - 0.5) * 0.2
        );
        lastEntry.do = Math.abs(lastEntry.do + (Math.random() - 0.5) * 0.2);
        lastEntry.alkaline = Math.abs(
          lastEntry.alkaline + (Math.random() - 0.5) * 0.2
        );
        lastEntry.time = incrementTime(lastEntry.time, period);
        newState.push({ ...lastEntry });

        return newState;
      });
    }, time);
    return () => clearInterval(interval);
  }, []);

  return (
    <div>
      <LineGraph
        title={name}
        description={description}
        data={data}
        period={period}
      />
    </div>
  );
}
