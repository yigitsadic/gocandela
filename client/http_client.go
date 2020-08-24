package client

import (
	"github.com/yigitsadic/gocandela/models"
	"github.com/yigitsadic/gocandela/shared"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type HTTPClient struct {
	Response string
	Re       *regexp.Regexp
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{Re: regexp.MustCompile(shared.LineMatcher)}
}

func (h *HTTPClient) ParseToLines() ([]models.Earthquake, error) {
	sm := h.Re.FindAllStringSubmatch(h.Response, -1)

	var earthquakes []models.Earthquake

	for _, element := range sm {
		earthquakes = append(earthquakes, models.Earthquake{
			Date:      element[1],
			Time:      element[2],
			Latitude:  element[3],
			Longitude: element[4],
			Depth:     element[5],
			Md:        element[6],
			Ml:        element[7],
			Mw:        element[8],
			Location:  element[9],
		})
	}

	return earthquakes, nil
}

func (h *HTTPClient) Fetch() {
	res, err := http.Get(shared.BaseURL)
	if err != nil {
		log.Fatalf("Unable to fetch from Kandilli")
	}

	defer res.Body.Close()

	read, _ := ioutil.ReadAll(res.Body)
	h.Response = string(read)
}
