package coordy

//Coord is
type Coord struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

//Response is
type Response struct {
	Coords []Coord `json:"coords"`
}
