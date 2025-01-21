# Go Weather CLI

## About
I built this CLI to learn Go. I wanted to do something that would be useful, fun and simple.

It uses the OpenWeatherMap API to get weather information.

The main functionality is to get the weather for the given coordinates.

I've also added a few other features:
- Adding a city to favourites
- Getting the weather for a favourite city


## Usage

To fetch the weather for the given coordinates, run in the format `get-weather -l <latitude> -n <longitude>`:
```bash
go run cmd/go-weather-cli/main.go get-weather -l 15.4027 -n 74.0078 
```

To add a city to favourites, run in the format `set-favourite <city> <latitude> <longitude>`:
```bash
go run cmd/go-weather-cli/main.go set-favourite Oslo 59.9139 10.7522  
```

You can add as many cities as you want.

To get the weather for a favourite city, run in the format `get-weather -c <favourite_city>`:
```bash
go run cmd/go-weather-cli/main.go get-weather -c Oslo
```

To get the list of favourite cities, run:
```bash
go run cmd/go-weather-cli/main.go get-favourites
```

I hope you find this useful. Feel free to contribute to the project.

## Convenience commands

I've added a Makefile to make it easier to run the commands.

To get the weather for the given coordinates, run `make weather-coords LAT=15.4027 LON=74.0078`.

To add a city to favourites, run `make set-fav CITY=Oslo LAT=59.9139 LON=10.7522`.

To get the weather for a favourite city, run `make weather-city CITY=Oslo`.

To get the list of favourite cities, run `make favourites`.


## Contributing

I'm not sure if this is the best way to do it. I'm open to suggestions.

Feel free to contribute to the project. 🎉🚀

## License

This project is licensed under the MIT License - see the LICENSE file for details.
