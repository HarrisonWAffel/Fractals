import Header from "./header";
import {Card, InputGroup} from "react-bootstrap";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import {useState} from "react";

export default function MandelbrotVideo() {


    const [moveX, setMoveX] = useState("148.1003")
    const [moveY, setMoveY] = useState("8.5001")
    const [zoom, setZoom] = useState("1000")
    const [zoomSteps, setZoomSteps] = useState("5")
    const [seconds, setSeconds] = useState("30")

    const [url, setUrl] = useState("")

    function updateUrl() {
        if (url !== "") {
            return
        }
        let x = "http://localhost:8989/mandelbrot.mp4?movex="+moveX+"&movey="+moveY+"&zoom="+zoom+"&zoom-step="+zoomSteps+"&seconds="+seconds
        document.getElementById('vid').defaultPlaybackRate = -1.0;
        console.log(x)
        setUrl(x)
    }

    return <>
        <Header/>
        <div style={{display: "flex", flexDirection: "row"}}>
            <Card style={{height: "28rem", marginLeft: "5%", marginRight: "5%", marginTop: "2.5%" }}>
                <Card.Body>
                    <Card.Title>Mandelbrot Set Creator</Card.Title>
                    <Card.Text>
                        Modify the below values to generate a Mandelbrot set.
                    </Card.Text>
                    <Card.Text>
                        Depending on the values you provide, the video may buffer several times
                    </Card.Text>
                    <InputGroup className="mb-3">
                        <InputGroup.Text >Move X</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setMoveX(v.target.value)}}
                            value={moveX}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text >Move Y</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setMoveY(v.target.value)}}
                            value={moveY}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text >Zoom</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setZoom(v.target.value)}}
                            value={zoom}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text >Zoom Steps</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setZoomSteps(v.target.value)}}
                            value={zoomSteps}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text >Duration in seconds</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setSeconds(v.target.value)}}
                            value={seconds}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <Button onClick={()=>{updateUrl()}} variant="primary">Update Mandelbrot Set</Button>
                </Card.Body>
            </Card>
            <video key={url} id={"vid"} style={{width: "1000px", height: "1000px", marginTop: "2.5%"}} controls onContextMenu="return false;">
                <source src={url} type="video/mp4"  onLoadStart="this.playbackRate = 0.5;"/>
            </video>
        </div>
    </>
}