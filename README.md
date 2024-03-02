# CLI Weather App

Welcome to the CLI Weather App! This application fetches weather data from the WeatherAPI and displays the current weather information along with the hourly forecast for a specified location.

## Features

- Displays current weather information (temperature, condition) for a given location.
- Shows the hourly forecast for the day including temperature, chance of rain, and weather condition.

### Prerequisites

- Go installed on your system. If not, you can download and install it from [here](https://golang.org/dl/).
- WeatherAPI key. You can sign up for a free key at [WeatherAPI](https://www.weatherapi.com/).

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/praveen-rikhari/CLI-Weather-App-in-Golang.git
    ```

2. Navigate to the project directory:

    ```bash
    cd CLI-Weather-App-in-Golang
    ```

3. Install this dependency for loading environment variables from `.env` file.

    ```bash
    go get github.com/joho/godotenv
    ```

4. Install this dependency for colored output.

    ```bash
    go get github.com/fatih/color
    ```

5. Create a `.env` file in the project root and add your WeatherAPI key:

    ```bash
    API_KEY="your_api_key_here"
    ```

6. Build the project:

    ```bash
    go build
    ```

7. To use this command anywhere from your terminal move `weather` file (created after running `go build` command) into your `/usr/local/bin` folder.
   ```
   mv weather /usr/local/bin
   ```
   OR
   ```
   sudo mv weather /usr/local/bin
   ```
### Usage

To run the Weather App, use the following command any where in your terminal:

```bash
weather [location]
```
replace `[location]` with the region for which you wanted the weather information.

## Preview

![Screenshot from 2024-03-03 01-48-26](https://github.com/praveen-rikhari/CLI-Weather-App-in-Golang/assets/84331681/6d647596-4648-4de3-b926-fd06b0fdf98d)

## Contribution

If you'd like to contribute to this project, feel free to open issues or submit pull requests. Contributions are welcome!
