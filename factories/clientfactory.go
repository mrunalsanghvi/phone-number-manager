package factories

import (
	"phone-number-manager/client"
)

// NewHTTPClient creates an outbound HTTP client.
func NewHTTPClient() client.HTTPClient {
	return client.NewHTTPClient()
}