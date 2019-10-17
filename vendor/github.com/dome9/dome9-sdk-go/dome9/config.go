package dome9

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	defaultBaseURL = "https://api.dome9.com/v2/"
	defaultTimeout = 120 * time.Second
	loggerPrefix   = "dome9-logger: "
)

// Config contains all the configuration data for the API client
type Config struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	// The logger writer interface to write logging messages to. Defaults to standard out.
	Logger *log.Logger
	// Credentials for basic authentication.
	AccessID, SecretKey string
}

/*
NewConfig returns a default configuration for the client.
By default it will try to read the access and te secret from the environment variable.
*/

// TODO Add healthCheck method to NewConfig
func NewConfig(accessID, secretKey, rawUrl string) (*Config, error) {
	if accessID == "" || secretKey == "" {
		accessID = os.Getenv("DOME9_ACCESS_ID")
		secretKey = os.Getenv("DOME9_SECRET_KEY")
	}
	if rawUrl == "" {
		rawUrl = defaultBaseURL
	}

	var logger *log.Logger
	if loggerEnv := os.Getenv("DOME9_SDK_LOG"); loggerEnv == "true" {
		logger = getDefaultLogger()
	}

	baseURL, err := url.Parse(rawUrl)
	return &Config{
		BaseURL:    baseURL,
		HTTPClient: getDefaultHTTPClient(),
		Logger:     logger,
		AccessID:   accessID,
		SecretKey:  secretKey,
	}, err
}

func getDefaultHTTPClient() *http.Client {
	return &http.Client{Timeout: defaultTimeout}
}

func getDefaultLogger() *log.Logger {
	return log.New(os.Stdout, loggerPrefix, log.LstdFlags|log.Lshortfile)
}
