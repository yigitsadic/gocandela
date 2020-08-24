package client

import "github.com/yigitsadic/gocandela/models"

type Client interface {
	ParseLines() ([]models.Earthquake, error)

	Fetch()
}
