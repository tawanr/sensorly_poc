import { Bar, BarChart, LabelList, XAxis, YAxis } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";

export interface StationData {
  name: string;
  temperature: number;
  ph: number;
  nitrite: number;
  chlorine: number;
  do: number;
  alkaline: number;
}

export default function StationCard(props: { data: StationData }) {
  return (
    <Card className="w-full" x-chunk="charts-01-chunk-4">
      <CardContent className="flex gap-4 p-4 pb-2">
        <CardTitle>{props.data.name}</CardTitle>
      </CardContent>
      <CardFooter className="flex flex-row border-t p-4">
        <div className="flex w-full items-center gap-2">
          <div className="grid flex-1 auto-rows-min gap-0.5">
            <div className="text-xs text-muted-foreground">Temperature</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              {props.data.temperature.toFixed(2)}
              <span className="text-sm font-normal text-muted-foreground"></span>
            </div>
          </div>
          <Separator orientation="vertical" className="mx-2 h-10 w-px" />
          <div className="grid flex-1 auto-rows-min gap-0.5">
            <div className="text-xs text-muted-foreground">pH</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              {props.data.ph.toFixed(1)}
              <span className="text-sm font-normal text-muted-foreground"></span>
            </div>
          </div>
          <Separator orientation="vertical" className="mx-2 h-10 w-px" />
          <div className="grid flex-1 auto-rows-min gap-0.5">
            <div className="text-xs text-muted-foreground">Nitrite</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              {props.data.nitrite.toFixed(1)}
              <span className="text-sm font-normal text-muted-foreground"></span>
            </div>
          </div>
        </div>
      </CardFooter>
      <CardFooter className="flex flex-row border-t p-4">
        <div className="flex w-full items-center gap-2">
          <div className="grid flex-1 auto-rows-min gap-0.5">
            <div className="text-xs text-muted-foreground">Chlorine</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              {props.data.chlorine.toFixed(1)}
              <span className="text-sm font-normal text-muted-foreground"></span>
            </div>
          </div>
          <Separator orientation="vertical" className="mx-2 h-10 w-px" />
          <div className="grid flex-1 auto-rows-min gap-0.5">
            <div className="text-xs text-muted-foreground">DO</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              {props.data.do.toFixed(1)}
              <span className="text-sm font-normal text-muted-foreground"></span>
            </div>
          </div>
          <Separator orientation="vertical" className="mx-2 h-10 w-px" />
          <div className="grid flex-1 auto-rows-min gap-0.5">
            <div className="text-xs text-muted-foreground">Alkaline</div>
            <div className="flex items-baseline gap-1 text-2xl font-bold tabular-nums leading-none">
              {props.data.alkaline.toFixed(1)}
              <span className="text-sm font-normal text-muted-foreground"></span>
            </div>
          </div>
        </div>
      </CardFooter>
    </Card>
  );
}
