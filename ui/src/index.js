import React from 'react';
import ReactDOM from 'react-dom/client';
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import './index.css';
import App from './App';
import Root from "./routes/root"
import ErrorPage from "./routes/not-found";
import JuliaSet from "./routes/julia-set";
import Mandelbrot from "./routes/mandelbrot";

import reportWebVitals from './reportWebVitals';
import 'bootstrap/dist/css/bootstrap.css';


const router = createBrowserRouter([
    {
        path: "/",
        element: <Root />,
        errorElement: <ErrorPage />
    },
    {
        path: "/julia-set",
        element: <JuliaSet />,
        errorElement: <ErrorPage/>
    },
    {
        path: "/mandelbrot",
        element: <Mandelbrot />,
        errorElement: <ErrorPage/>
    }
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
      <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
