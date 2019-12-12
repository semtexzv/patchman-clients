/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// DnfModule Representation of one DNF module
type DnfModule struct {
	Name string `json:"name,omitempty"`
	Stream string `json:"stream,omitempty"`
}
