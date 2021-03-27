package tz_converter

import (
	"log"
	"testing"
)

func TestFuzzyFindTimeZone(t *testing.T) {
	log.Println(FuzzyFindTimeZone("Denver"))
}
