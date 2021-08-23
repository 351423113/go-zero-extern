package internal

import (
	"fmt"
	"strings"

	"github.com/351423113/go-zero-extern/zrpc/internal/resolver"
)

// BuildDirectTarget returns a string that represents the given endpoints with direct schema.
func BuildDirectTarget(endpoints []string) string {
	return fmt.Sprintf("%s:///%s", resolver.DirectScheme,
		strings.Join(endpoints, resolver.EndpointSep))
}

// BuildDiscovTarget returns a string that represents the given endpoints with discov schema.
func BuildDiscovTarget(endpoints []string, key string) string {
	return fmt.Sprintf("%s://%s/%s", resolver.DiscovScheme,
		strings.Join(endpoints, resolver.EndpointSep), key)
}
