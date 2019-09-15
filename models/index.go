package models

import "github.com/somprabhsharma/the-lazy-traveler/models/redis"

// Dao dao struct
type Dao struct {
	Cache           *redis.Client
	FlightPathModel *flightPathModel
}

// NewDao creates instance of Dao
func NewDao() *Dao {
	redisClient := redis.NewClient()
	return &Dao{
		Cache:           redisClient,
		FlightPathModel: newFlightPathModel(redisClient),
	}
}
