package tz_converter

import (
	"encoding/json"
	fuzz "github.com/agext/levenshtein"
	"os"
	"time"
)

const (
	tzCitiesFile = "tz_cities.json"
)

type City struct {
	Country    string
	Name       string
	Lat        string
	Lng        string
	TimezoneID string
}

// TimeIn returns a time for a particular time zone based on the IANA name given
func TimeIn(t time.Time, name string) (time.Time, error) {
	location, err := time.LoadLocation(name)
	if err != nil {
		return time.Time{}, err
	}
	t = t.In(location)
	return t, nil
}

func FuzzyFindTimeZone(city string) ([]City, error) {
	var matches []City
	var locations []City
	data, err := os.ReadFile(tzCitiesFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return nil, err
	}
	for _, location := range locations {
		if fuzz.Match(city, location.Name, nil) == 1 {
			matches = append(matches, location)
		}
	}
	return matches, nil
}


func ReduceMatches(matches []City) []City {
	// remove slices
	return nil
}