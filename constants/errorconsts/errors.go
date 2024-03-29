package errorconsts

import (
	"strconv"
)

// These error constants are to be used in code for corresponding errors
// Their value doesn't have much importance but try to keep it same as the error
// just to be consistent with the convention
const (
	// InvalidRequest key
	InvalidRequest = "InvalidRequest"
	// NoFlightsAvailable key
	NoFlightsAvailable = "NoFlightsAvailable"
	// SameStartEndCity key
	SameStartEndCity = "SameStartEndCity"
	// InvalidFlightSchedule key
	InvalidFlightSchedule = "InvalidFlightSchedule"
)

const (
	// GenericErrorMessage message
	GenericErrorMessage = "Oops! I missed that test case. Please contact me on somprabhsharma@outlook.com with details and I will fix it."
	// GenericErrorCode code
	GenericErrorCode = 100
)

const (
	// InvalidRequestCode code
	InvalidRequestCode = 101
	// NoFlightsAvailableCode code
	NoFlightsAvailableCode = 102
	// SameStartEndCityCode code
	SameStartEndCityCode = 103
	// InvalidFlightScheduleCode code
	InvalidFlightScheduleCode = 104
)

// LTError is custom error for the micro service
type LTError struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	Title    string `json:"title,omitempty"`
	Details  string `json:"details,omitempty"`
	Err      string `json:"-"`
	HTTPCode int    `json:"-"`
}

// Error prints code & message of the error along with details
func (err LTError) Error() string {
	message := strconv.Itoa(err.Code) + ": " + err.Message
	if err.Details != "" {
		message = message + " " + err.Details
	}
	return message
}

// LTErrorMap is a map of error strings against LTError struct instances
var LTErrorMap = map[string]LTError{
	InvalidRequest: {
		Message: "Invalid request. Please provide all required parameters in the request.",
		Code:    InvalidRequestCode,
	},
	NoFlightsAvailable: {
		Message: "No flights available for the given cities.",
		Code:    NoFlightsAvailableCode,
	},
	SameStartEndCity: {
		Message: "Source and Destination cannot be same",
		Code:    SameStartEndCityCode,
	},
	InvalidFlightSchedule: {
		Message: "One or more flight schedule provided in the request are invalid. Please make sure each flight schedule has valid arrival and departure details.",
		Code:    InvalidFlightScheduleCode,
	},
}
