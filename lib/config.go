package lib

// Config contains the configuration
type Config struct {
	AuthEmail string // used for the cloudflare api
	AuthKey   string // used for the cloudflare api

	// the domain names which will be updated
	DomainNames []string
}
