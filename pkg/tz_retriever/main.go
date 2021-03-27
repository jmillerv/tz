package main

import (
	"encoding/json"
	tz "github.com/jmillerv/timezone/pkg/tz_converter"
	timezone "gopkg.in/ugjka/go-tz.v2/tz"
	"log"
	"os"
	"strconv"
)

const (
	sourceCities      = "cities.json"
	destinationCities = "tz_cities.json"
)


func loadCityJson() []tz.City {
	data, err := os.ReadFile(sourceCities)
	if err != nil {
		log.Println(err)
	}

	var cities []tz.City
	err = json.Unmarshal(data, &cities)
	if err != nil {
		log.Println(err)
	}
	return cities
}

func addTimezone(cities []tz.City) []tz.City {
	var tzCities []tz.City
	for _, city := range cities {
		lon, _ := strconv.Atoi(city.Lng)
		lat, _ := strconv.Atoi(city.Lat)
		zone, err := timezone.GetZone(timezone.Point{
			Lon: float64(lon),
			Lat: float64(lat),
		})
		if err != nil {
			log.Println(err)
		}
		city.TimezoneID = zone[0]
		tzCities = append(tzCities, city)
	}
	return tzCities
}

func main() {
	cities := loadCityJson()
	tzCities := addTimezone(cities)
	file, _ := json.MarshalIndent(tzCities, "", " ")
	err := os.WriteFile(destinationCities, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
