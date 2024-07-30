package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Float64String float64

func (f *Float64String) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		*f = Float64String(v)
		return nil
	}

	var num float64
	if err := json.Unmarshal(data, &num); err == nil {
		*f = Float64String(num)
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into Float64String", data)
}

// Structure of cities to request
type City struct {
	Value string        `json:"value"`
	Lat   Float64String `json:"lat"`
	Lng   Float64String `json:"lng"`
}

type CityResponse struct {
	Cities []City `json:"cities"`
}

// Weather structure
type Weather struct {
	Current float64 `json:"current"`
	Max     float64 `json:"max"`
	Min     float64 `json:"min"`
}

type CityWeather struct {
	City    string
	Current string
	Max     string
	Min     string
	Temp    float64
}

// Get the list of cities from the dastyar.io api
func getCities() ([]City, error) {
	url := "https://api.dastyar.io/express/clock/cities"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cityResponse CityResponse
	if err := json.NewDecoder(resp.Body).Decode(&cityResponse); err != nil {
		return nil, err
	}
	return cityResponse.Cities, nil
}

// Get real-time weather conditions based on geographic location and city
func getWeather(lat, lng float64) (float64, string, string, string, error) {
	url := fmt.Sprintf("https://api.dastyar.io/express/weather?lat=%f&lng=%f&theme=light", lat, lng)
	resp, err := http.Get(url)
	if err != nil {
		return 0, "", "", "", err
	}
	defer resp.Body.Close()

	var weatherResponse []Weather
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return 0, "", "", "", err
	}
	if len(weatherResponse) > 0 {
		current := math.Round(weatherResponse[0].Current)
		max := math.Round(weatherResponse[0].Max)
		min := math.Round(weatherResponse[0].Min)
		return current, fmt.Sprintf("%.0f°C", current), fmt.Sprintf("%.0f°C", max), fmt.Sprintf("%.0f°C", min), nil
	}
	return 0, "No data", "No data", "No data", nil
}

func updateTable() {
	cities, err := getCities()
	if err != nil {
		fmt.Printf("Error getting cities: %v\n", err)
		return
	}

	// Because the list of cities are many and I don't want to put too much pressure on the web services,
	// I have listed only a part of the important cities. You can add any city you want in English.

	targetCities := map[string]bool{
		"Ahvaz":        true,
		"Arak":         true,
		"Bandar Abbas": true,
		"Mashhad":      true,
		"Qom":          true,
		"Rasht":        true,
		"Shiraz":       true,
		"Tabriz":       true,
		"Tehran":       true,
		"Yazd":         true,
	}

	var cityWeathers []CityWeather

	for _, city := range cities {
		if _, ok := targetCities[strings.Title(city.Value)]; ok {
			lat := float64(city.Lat)
			lng := float64(city.Lng)

			temp, current, max, min, err := getWeather(lat, lng)
			if err != nil {
				fmt.Printf("Error getting weather for %s: %v\n", city.Value, err)
				continue
			}

			cityWeathers = append(cityWeathers, CityWeather{
				City:    city.Value,
				Current: current,
				Max:     max,
				Min:     min,
				Temp:    temp,
			})
		}
	}

	// Sort by current temperature in descending order
	sort.Slice(cityWeathers, func(i, j int) bool {
		return cityWeathers[i].Temp > cityWeathers[j].Temp
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"City", "Current Weather", "Max Temperature", "Min Temperature"})

	// Additional settings to make the table larger
	table.SetColWidth(20)
	table.SetRowLine(true)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetFooterAlignment(tablewriter.ALIGN_CENTER)

	for _, cityWeather := range cityWeathers {
		table.Append([]string{cityWeather.City, cityWeather.Current, cityWeather.Max, cityWeather.Min})
	}

	table.Render()
}

func main() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		//Re-call every 10 minutes for instant weather display
		updateTable() 

		fmt.Println("Next update in 10 minutes...")
		for i := 600; i > 0; i-- {
			fmt.Printf("\rTime remaining: %02d:%02d", i/60, i%60)
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
		<-ticker.C
	}
}
