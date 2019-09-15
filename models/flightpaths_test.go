package models

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/somprabhsharma/the-lazy-traveler/entities/flightpath"
	"testing"
)

func TestFlightPaths(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "The Lazy Traveler Suite")
}

var _ = Describe("models", func() {
	Context("##flightpaths", func() {
		dao := NewDao()
		data := flightpath.LazyJackRequest{
			Schedules: []*flightpath.FlightDetail{
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "A",
						Timestamp: 1,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 10,
					},
				},
			},
			TripPlan: &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			},
		}

		shortestPath := []flightpath.ScheduleDetail{
			{
				City:      "A",
				Timestamp: 1,
			},
			{
				City:      "Z",
				Timestamp: 10,
			},
		}

		It("should save shortest path in redis cache", func() {
			_ = dao.FlightPathModel.Put(shortestPath, data)
			val, _ := dao.FlightPathModel.Get(data)
			Expect(val).ShouldNot(BeNil())
			Expect(len(val)).To(Equal(2))
			Expect(val[0].City).To(Equal("A"))
			Expect(val[0].Timestamp).To(Equal(int64(1)))
			Expect(val[1].City).To(Equal("Z"))
			Expect(val[1].Timestamp).To(Equal(int64(10)))
		})

		It("should get shortest path from redis cache", func() {
			_ = dao.FlightPathModel.Put(shortestPath, data)
			val, _ := dao.FlightPathModel.Get(data)
			Expect(val).ShouldNot(BeNil())
			Expect(len(val)).To(Equal(2))
			Expect(val[0].City).To(Equal("A"))
			Expect(val[0].Timestamp).To(Equal(int64(1)))
			Expect(val[1].City).To(Equal("Z"))
			Expect(val[1].Timestamp).To(Equal(int64(10)))
		})
	})
})
