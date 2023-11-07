import Header from "./header";
import Button from "react-bootstrap/Button";
import useWebSocket from "react-use-websocket";

export default function JuliaSet() {

    const { sendMessage, lastMessage, readyState } = useWebSocket(
        "ws://localhost:8989/fractal",
        {
            retryOnError: true,
            reconnectAttempts: 5,
        }
    );

    function makereq() {
        sendMessage("julia-set")
    }

    return <>
        <Header/>
        <Button onClick={()=> makereq()}> Make Request </Button>
    </>
}