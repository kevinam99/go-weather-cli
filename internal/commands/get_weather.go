package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type WeatherInfo struct {
	Temperature float64 `json:"temperature"`
	Conditions  string  `json:"conditions"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
}

func getWeather() *cobra.Command {
	var city string
	var lat, lon float64

	cmd := &cobra.Command{
		Use:   "get-weather",
		Short: "Get weather information for a city or coordinates",
		Run: func(cmd *cobra.Command, args []string) {
			if city != "" && (lat != 0 || lon != 0) {
				fmt.Println("Error: Please provide either a city name or coordinates, not both")
				return
			}

			if city == "" && (lat == 0 && lon == 0) {
				fmt.Println("Error: Please provide either a city name or coordinates")
				return
			}

			if city != "" {
				coords, err := getFavouriteCoordinates(city)
				if err != nil {
					fmt.Printf("Error: Could not find coordinates for city %s: %v\n", city, err)
					return
				}
				lat, lon = coords.Lat, coords.Lon
			}

			weather, err := fetchWeather(lat, lon)
			if err != nil {
				fmt.Printf("Error fetching weather: %v\n", err)
				return
			}

			// Display weather information
			fmt.Printf("Weather for location %s, %s (%.4f, %.4f):\n", weather.City, weather.Country, lat, lon)
			fmt.Printf("Temperature: %.1f°C\n", weather.Temperature)
			fmt.Printf("Conditions: %s\n", weather.Conditions)
			fmt.Printf("Humidity: %d%%\n", weather.Humidity)
			fmt.Printf("Wind Speed: %.1f m/s\n", weather.WindSpeed)
		},
	}

	cmd.Flags().StringVarP(&city, "city", "c", "", "City name to get weather for")
	cmd.Flags().Float64VarP(&lat, "latitude", "l", 0, "Latitude coordinate")
	cmd.Flags().Float64VarP(&lon, "longitude", "n", 0, "Longitude coordinate")

	return cmd
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func getFavouriteCoordinates(city string) (*Coordinates, error) {
	// TODO: Implement reading from favourites.json
	// This function should:
	// 1. Read the favourites.json file
	// 2. Parse the JSON
	// 3. Find the coordinates for the given city
	// 4. Return  an error if city not found
	favorites, err := readFavorites()
	if err != nil {
		return &Coordinates{}, fmt.Errorf("failed to read favorites: %v", err)
	}

	if len(favorites) == 0 {
		return &Coordinates{}, fmt.Errorf("no favorites found")
	}

	for _, favorite := range favorites {
		if favorite.City == city {
			return &Coordinates{Lat: favorite.Latitude, Lon: favorite.Longitude}, nil
		}
	}
	return nil, fmt.Errorf("not implemented")
}

func fetchWeather(lat, lon float64) (WeatherInfo, error) {
	apiKey := "7dd9e0f7496e2aef381cf31d157c4377" // Replace with your actual API key
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.6f&lon=%.6f&appid=%s&units=metric", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherInfo{}, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherInfo{}, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var result struct {
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
		City string `json:"name"`
		Sys  struct {
			Country string `json:"country"`
		} `json:"sys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return WeatherInfo{}, fmt.Errorf("failed to parse weather data: %v", err)
	}

	// Create a new struct for our formatted response
	weatherInfo := WeatherInfo{
		Temperature: result.Main.Temp,
		Conditions:  result.Weather[0].Description,
		Humidity:    result.Main.Humidity,
		WindSpeed:   result.Wind.Speed,
		City:        result.City,
		Country:     result.Sys.Country,
	}

	return weatherInfo, nil
}
