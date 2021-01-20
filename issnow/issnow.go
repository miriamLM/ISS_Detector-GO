package issnow

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Coordinates struct {
	IssPosition struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	} `json:"iss_position"`
	Timestamp int    `json:"timestamp"`
	Message   string `json:"message"`
}

type Position struct {
	State     string `json:"state"`
	Latt      string `json:"latt"`
	City      string `json:"city"`
	Prov      string `json:"prov"`
	Geocode   string `json:"geocode"`
	Geonumber string `json:"geonumber"`
	Country   string `json:"country"`
	Stnumber  string `json:"stnumber"`
	Staddress string `json:"staddress"`
	Inlatt    string `json:"inlatt"`
}

func PositionIssNow(c chan string) {
	response, err := http.Get("http://api.open-notify.org/iss-now.json")

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
	CountryIssNow(coordinates.IssPosition.Latitude, coordinates.IssPosition.Longitude, &name)

	c <- name
}

func CountryIssNow(lat string, lon string, name *string) {
	response, err := http.Get("https://geocode.xyz/" + lat + "," + lon + "?json=1")

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

	if position.Prov != "" {
		*name = position.Prov
	}

}
