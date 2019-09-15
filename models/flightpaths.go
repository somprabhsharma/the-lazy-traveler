package models

import (
	"encoding/base64"
	"encoding/json"
	"github.com/somprabhsharma/the-lazy-traveler/entities/flightpath"
	"github.com/somprabhsharma/the-lazy-traveler/models/redis"
	"github.com/somprabhsharma/the-lazy-traveler/utils/logger"
	"strconv"
	"time"
)

const (
	flightPathSuffix = "-flight-path"
	flightPathTTL    = 24 * time.Hour
)

type flightPathModel struct {
	Cache *redis.Client
}

func newFlightPathModel(redis *redis.Client) *flightPathModel {
	return &flightPathModel{
		Cache: redis,
	}
}

// Put puts shortest path result in cache
func (t *flightPathModel) Put(shortestPath []flightpath.ScheduleDetail, data flightpath.LazyJackRequest) error {
	// generate key from input data
	key := generateCacheKey(data)

	// stringify the shortest path result
	shortestPathBytes, err := json.Marshal(shortestPath)
	if err != nil {
		logger.Warn("lazy-jack", "error while saving shortest path data in cache for key: "+key, err)
		return nil
	}

	// save it in cache
	return t.Cache.Put(key, string(shortestPathBytes), flightPathTTL)
}

// Get gets shortest path result from cache
func (t *flightPathModel) Get(data flightpath.LazyJackRequest) ([]flightpath.ScheduleDetail, error) {
	// generate key from input data
	key := generateCacheKey(data)

	// save it in cache
	value, err := t.Cache.Get(key)
	if err != nil {
		logger.Warn("lazy-jack", "error while getting shortest path data from cache for key: "+key, err)
		return nil, err
	}

	// unmarshal data obtained from cache
	var shortestPath []flightpath.ScheduleDetail
	err = json.Unmarshal([]byte(value), &shortestPath)
	if err != nil {
		logger.Warn("lazy-jack", "error while unmarshalling shortest path data obtained from cache for key: "+key, err)
		return nil, err
	}

	return shortestPath, nil
}

// generateCacheKey generates unique key for input data
func generateCacheKey(data flightpath.LazyJackRequest) string {
	key := data.TripPlan.StartCity + "_" + data.TripPlan.EndCity + "_" + strconv.FormatInt(data.PreferredTime, 10)
	base64key := base64.StdEncoding.EncodeToString([]byte(key))
	return base64key + flightPathSuffix
}
