package flightpath

// LazyJackRequest is struct of body for lazy jack api
type LazyJackRequest struct {
	PreferredTime int64           `json:"preferred_time,omitempty"`
	TripPlan      *TripDetail     `json:"trip_plan" binding:"required"`
	Schedules     []*FlightDetail `json:"schedules" binding:"required"`
}

// TripDetail is the details of the trip i.e. start, end city
type TripDetail struct {
	StartCity string `json:"start_city" binding:"required"`
	EndCity   string `json:"end_city" binding:"required"`
}

// ScheduleDetail is the schedule detail of a flight either arrival or departure schedule
type ScheduleDetail struct {
	City      string `json:"city" binding:"required"`
	Timestamp int64  `json:"timestamp" binding:"required"`
}

// FlightDetail is the flight details i.e arrival, departure details
type FlightDetail struct {
	Departure *ScheduleDetail `json:"departure" binding:"required"`
	Arrival   *ScheduleDetail `json:"arrival" binding:"required"`
}
