/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.3.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// UpdatesResponseUpdateList struct for UpdatesResponseUpdateList
type UpdatesResponseUpdateList struct {
	AvailableUpdates []UpdatesResponseAvailableUpdates `json:"available_updates,omitempty"`
	Description string `json:"description,omitempty"`
	Summary string `json:"summary,omitempty"`
}
