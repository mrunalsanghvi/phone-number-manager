package factories

import (
	"phone-number-manager/internal/client"
)

// NewHTTPClient creates an outbound HTTP client.
func NewHTTPClient() client.HTTPClient {
	return client.NewHTTPClient()
}
