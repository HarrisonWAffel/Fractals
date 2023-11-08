import Header from "./header";
import {Card, InputGroup} from "react-bootstrap";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import {useState} from "react";

export default function Mandelbrot() {


    const [moveX, setMoveX] = useState("148.5")
    const [moveY, setMoveY] = useState("9.5")
    const [zoom, setZoom] = useState("1000")


    const [url, setUrl] = useState("")

    function updateUrl() {
        setUrl("http://localhost:8989/mandelbrot.png?movex="+moveX+"&movey="+moveY+"&zoom="+zoom)
    }

    return <>
        <Header/>
        <div style={{display: "flex", flexDirection: "row"}}>
            <Card style={{height: "26rem", marginLeft: "5%", marginRight: "5%", marginTop: "2.5%" }}>
                <Card.Body>
                    <Card.Title>Mandelbrot Set Creator</Card.Title>
                    <Card.Text>
                        Modify the below values to generate a mandelbrot set
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
                        <InputGroup.Text id="constant-imaginary">Move Y</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setMoveY(v.target.value)}}
                            value={moveY}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text id="total-range">Zoom</InputGroup.Text>
                        <Form.Control
                            onChange={(v)=>{setZoom(v.target.value)}}
                            value={zoom}
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <Button onClick={()=>{updateUrl()}} variant="primary">Preview Mandelbrot Set</Button>
                    <Button onClick={()=>{}} variant="primary">Download Mandelbrot Set</Button>
                </Card.Body>
            </Card>
            <img key={url} style={{marginTop: "2.5%"}} src={url}/>
        </div>
    </>
}