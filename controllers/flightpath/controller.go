package flightpath

import (
	"errors"
	"github.com/somprabhsharma/the-lazy-traveler/constants/errorconsts"
	"github.com/somprabhsharma/the-lazy-traveler/entities/flightpath"
	"github.com/somprabhsharma/the-lazy-traveler/models"
)

// Controller is a struct which will act like a controller
type Controller struct {
	Dao *models.Dao
}

// NewController is a constructor for Controller struct
func NewController(dao *models.Dao) *Controller {
	return &Controller{
		Dao: dao,
	}
}

// FindShortestFlightPath finds shortest flight path
func (c *Controller) FindShortestFlightPath(data flightpath.LazyJackRequest) ([]flightpath.ScheduleDetail, error) {
	if data.TripPlan.StartCity == data.TripPlan.EndCity {
		return nil, errors.New(errorconsts.SameStartEndCity)
	}

	// get shortest path data from cache if present
	shortestPath, err := c.Dao.FlightPathModel.Get(data)
	if err == nil && shortestPath != nil {
		return shortestPath, nil
	}

	// filter flight schedules
	data.Schedules = filterFlightSchedules(data.Schedules, data.PreferredTime)

	// convert schedules array into graph
	scheduleGraph := generateGraphOfSchedules(data.Schedules)

	// generate origin and destination city parameters
	origin := flightpath.ScheduleDetail{
		City: data.TripPlan.StartCity,
	}
	destination := flightpath.ScheduleDetail{
		City: data.TripPlan.EndCity,
	}

	// execute dijkstra's algorithm to get array of paths from origin to destination
	shortestDuration, paths := scheduleGraph.GetShortestPaths(origin, destination)

	// select the better paths among shortest paths
	shortestPath, err = getShortestPath(shortestDuration, paths)
	if err != nil {
		return nil, err
	}

	// save this shortest path in redis
	_ = c.Dao.FlightPathModel.Put(shortestPath, data)
	return shortestPath, nil
}

// generateGraphOfSchedules converts flight schedules into graph data structure
func generateGraphOfSchedules(schedules []*flightpath.FlightDetail) *graph {
	graph := newGraph()
	for _, schedule := range schedules {
		duration := schedule.Arrival.Timestamp - schedule.Departure.Timestamp
		graph.addEdge(*schedule.Departure, *schedule.Arrival, duration)
	}

	return graph
}

// filterFlightSchedules filters flight schedule by cutoffTimestamp i.e. returns flights that have departure time after cutoffTimestamp
func filterFlightSchedules(schedules []*flightpath.FlightDetail, cutoffTimestamp int64) []*flightpath.FlightDetail {
	if cutoffTimestamp == 0 {
		return schedules
	}

	filteredSchedules := make([]*flightpath.FlightDetail, 0)
	for _, schedule := range schedules {
		if schedule.Departure.Timestamp > cutoffTimestamp {
			filteredSchedules = append(filteredSchedules, schedule)
		}
	}
	return filteredSchedules
}

// getShortestPath returns one single most relevant flight path among all the shortest path
func getShortestPath(shortestDuration int64, paths map[int64][][]flightpath.ScheduleDetail) ([]flightpath.ScheduleDetail, error) {
	if paths == nil || len(paths) == 0 {
		return nil, errors.New(errorconsts.NoFlightsAvailable)
	}
	shortestPaths := paths[shortestDuration]
	if shortestPaths == nil || len(shortestPaths) == 0 {
		return nil, errors.New(errorconsts.NoFlightsAvailable)
	}

	// if only path available then it is the shortest path
	if len(shortestPaths) == 1 {
		return shortestPaths[0], nil
	}

	// there can be multiple shortest path which can have same number of stops
	// but we are picking the first one, since the api response can only contain one flight path
	shortestPathIndex := 0
	shortestPathLength := len(shortestPaths[0])
	for index, path := range shortestPaths {
		if len(path) < shortestPathLength {
			shortestPathLength = len(path)
			shortestPathIndex = index
		}
	}

	return shortestPaths[shortestPathIndex], nil
}
