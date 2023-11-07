import Button from "react-bootstrap/Button";
import {Container, Nav, Navbar, NavDropdown} from "react-bootstrap";

export default function Root() {
    return (
        <>
            <Navbar expand={"lg"} className={"bg-body-tertiary"} id="sidebar">
                <Container>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Navbar.Brand>Fractal Generator</Navbar.Brand>
                        <Nav>
                            <NavDropdown title="Available Fractals" id="basic-nav-dropdown">
                                <NavDropdown.Item href={`/julia-set`}>Julia Sets</NavDropdown.Item>
                                <NavDropdown.Item href={`/mandelbrot`}>Mandelbrot</NavDropdown.Item>
                            </NavDropdown>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
            <div id="detail"></div>
        </>
    );
}