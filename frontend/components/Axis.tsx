import { AxisBottom, AxisRight } from "@visx/axis";
import { ScaleBand, ScaleLinear} from "d3-scale";

const Axis = ({
  width,
  height,
  margin,
  xScale,
  yScale,
}: {
  width: number;
  height: number;
  margin: number;
  xScale: ScaleBand<string>;
  yScale: ScaleLinear<number, number, never>;

}) => {
  return (
    <>
      <AxisBottom
        top={height - margin}
        scale={xScale}
        stroke={"#545353"}
        tickStroke={"#545353"}
        tickLabelProps={{
          fontSize: 12,
          fill: "#545353",
          textAnchor: "middle",
          fontFamily: "inherit",
        }}
      />
      <AxisRight
        left={width - margin}
        scale={yScale}
        stroke={"#545353"}
        tickStroke={"#545353"}
        tickLabelProps={{
          fontSize: 12,
          fill: "#545353",
          textAnchor: "middle",
          fontFamily: "inherit",
          dx: 18,
        }}
      />
    </>
  );
};

export default Axis;
