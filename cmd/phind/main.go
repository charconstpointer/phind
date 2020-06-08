package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//TripDesc is
type TripDesc struct {
	Length float32 `json:"length"`
	Points int     `json:"points"`
	Seed   int     `json:"seed"`
}

//TripOpt is
type TripOpt struct {
	Desc TripDesc `json:"round_trip"`
}

//TripRequest is
type TripRequest struct {
	Coordinates [][2]float32 `json:"coordinates"`
	Options     TripOpt      `json:"options"`
}

//TripCoord return lng first then lat
type TripCoord struct {
	Coords [][2]float32 `json:"coordinates"`
}

type TripResFeature struct {
	Geomerty TripCoord `json:"geometry"`
}
type TripResponse struct {
	Features []TripResFeature `json:"features"`
}

func main() {
	const token = "5b3ce3597851110001cf6248479488b571734c25b1df9564a229ddb5"
	const api = "https://api.openrouteservice.org/v2/directions/foot-walking/geojson"
	coords := make([][2]float32, 1)

	coords[0] = [2]float32{21.1012782, 52.2954492}
	tripData := TripRequest{
		Coordinates: coords,
		Options: TripOpt{
			Desc: TripDesc{
				Length: 1000,
				Points: 10,
				Seed:   3,
			},
		},
	}

	b, err := json.Marshal(tripData)

	c := http.Client{}
	r, err := http.NewRequest("POST", api, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err.Error())
	}
	r.Header.Set("Authorization", token)
	r.Header.Set("Content-Type", "application/json")

	res, err := c.Do(r)
	if err != nil {
		fmt.Println(err.Error())
	}
	var t TripResponse
	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, f := range t.Features {
		for _, g := range f.Geomerty.Coords {
			fmt.Println(g)
		}
	}

}
