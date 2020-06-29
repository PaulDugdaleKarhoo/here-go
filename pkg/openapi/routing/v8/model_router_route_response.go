/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// RouterRouteResponse Returns a list of routes.
type RouterRouteResponse struct {
	// Contains a list of issues encountered during the processing of this response.
	Notices []RouteResponseNotice `json:"notices,omitempty"`
	// List of possible routes
	Routes []RouterRoute `json:"routes"`
}
