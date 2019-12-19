package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cezarmathe/cf-dns-auto-updater/lib"
	"github.com/joho/godotenv"
)

var (
	config *lib.Config
)

func init() {
	// check if we have a connection to the internet
	// required to check the new public ip and accessing the cloudflare api
	_, err := http.Get("https://1.1.1.1")
	if err != nil {
		fmt.Println("no internet connection")
		os.Exit(1)
	}

	// check if .env should not be loaded
	noDotEnv, isSet := os.LookupEnv("CF_DNS_AUTO_UPDATER_NO_DOTENV")
	if isSet && len(noDotEnv) > 0 {
		fmt.Println("not loading .env")
		return
	}

	// load .env
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("error opening .env - %s\n", err.Error())
	}
}

func main() {
	// load the configuration
	authEmail, isSet := os.LookupEnv("CF_DNS_AUTO_UPDATER_AUTH_EMAIL")
	if !isSet || len(authEmail) == 0 {
		fmt.Println("CF_DNS_AUTO_UPDATER_AUTH_EMAIL not set")
		os.Exit(1)
	}
	authKey, isSet := os.LookupEnv("CF_DNS_AUTO_UPDATER_AUTH_KEY")
	if !isSet || len(authKey) == 0 {
		fmt.Println("CF_DNS_AUTO_UPDATER_AUTH_KEY not set")
		os.Exit(1)
	}
	domainNames, isSet := os.LookupEnv("CF_DNS_AUTO_UPDATER_DOMAIN_NAMES")
	if !isSet || len(domainNames) == 0 {
		fmt.Println("CF_DNS_AUTO_UPDATER_DOMAIN_NAMES not set")
		os.Exit(1)
	}

	config = new(lib.Config)
	config.AuthEmail = authEmail
	config.AuthKey = authKey
	config.DomainNames = strings.Split(domainNames, " ")
	// FIXME 18/12/2019: validate domain names

	fmt.Printf("Domain names to check: %s\n", config.DomainNames)
}
