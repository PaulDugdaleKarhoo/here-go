/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// Ev **Disclaimer: This parameter is currently in beta release, and is therefore subject to breaking changes.**  EV parameters to be used for calculating consumption and for calculating EV routes with automatically added charging stations.  The following attributes are required for calculating consumption: * `freeFlowSpeedTable` * `ascent` * `descent`  The following attributes are additionally required in order to calculate reachable routes:  * `initialCharge` * `maxCharge` * `chargingCurve` * `maxChargeAfterChargingStation` * `makeReachable` set to `true`  All remaining attributes are optional. 
type Ev struct {
	// Function curve specifying consumption rate at a given free flow speed on a flat stretch of road.  The format of the string is a comma-separated list of numbers, as follows:  ``` <SPEED_0>,<CONSUMPTION_0>,<SPEED_1>,<CONSUMPTION_1>,...,<SPEED_N>,<CONSUMPTION_N>` ```  where speed values are strictly increasing, non-negative integers in units of (km/h), and consumption values are non-negative floating point values in units of (Wh/m).  The function is linearly interpolated between each successive pair of data points. For speeds less than `SPEED_0` the value of the function is `CONSUMPTION_0`. For speeds greater than `SPEED_N` the value of the function is `CONSUMPTION_N`. 
	FreeFlowSpeedTable string `json:"freeFlowSpeedTable"`
	// Function curve specifying consumption rate at a given traffic-reduced speed on a flat stretch of road.  See `freeFlowSpeedTable` for a description of the string format. 
	TrafficSpeedTable string `json:"trafficSpeedTable,omitempty"`
	// Rate of energy consumed per meter rise in elevation (in Wh/m, i.e., Watt-hours per meter). 
	Ascent float32 `json:"ascent"`
	// Rate of energy recovered per meter fall in elevation (in Wh/m, i.e., Watt-hours per meter). 
	Descent float32 `json:"descent"`
	// Rate of energy (in Wh/s) consumed by the vehicle's auxiliary systems (e.g., air conditioning, lights). The value represents the number of Watt-hours consumed per second of travel. 
	AuxiliaryConsumption float32 `json:"auxiliaryConsumption,omitempty"`
	// Charge level of the vehicle's battery at the start of the route (in kWh). Value must be less than or equal to the value of `maxCharge`. 
	InitialCharge float32 `json:"initialCharge,omitempty"`
	// Total capacity of the vehicle's battery (in kWh). 
	MaxCharge float32 `json:"maxCharge,omitempty"`
	// Function curve describing the maximum battery charging rate (in kW) at a given charge level (in kWh).  The format of the string is a comma-separated list of numbers, as follows:  ``` <CHARGE_0>,<RATE_0>,<CHARGE_1>,<RATE_1>,...,<RATE_N>,<CHARGE_N> ```  where charge values are strictly increasing, non-negative floating-point values in units of (kWh), and rate values are positive floating point values in units of (kW).  Charge values must cover the entire range of `[0, maxChargeAfterChargingStation`]. The charging curve is piecewise constant, e.g., for any charge in the range `[CHARGE_0, CHARGE_1)`, the value of the function is `RATE_0`. 
	ChargingCurve string `json:"chargingCurve,omitempty"`
	// Maximum charging voltage supported by the vehicle's battery (in Volt). 
	MaxChargingVoltage float32 `json:"maxChargingVoltage,omitempty"`
	// Maximum charging current supported by the vehicle's battery (in Ampere). 
	MaxChargingCurrent float32 `json:"maxChargingCurrent,omitempty"`
	// Maximum charge to which the battery should be charged at a charging station (in kWh). Value must be less than or equal to the value of `maxCharge`. 
	MaxChargeAfterChargingStation float32 `json:"maxChargeAfterChargingStation,omitempty"`
	// Minimum charge when arriving at a charging station (in kWh). Value must be less than the value of `maxChargeAfterChargingStation`. 
	MinChargeAtChargingStation float32 `json:"minChargeAtChargingStation,omitempty"`
	// Minimum charge at the final route destination (in kWh). Value must be less than the value of `maxChargeAfterChargingStation`. 
	MinChargeAtDestination float32 `json:"minChargeAtDestination,omitempty"`
	// Time spent after arriving at a charging station, but before actually charging, e.g., time spent for payment processing (in seconds). 
	ChargingSetupDuration int32 `json:"chargingSetupDuration,omitempty"`
	// When set to `true`, the router ensures that the calculated route is reachable within the given constraints (i.e., `minChargeAtChargingStation`, `minChargeAtDestination`). If necessary, charging stations are added to the route to achieve reachability.  The following conditions must be met in order to enable this option: * `transportMode=car` * `routingMode=fast` * No `avoid` options requested * No `alternatives` requested 
	MakeReachable bool `json:"makeReachable,omitempty"`
}
