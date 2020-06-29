/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// ErrorResponse Response in case of error
type ErrorResponse struct {
	// Human readable error description
	Title string `json:"title"`
	// HTTP status code
	Status int32 `json:"status"`
	// Error code.  All error codes start with \"`E60`\". 
	Code string `json:"code"`
	// Human readable explanation for the error
	Cause string `json:"cause"`
	// Human readable action that can be taken to correct the error
	Action string `json:"action"`
	// Auto generated id that univocally identify this request
	CorrelationId string `json:"correlationId"`
}
