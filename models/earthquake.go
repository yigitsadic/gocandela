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
