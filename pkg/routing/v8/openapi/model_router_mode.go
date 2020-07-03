/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// RouterMode Mode of transport to be used for route calculation.
type RouterMode string

// List of RouterMode
const (
	ROUTERMODE_CAR RouterMode = "car"
	ROUTERMODE_TRUCK RouterMode = "truck"
	ROUTERMODE_PEDESTRIAN RouterMode = "pedestrian"
)