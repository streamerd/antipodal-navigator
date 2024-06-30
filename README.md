# Antipodal Globe Explorer

This application allows users to explore antipodal points on a 3D globe using Mapbox GL JS and Go Fiber. When a user clicks on any point on the globe, the application calculates its antipodal point* and smoothly flies to that location.

Inspired by Neil deGrasse Tyson's explanation of antipodal points: [Startalk Short Video](https://www.youtube.com/shorts/HJwZoQq0-a0)


[![Antipodal Points Navigator](https://img.youtube.com/vi/W1R4Xj2NvTA/0.jpg)](https://www.youtube.com/watch?v=W1R4Xj2NvTA)


## What are Antipodal Points?

Antipodal points are two points on a sphere that are diametrically opposite to each other. On Earth, they are sometimes referred to as antipodes. If you were to draw a straight line from one point through the center of the Earth, it would emerge at the antipodal point.

## Features

- Interactive 3D globe using Mapbox GL JS
- Click on any point to fly to its antipodal location
- Display of starting and destination coordinates
- Smooth animation using Mapbox's flyTo effect

## Prerequisites

- Go 1.16 or later
- Mapbox API token

## Setup

1. Clone this repository:

`git clone https://github.com/yourusername/antipodal-navigator.git && cd antipodal-navigator`


2. Install dependencies: `go mod tidy`

3. Get a Mapbox API token:
- Sign up at [Mapbox](https://www.mapbox.com/)
- Navigate to your account's tokens page
- Create a new token or use the default public token

4. Set up your environment variables:
- Copy the `.env.example` file to `.env`
- Open `.env` and add your Mapbox token:
  ```
  MAPBOX_TOKEN=your_mapbox_public_key_here
  PORT=:3000
  ```

## Running the Application (Dev mode)

0. Set an environment var named DEV: `export DEV=1`
1. Start the server:
`go run main.go`

2. Open a web browser and navigate to `http://localhost:3000`

3. Click on any point on the globe to see the animation to its antipodal point!

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Neil deGrasse Tyson and Startalk for the inspiration
- Mapbox for their excellent mapping platform
- The Go Fiber team for the fast web framework