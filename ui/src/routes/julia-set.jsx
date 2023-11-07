import Header from "./header";
import Button from "react-bootstrap/Button";
import useWebSocket from "react-use-websocket";
import {useEffect, useRef, useState} from "react";
import {Card, InputGroup} from "react-bootstrap";
import Form from 'react-bootstrap/Form';

export default function JuliaSet() {


    // need a button to start the render
    // need fields for the following

    /*
        Initial Constant Real
        Constant Imaginary
        Total Range
        Step Size

        Video height and width are going to be hardcoded to 1000x1000
     */


    function SideBar() {
        return <>

        </>
    }

    function getUrl() {
        setUrl("http://localhost:8989/julia-set.mp4?constant-real="+constantReal+"&constant-imaginary="+constantImaginary+"&total-range="+totalRange+"&step-size="+stepSize+"&zoom="+zoom)
    }


    // probably not using this right
    const [url, setUrl] = useState("")
    const [displayedVideo, setDisplayedVideo] = useState(false)
    const [constantReal, setConstantReal] = useState("0.280")
    const [constantImaginary, setConstantImaginary] = useState("0.01")
    const [totalRange, setTotalRange] = useState("0.005")
    const [stepSize, setStepSize] = useState("0.000001")
    const [zoom, setZoom] = useState("1.0")

    return <>
        <Header/>
        <div style={{display: "flex", flexDirection: "row"}}>
            <Card style={{height: "26rem", marginLeft: "5%", marginRight: "5%", marginTop: "2.5%" }}>
                <Card.Body>
                    <Card.Title>Julia Set Creator</Card.Title>
                    <Card.Text>
                        Modify the below values to generate a julia set
                    </Card.Text>

                    <InputGroup className="mb-3">
                        <InputGroup.Text id="constant-real">Constant Real</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setConstantReal(v.target.value)}}
                            value={constantReal}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text id="constant-imaginary">Constant Imaginary</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setConstantImaginary(v.target.value)}}
                            value={constantImaginary}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text id="total-range">Total Range</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setTotalRange(v.target.value)}}
                            value={totalRange}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text >Step Size</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setStepSize(v.target.value)}}
                            id="step-size"
                            value={stepSize}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text >Zoom</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setZoom(v.target.value)}}
                            id="step-size"
                            value={zoom}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <Button onClick={()=>{
                        getUrl()
                    }} variant="primary">Generate Julia Set</Button>
                </Card.Body>
            </Card>
            <video key={url} id={"vid"} style={{width: "1000px", height: "1000px", marginTop: "2.5%"}} controls autoPlay controlsList="nodownload" onContextMenu="return false;">
                <source src={url} type="video/mp4"/>
            </video>
        </div>
    </>
}