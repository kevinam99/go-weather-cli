# Makefile

# Define variables for convenience
# GO=go
# APP_NAME=myapp
# PORT=8080

.PHONY: all weather-coords set-fav weather-city favourites

all: build

# Get weather for specific coordinates
# Usage: make weather-coords LAT=15.4027 LON=74.0078
weather-coords:
	go run cmd/go-weather-cli/main.go get-weather -l ${LAT} -n ${LON}

# Add a city to favourites
# Usage: make set-fav CITY=Oslo LAT=59.9139 LON=10.7522
set-fav:
	go run cmd/go-weather-cli/main.go set-favourite ${CITY} ${LAT} ${LON}

# Get weather for a favourite city
# Usage: make weather-city CITY=Oslo
weather-city:
	go run cmd/go-weather-cli/main.go get-weather -c ${CITY}

# List all favourite cities
favourites:
	go run cmd/go-weather-cli/main.go get-favourites

