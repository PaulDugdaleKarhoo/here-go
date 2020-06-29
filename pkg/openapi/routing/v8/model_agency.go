/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// Agency Contains information about a particular agency.
type Agency struct {
	// Unique code of the agency. Specifies if the same agency is used on different sections of the same route.
	Id string `json:"id"`
	// Human readable name of the owner of the transport service.
	Name string `json:"name"`
	// Link to the agency's website.
	Website string `json:"website,omitempty"`
	// Agency icon url
	Icon string `json:"icon,omitempty"`
}
