package wheretheiss

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Position struct {
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	TimezoneID  string `json:"timezone_id"`
	Offset      int    `json:"offset"`
	CountryCode string `json:"country_code"`
	MapURL      string `json:"map_url"`
}

func PositionWhereTheIss(d chan string) {

	response, err := http.Get("https://api.wheretheiss.at/v1/satellites/25544")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	coordinates := &Coordinates{}
	json.Unmarshal(body, coordinates)

	if err != nil {
		log.Fatal(err)
	}

	var name = "Water"
	CountryWhereTheIss(coordinates.Latitude, coordinates.Longitude, &name)

	d <- name
}

func CountryWhereTheIss(lat float64, long float64, name *string) {
	s := fmt.Sprintf("%f", lat)
	r := fmt.Sprintf("%f", long)

	response, err := http.Get("https://api.wheretheiss.at/v1/coordinates/" + s + "," + r)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	position := &Position{}
	json.Unmarshal(body, position)

	if err != nil {
		log.Fatal(err)
	}

	if position.CountryCode != "" {
		*name = position.CountryCode
	}

}
