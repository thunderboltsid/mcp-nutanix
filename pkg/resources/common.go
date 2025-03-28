package resources

import (
	"fmt"
	"regexp"
)

// nutanixURI returns a URI for a resource type and UUID
func nutanixURI(resourceType, uuid string) string {
	return fmt.Sprintf("%s://%s", resourceType, uuid)
}

// extractIDFromURI extracts the UUID from a URI
// it uses regex to extract the UUID from the URI
// uri is expected to be in the format of resourceType://uuid or resourceType://name
func extractIDFromURI(uri string) string {
	re := regexp.MustCompile(`^[^:]+://([^/]+)$`)
	matches := re.FindStringSubmatch(uri)
	if len(matches) != 2 {
		return ""
	}

	return matches[1]
}

// extractTypeFromURI extracts the resource type from a URI
// it uses regex to extract the resource type from the URI
func extractTypeFromURI(uri string) string {
	re := regexp.MustCompile(`^([^:]+)://[^/]+$`)
	matches := re.FindStringSubmatch(uri)
	if len(matches) != 2 {
		return ""
	}

	return matches[1]
}
