// Config loads configuration.
package config

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const Version = "1.1"

func init() {
	// set default downstream url if no env present
	if downstreamUrl := os.Getenv("DOWNSTREAM_URL"); downstreamUrl == "" {
		log.Println("No Env DOWNSTREAM_URL configured. Defaulting to http://localhost:9090")
		os.Setenv("DOWNSTREAM_URL", "http://localhost:9090")
	}
}

// GetInt loads the environment variable varName, converts it to an integer,
// and returns that integer or an error.
func GetInt(varName string) (int, error) {
	envVar := os.Getenv(varName)
	return strconv.Atoi(envVar)
}

func MustGetURL(urlEnvVar string) *url.URL {
	urlFromEnv := os.Getenv(urlEnvVar)
	if urlFromEnv == "" {
		log.Fatalf("No Env %s configured.", urlEnvVar)
	}
	parsedUrl, err := url.Parse(urlFromEnv)
	if err != nil {
		log.Fatalf("Invalid url: %s. %s\n", urlFromEnv, err.Error())
	}
	return parsedUrl
}

// SetMaxIdleConnsPerHost sets the MaxIdleConnsPerHost value for the default
// HTTP transport. If you are using a custom transport, calling this function
// won't change anything.
func SetMaxIdleConnsPerHost(maxConns int) {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = maxConns
}
