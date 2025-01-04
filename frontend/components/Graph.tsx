import { Line } from "@visx/shape";
import { FireballData } from "@/lib/types";
import { useMemo, useRef } from "react";
import Axis from "./Axis";
import { scaleLinear, scaleBand } from "@visx/scale";

const Graph = (props: { data: FireballData[] }) => {
  const backgroundRef = useRef<HTMLDivElement>(null);
  const width = backgroundRef?.current?.clientWidth || 0;
  const height = backgroundRef?.current?.clientHeight || 0;
  const margin = 30;

  const xScale = useMemo(
    () =>
      scaleBand<string>({
        range: [0, width - margin],
        domain: props.data.map((fireball) => fireball.id),
        paddingOuter: 1,
      }),
    [width, props.data]
  );

  const yScale = useMemo(
    () =>
      scaleLinear<number>({
        range: [height - margin, margin / 2],
        domain: [
          Math.min(
            ...props.data.map((fireball) => fireball.total_radiated_energy)
          ),
          Math.max(
            ...props.data.map((fireball) => fireball.total_radiated_energy)
          ),
        ],
        round: true,
        nice: true,
      }),
    [width, props.data]
  );

  const chartData = useMemo(() => {
    return props.data.map((fireball) => ({
      x: (xScale(fireball.id) ?? 0) + xScale.bandwidth() / 2,
      y: yScale(fireball.total_radiated_energy),
    }));
  }, [props.data, xScale, yScale]);

  return (
    <div
      ref={backgroundRef}
      className="relative w-[500px] h-[500px] cursor-crosshair"
    >
      <svg width={width} height={height}>
        <rect x={0} y={0} width={width} height={height} fill="black" rx={10} />
        {chartData.map((point, index) => {
          if (index === 0) return null;
          const previousPoint = chartData[index - 1];
          return (
            <Line
              key={index}
              x1={previousPoint.x}
              y1={previousPoint.y}
              x2={point.x}
              y2={point.y}
              stroke="red"
              strokeWidth={2}
            />
          );
        })}
        {chartData?.map((data, i) => (
          <circle key={i} cx={data.x} cy={data.y} r={3} fill="red" />
        ))}
        <Axis
          width={width}
          height={height}
          margin={margin}
          xScale={xScale}
          yScale={yScale}
        />
      </svg>
    </div>
  );
};

export default Graph;
