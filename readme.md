
# Go-Weather

This Go package fetches the current weather data for a list of specified cities and displays the information in a neatly formatted table in the terminal. The table is updated every 10 minutes, and the cities are sorted based on their current temperature in descending order.

# Features
- Monitoring the weather of Iranian top cities.


- Displays the current temperature, maximum temperature, and minimum temperature.

- Updates the table every 10 minutes.
-Sorts the cities by current temperature in descending order.



# Installation
To use this package, you need to have Go installed on your machine. If you don't have Go installed, you can download and install it from here.


1. Clone the repository:
   ```bash
   git clone https://github.com/mohammadhasanii/Go-Weather.git 
   ```

2. Run the application
 ```bash
   go run main.go 
   ```

# Usage
Once the application is running, it will display a table with the current weather data for the specified cities. The table will be updated every 10 minutes, and a countdown timer will show the time remaining until the next update.

# Example Output

```mathematica
+---------+----------------+-----------------+-----------------+
|  CITY   | CURRENT WEATHER| MAX TEMPERATURE | MIN TEMPERATURE |
+---------+----------------+-----------------+-----------------+
|  Tehran |      35°C      |       38°C      |       30°C      |
| Esfahan |      32°C      |       34°C      |       28°C      |
|  Shiraz |      30°C      |       33°C      |       27°C      |
|   Arak  |      28°C      |       30°C      |       25°C      |
+---------+----------------+-----------------+-----------------+
Next update in 10 minutes...
Time remaining: 09:59

```

![Logo](./demo.png)


# Contributing
If you would like to contribute to this project, please fork the repository and submit a pull request.






