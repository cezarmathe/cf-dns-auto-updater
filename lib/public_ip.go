package lib

import (
	"io/ioutil"
	"net"
	"net/http"
)

const (
	// IPGetURL is the URL used for getting the IP
	IPGetURL = "https://api.ipify.org/"
)

// GetPublicIP attempts to get the public ip address
func GetPublicIP() (*net.IPAddr, error) {

	// make a request to IPGetURL to get the public ip address
	resp, err := http.Get(IPGetURL)
	if err != nil {
		return nil, err
	}

	// process the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// convert the response into the ip address
	return net.ResolveIPAddr("", string(body))
}
