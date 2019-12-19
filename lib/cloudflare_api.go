package lib

import (
	"net"

	"github.com/cloudflare/cloudflare-go"
)

var (
	// API is the cloudflare api
	API *cloudflare.API
)

// NewAPI initializes a new cloudflare api
func NewAPI(config *Config) error {
	// create a new cloudflare api
	api, err := cloudflare.New(config.AuthKey, config.AuthEmail)
	if err != nil {
		return err
	}

	// set the api object
	API = api

	return nil
}

// PerformUpdate check the zones from Cloudflare and updates the IPs if necessary
func PerformUpdate(ip *net.IPAddr, zones []Zone) error {

	return nil
}
