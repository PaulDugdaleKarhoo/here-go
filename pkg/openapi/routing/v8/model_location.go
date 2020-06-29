/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// Location Location on the Earth
type Location struct {
	// Location of a point on Earth north or south of the equator in decimal degrees.
	Lat float64 `json:"lat"`
	// Location of a place on Earth east or west of the prime meridian in decimal degrees.
	Lng float64 `json:"lng"`
	// The elevation of a point above mean sea level in meters.
	Elv float32 `json:"elv,omitempty"`
}
