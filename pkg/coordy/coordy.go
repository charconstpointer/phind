package coordy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//GetCoords is
func GetCoords(start [2]float32) Response {
	coords := make([][2]float32, 1)
	coords[0] = start
	r := TripRequest{
		Coordinates: coords,
		Options: TripOpt{
			Desc: TripDesc{
				Length: 1000,
				Points: 10,
				Seed:   3,
			},
		},
	}
	return getPath(r)
}

//GetPath is
func getPath(tripData TripRequest) Response {
	fmt.Println("getPath")
	const api = "https://api.openrouteservice.org/v2/directions/foot-walking/geojson"
	const token = "5b3ce3597851110001cf6248479488b571734c25b1df9564a229ddb5"
	b, err := json.Marshal(tripData)
	if err != nil {
		fmt.Println(err.Error())
	}
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
	var response Response
	for _, f := range t.Features {
		for _, g := range f.Geomerty.Coords {

			c := Coord{Lng: g[0], Lat: g[1]}
			fmt.Println(c)
			response.Coords = append(response.Coords, c)
		}
	}
	return response
}
