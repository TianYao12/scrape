"use client";

import { useState, useEffect } from "react";

interface FireballData {
  id: string;
  created_at: string;
  feed_id: string;
  updated_at: string;
  total_radiated_energy: number
}

export default function Home() {
  const [fireballData, setFireballData] = useState<FireballData[]>([]);

  useEffect(() => {
    const fetchFireballData = async () => {
      try {
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/v1/fireballs`,
          {
            method: "GET",
            headers: {
              "Content-type": "application/json",
              "Authorization": `${process.env.NEXT_PUBLIC_AUTH_KEY}`,
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

  useEffect(()=> console.log(fireballData))
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      Fireball
      <div>{fireballData.map((fireball, index) => <div key={`fireball-${index}`}>123</div>)}</div>
    </div>
  );
}
