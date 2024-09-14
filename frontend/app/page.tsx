"use client";

import React, { useEffect } from "react";
import { Notification, Notifications } from "./components/notifications";
import StationCard, { StationData } from "./components/station-card";

export default function Home() {
  const notifications: Notification[] = [
    {
      title: "Temperature for Station 1 is slightly high.",
      description: "1 hour ago",
    },
  ];
  const [notificationsState, setNotificationsState] =
    React.useState(notifications);
  const readNoti = () => {
    setNotificationsState([]);
  };

  const [stationsData, setStationsData] = React.useState<StationData[]>([
    {
      name: "Station 1",
      temperature: 29,
      ph: 7.5,
      nitrite: 0.5,
      chlorine: 0.5,
      do: 0.5,
      alkaline: 0.5,
    },
    {
      name: "Station 2",
      temperature: 28,
      ph: 8.0,
      nitrite: 0.5,
      chlorine: 0.5,
      do: 0.5,
      alkaline: 0.5,
    },
    {
      name: "Station 3",
      temperature: 27.9,
      ph: 7.9,
      nitrite: 0.5,
      chlorine: 0.5,
      do: 0.5,
      alkaline: 0.5,
    },
  ]);
  useEffect(() => {
    const interval = setInterval(() => {
      setStationsData((prevState) => {
        const newState = [...prevState];
        newState.map((station) => {
          station.temperature += (Math.random() - 0.5) * 0.5;
          station.ph = Math.abs(station.ph + (Math.random() - 0.5) * 0.1);
          station.nitrite = Math.abs(
            station.nitrite + (Math.random() - 0.5) * 0.1
          );
          station.chlorine = Math.abs(
            station.chlorine + (Math.random() - 0.5) * 0.1
          );
          station.do = Math.abs(station.do + (Math.random() - 0.5) * 0.1);
          station.alkaline = Math.abs(
            station.alkaline + (Math.random() - 0.5) * 0.1
          );
          return station;
        });
        return newState;
      });
    }, 10000);
    return () => clearInterval(interval);
  }, []);

  const stationJSX = [];
  for (let i = 0; i < stationsData.length; i++) {
    stationJSX.push(
      <div id={"station-" + (i + 1)}>
        <a href={`/sensors/${i + 1}`}>
          <StationCard data={stationsData[i]} />
        </a>
      </div>
    );
  }

  return (
    <div className="flex flex-wrap flex-col items-center justify-center mt-4 gap-4">
      <div className="flex-1 items-center justify-center w-full">
        <Notifications notifications={notificationsState} readNoti={readNoti} />
      </div>
      <div className="lg:grid flex-1 items-center justify-center w-full lg:grid-cols-3 xl:grid-cols-3 gap-4 space-y-4">
        {stationJSX}
      </div>
    </div>
  );
}
