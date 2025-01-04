"use client";

import { useState, useEffect } from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { FireballData } from "@/lib/types";
import FireballTable from "@/components/Table";
import { Table } from "lucide-react";
import Graph from "@/components/Graph";

export default function Home() {
  const [fireballData, setFireballData] = useState<FireballData[]>([]);
  const [isTable, setIsTable] = useState<boolean>(true);

  useEffect(() => {
    const fetchFireballData = async () => {
      try {
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/v1/fireballs`,
          {
            method: "GET",
            headers: {
              "Content-type": "application/json",
              Authorization: `${process.env.NEXT_PUBLIC_AUTH_KEY}`,
            },
          }
        );
        const data = await response.json();
        setFireballData(data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchFireballData();
  }, []);

  useEffect(() => console.log(fireballData));
  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <div>
        <Select
          onValueChange={(value) => {
            setIsTable(value == "table");
          }}
        >
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Display" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="graph">Graph</SelectItem>
            <SelectItem value="table">Table</SelectItem>
          </SelectContent>
        </Select>
      </div>
      <div>
        {isTable ? (
          <Graph data={fireballData} />
        ) : (
          <FireballTable data={fireballData} />
        )}
      </div>
    </div>
  );
}
