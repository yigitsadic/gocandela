package client

import "github.com/yigitsadic/gocandela/models"

type Client interface {
	ParseToLines() ([]models.Earthquake, error)

	Fetch()
}
