package services

import (
	"net/http"
	"sync"
)

// HTTPClientService provides a singleton HTTP client
type HTTPClientService struct {
	client *http.Client
}

var (
	httpClientInstance *HTTPClientService
	once               sync.Once
)

// GetHTTPClient returns the singleton instance of HTTPClientService
func GetHTTPClient() *HTTPClientService {
	once.Do(func() {
		httpClientInstance = &HTTPClientService{
			client: &http.Client{},
		}
	})
	return httpClientInstance
}

// Client returns the http.Client instance
func (s *HTTPClientService) Client() *http.Client {
	return s.client
}
