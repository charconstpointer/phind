package coordy

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
