import Header from "./header";
import Button from "react-bootstrap/Button";
import useWebSocket from "react-use-websocket";
import {useEffect, useState} from "react";

export default function JuliaSet() {

    const { sendMessage, lastMessage, readyState } = useWebSocket(
        "ws://localhost:8989/fractal",
        {
            retryOnError: true,
            reconnectAttempts: 5,
        }
    );

    useEffect(() => {
        console.log(lastMessage)
        setData((d) => d +1)
    }, [lastMessage])

    const [data, setData] = useState(0);

    function makereq() {
        sendMessage("julia-set")
    }

    return <>
        <Header/>
        <Button onClick={()=> makereq()}> Make Request </Button>
        <span>
            Number of chunks {data}
        </span>
    </>
}