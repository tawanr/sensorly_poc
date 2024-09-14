import { GraphData } from "../components/graph";

export interface TempData {
  temperature: number;
  station_id: number;
  recorded_at: string;
  id: number;
}

export async function fetchSensorData(): Promise<TempData[]> {
  const response = await fetch("http://localhost:8000/api/v1/sensors?station_id=1", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + localStorage.getItem("token"),
    },
  });
  const data = await response.json();
  console.log(data);
  return data?.data || [];
}

export function incrementTime(currentDate: Date, period: string): Date {
  const date = new Date(currentDate.getTime());
  switch (period) {
    case "second":
      date.setSeconds(date.getSeconds() + 5);
      break;
    case "minute":
      date.setMinutes(date.getMinutes() + 1);
      break;
    case "hour":
      date.setHours(date.getHours() + 1);
      break;
    case "day":
      date.setDate(date.getDate() + 1);
      break;
    case "week":
      date.setDate(date.getDate() + 7);
      break;
    case "month":
      date.setMonth(date.getMonth() + 1);
      break;
    case "year":
      date.setFullYear(date.getFullYear() + 1);
      break;
  }
  return date;
}

export function getInitialData(limit: number = 80, period: string = "day"): GraphData[] {
  const data: GraphData[] = [];
  let currentDate = new Date("2024-04-01");
  const entry: GraphData = {
    time: new Date(currentDate.getTime()),
    temperature: 33,
    ph: 7.5,
    nitrite: 0.5,
    chlorine: 0.5,
    do: 0.5,
    alkaline: 0.5,
  };
  for (let i = 0; i < limit; i++) {
    entry.temperature += Math.random() - 0.5;
    entry.ph = Math.abs(entry.ph + (Math.random() - 0.5) * 0.2);
    entry.nitrite = Math.abs(entry.nitrite + (Math.random() - 0.5) * 0.2);
    entry.chlorine = Math.abs(entry.chlorine + (Math.random() - 0.5) * 0.2);
    entry.do = Math.abs(entry.do + (Math.random() - 0.5) * 0.2);
    entry.alkaline = Math.abs(entry.alkaline + (Math.random() - 0.5) * 0.2);
    entry.time = incrementTime(currentDate, period);
    currentDate = new Date(entry.time.getTime());
    data.push({ ...entry });
  }
  return data;
}