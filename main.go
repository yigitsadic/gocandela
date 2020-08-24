package main

import (
	"encoding/json"
	"github.com/yigitsadic/gocandela/client"
	"os"
)

func main() {
	c := client.NewHTTPClient()
	c.Fetch()
	res, err := c.ParseLines()
	if err != nil {
		os.Exit(1)
	}

	err = json.NewEncoder(os.Stdout).Encode(res)
	if err != nil {
		os.Exit(1)
	}
}
