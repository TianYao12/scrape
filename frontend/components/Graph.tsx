import { FireballData } from "@/lib/types";
import { useRef } from "react";

const Graph = (props: { data: FireballData[]}) => {
    const backgroundRef = useRef<HTMLDivElement>(null);
    const width = backgroundRef?.current?.clientWidth || 0;
    const height = backgroundRef?.current?.clientHeight || 0;
    const margin = 30;

    return (
        <div ref={backgroundRef} className="relative h-full cursor-crosshair">
            <svg width={width} height={height}>
                
            </svg>

        </div>
    )
}

export default Graph;