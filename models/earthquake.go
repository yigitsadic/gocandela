package models

type Earthquake struct {
	Date      string `json:"date"`
	Time      string `json:"time"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Depth     string `json:"depth"`
	Md        string `json:"md"`
	Ml        string `json:"ml"`
	Mw        string `json:"mw"`
	Location  string `json:"location"`
}

func NewEarthquake(element []string) *Earthquake {
	return &Earthquake{
		Date:      element[1],
		Time:      element[2],
		Latitude:  element[3],
		Longitude: element[4],
		Depth:     element[5],
		Md:        element[6],
		Ml:        element[7],
		Mw:        element[8],
		Location:  element[9],
	}
}
