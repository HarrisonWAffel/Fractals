#### SUSE Rancher Hackweek '23 project.

# Fractal Generator 

The goal of this project is to improve the implementation of my prior [Julia-Set generator](https://github.com/harrisonwaffel/julia-sets) as well as to provide a base to implement generators for other fractals. Currently, only the Julia set and Mandelbrot set are implemented. 


#### REST API Details

This project provides a REST API for the Julia and Mandelbrot set generators. By default, the API will listen on port `8989`. A Dockerfile has been included which packages the API as well as the ffmpeg dependencies. 

+ Endpoint: `/mandelbrot.png`
  + Generates a single image of the Mandelbrot set 
+ Query Parameters:
  + `zoom`: `int`, optional
    + The amount of zoom to apply to the image
  + `MoveX`: `float`, optional
    + Translates the image on the X axis.
  + `MoveY`: float, optional
    + Translates the image on the Y axis
+ Recommended request:
  + `/mandelbrot.png?movex=148.1003&movey=8.5001&zoom=1000`

+ Endpoint: `/mandelbrot.mp4`
  + Generates a video which slowly zooms into the Mandelbrot set. In order to not zoom into the middle of the set, translating on the X and Y axis is required.
+ Query Parameters:
    + `zoom`: `int`, optional
      + The amount of zoom to apply to the image
    + `zoom-step`: `float`, required
      + Controls the speed of the zoom animation
    + `MoveX`: `float`, optional
        + Translates the image on the X axis.
    + `MoveY`: float, optional
        + Translates the image on the Y axis
    + `duration`: `int`, required
      + The desired length of the resulting video
+ Recommended request: 
    + `/mandelbrot.mp4?movex=148.1003&movey=8.5001&zoom=1&duration=30&zoom-step-size=1`


+ Endpoint: `/julia-set.mp4`
  + Generates a video of a Julia-Set
+ Query Parameters:
  + `constant-real`: `float`, required
  + `constant-imaginary`: `float`, required
  + `total-range`: `float`, required
  + `step-size`: `float`, required
  + `movex`: `float`, optional
  + `movey`: `float`, optional
+ Recommended request: 
  + `/julia-set.mp4?constant-real=0.28&constant-imaginary=0.01&total-range=0.005&step-size=0.00001`


#### Included UI 

This project includes a **very simple** UI implemented in react it is not very fleshed out. However, the UI is useful for testing values used to generate videos and images. 

If you intend on viewing and saving a video, using the API directly is less prone to stuttering.