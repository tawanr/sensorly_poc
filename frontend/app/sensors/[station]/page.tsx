import { File, ListFilter } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import StationGraph from "@/app/components/station-graph";

export default function Page({ params }: { params: { station: number } }) {
  const name: string = "Station " + params.station.toString();

  return (
    <Tabs defaultValue="second">
      <div className="flex items-center">
        <TabsList>
          <TabsTrigger value="second">Second</TabsTrigger>
          <TabsTrigger value="minute">Minute</TabsTrigger>
          <TabsTrigger value="hour">Hour</TabsTrigger>
          <TabsTrigger value="day">Day</TabsTrigger>
          <TabsTrigger value="week">Week</TabsTrigger>
        </TabsList>
        <div className="ml-auto flex items-center gap-2">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="outline" size="sm" className="h-7 gap-1 text-sm">
                <ListFilter className="h-3.5 w-3.5" />
                <span className="sr-only sm:not-sr-only">Filter</span>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>Filter by</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuCheckboxItem checked>
                Alerted
              </DropdownMenuCheckboxItem>
              <DropdownMenuCheckboxItem>Date</DropdownMenuCheckboxItem>
            </DropdownMenuContent>
          </DropdownMenu>
          <Button size="sm" variant="outline" className="h-7 gap-1 text-sm">
            <File className="h-3.5 w-3.5" />
            <span className="sr-only sm:not-sr-only">Export</span>
          </Button>
        </div>
      </div>
      <TabsContent value="second">
        <StationGraph name={name} period="second" />
      </TabsContent>
      <TabsContent value="minute">
        <StationGraph name={name} period="minute" />
      </TabsContent>
      <TabsContent value="hour">
        <StationGraph name={name} period="hour" />
      </TabsContent>
      <TabsContent value="day">
        <StationGraph name={name} period="day" />
      </TabsContent>
      <TabsContent value="week">
        <StationGraph name={name} period="week" />
      </TabsContent>
    </Tabs>
  );
}
