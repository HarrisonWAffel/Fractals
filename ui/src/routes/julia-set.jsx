import Header from "./header";
import Button from "react-bootstrap/Button";
import useWebSocket from "react-use-websocket";
import {useEffect, useRef, useState} from "react";

export default function JuliaSet() {


    return <>
        <Header/>

        <video controls autoplay controlsList="nodownload" oncontextmenu="return false;">
            <source src={"http://localhost:8989/fractal.mp4"} type="video/mp4"/>
        </video>
    </>
}