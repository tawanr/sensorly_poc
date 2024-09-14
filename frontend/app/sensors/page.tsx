"use client";

import React, { useEffect } from "react";
import { GraphData, LineGraph } from "../components/graph";
import { fetchSensorData } from "../lib/sensor-data";
import { redirect } from "next/navigation";

export default function Page() {
  // const title = "Temperature";
  // const description = "Sensor Data";
  // const [data, setData] = React.useState<GraphData[]>([]);

  // return <LineGraph title={title} description={description} data={data} />;
  redirect("/sensors/1");
}
