package resources

import (
	"fmt"
	"strings"
)

// NutanixURI returns a URI for a resource type and UUID
func NutanixURI(resourceType, uuid string) string {
	return fmt.Sprintf("%s://%s", resourceType, uuid)
}

// extractIDFromURI extracts the UUID from a URI
// it uses regex to extract the UUID from the URI
// uri is expected to be in the format of resourceType://uuid or resourceType://name
func extractIDFromURI(uri string) string {
	parts := strings.Split(uri, "://")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

// extractTypeFromURI extracts the resource type from a URI
// uri is expected to be in the format of resourceType://uuid or resourceType://name
func extractTypeFromURI(uri string) string {
	parts := strings.Split(uri, "://")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}
