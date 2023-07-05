import { useState } from "react";
import "./canvas.css";
import Pallet from "./Pallet";
import Renderer from "./renderer/Renderer";
import Syncer from "./syncer/Syncer";

export default function Canvas(){

    const start: Drawing = { draw_objs: []}

    const [drawing, setDrawing] = useState(start);
    const [palleton, setPalletOn] = useState(false);

    function pallet(e: React.MouseEvent<HTMLDivElement, MouseEvent>) {
        e.preventDefault();
        setPalletOn(true);
        }

    return (
            <div className="canvas" onContextMenu={(e) => pallet(e)}>
                {palleton && <Pallet pitems={[]}/>}
                <Renderer drawing={drawing}/>
                <Syncer drawing={drawing} setDrawing={setDrawing}/>
            </div>
           ); }
