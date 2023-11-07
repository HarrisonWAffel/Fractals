import './App.css';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import React, { useState, useCallback, useEffect } from 'react';
import Button from 'react-bootstrap/Button';

function App() {

  const { sendMessage, lastMessage, readyState } = useWebSocket(
      "ws://localhost:8989/fractal",
      {
        retryOnError: true,
        reconnectAttempts: 5,
      }
  );
  const [messageHistory, setMessageHistory] = useState([]);

  useEffect(() => {
    if (lastMessage !== null) {
      console.log(lastMessage)
      setMessageHistory((prev) => prev.concat(lastMessage.data));
    }
  }, [lastMessage, setMessageHistory]);

  function sendData(fractal) {
    console.log("sending data")
    sendMessage(fractal)
  }

  const fractals = [
      "julia-set",
      "mandelbrot"
  ]


  return (
    <div className="App">
      {lastMessage ? <span>Last message: {lastMessage.data}</span> : null}
      <ul>
        {messageHistory.map((message, idx) => (
            <span key={idx}>{message ? message.data : null}</span>
        ))}
      </ul>
      <Button className="btn btn-primary btn-lg mx-3 px-5 py-3 mt-2" onClick={() => sendData(fractals[0])}> julia-set </Button>
      <Button onClick={() => sendData(fractals[1])}> mandelbrot </Button>
    </div>
  );
}

export default App;
