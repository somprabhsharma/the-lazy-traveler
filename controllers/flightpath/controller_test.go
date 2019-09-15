package flightpath

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/somprabhsharma/the-lazy-traveler/constants/errorconsts"
	"github.com/somprabhsharma/the-lazy-traveler/entities/flightpath"
	"github.com/somprabhsharma/the-lazy-traveler/models"
	"testing"
)

func TestFlightPathController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "The Lazy Traveler Suite")
}

var _ = Describe("controllers", func() {
	Context("##flightpath", func() {
		controller := NewController(models.NewDao())
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
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "A",
						Timestamp: 2,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
				},
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 15,
					},
				},
			},
			TripPlan: &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			},
		}

		It("should throw error if start city and end city are same", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "A",
			}
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(shortestPath).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal(errorconsts.SameStartEndCity))
		})

		It("should throw error if no path found between start and end city", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "AA",
			}
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(shortestPath).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal(errorconsts.NoFlightsAvailable))
		})

		It("should return shortest path between start and end city", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			}
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(err).Should(BeNil())
			Expect(shortestPath).ShouldNot(BeNil())
			Expect(len(shortestPath)).To(Equal(2))
			Expect(shortestPath[0].City).To(Equal("A"))
			Expect(shortestPath[0].Timestamp).To(Equal(int64(1)))
			Expect(shortestPath[1].City).To(Equal("Z"))
			Expect(shortestPath[1].Timestamp).To(Equal(int64(10)))
		})

		It("should return shortest path between start and end city after preferred time", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			}
			data.PreferredTime = int64(2)
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(err).Should(BeNil())
			Expect(shortestPath).ShouldNot(BeNil())
			Expect(len(shortestPath)).To(Equal(3))
			Expect(shortestPath[0].City).To(Equal("A"))
			Expect(shortestPath[0].Timestamp).To(Equal(int64(2)))
			Expect(shortestPath[1].City).To(Equal("B"))
			Expect(shortestPath[1].Timestamp).To(Equal(int64(8)))
			Expect(shortestPath[2].City).To(Equal("Z"))
			Expect(shortestPath[2].Timestamp).To(Equal(int64(15)))
		})

		It("should throw error if any flight schedule provided does not have arrival time", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			}
			data.Schedules = []*flightpath.FlightDetail{
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
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "A",
						Timestamp: 2,
					},
				},
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 15,
					},
				},
			}
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(shortestPath).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal(errorconsts.InvalidFlightSchedule))
		})

		It("should throw error if any flight schedule provided does not have departure time", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			}
			data.Schedules = []*flightpath.FlightDetail{
				{
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 10,
					},
				},
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 15,
					},
				},
			}
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(shortestPath).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal(errorconsts.InvalidFlightSchedule))
		})

		It("should throw error if any flight schedule provided have invalid timestamp i.e. negative", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			}
			data.Schedules = []*flightpath.FlightDetail{
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "A",
						Timestamp: 1,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: -1,
					},
				},
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "A",
						Timestamp: 2,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
				},
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 15,
					},
				},
			}
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(shortestPath).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal(errorconsts.InvalidFlightSchedule))
		})

		It("should throw error if any flight schedule provided have invalid timestamp i.e. zero", func() {
			data.TripPlan = &flightpath.TripDetail{
				StartCity: "A",
				EndCity:   "Z",
			}
			data.Schedules = []*flightpath.FlightDetail{
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "A",
						Timestamp: 0,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 10,
					},
				},
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "A",
						Timestamp: 2,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
				},
				{
					Departure: &flightpath.ScheduleDetail{
						City:      "B",
						Timestamp: 8,
					},
					Arrival: &flightpath.ScheduleDetail{
						City:      "Z",
						Timestamp: 15,
					},
				},
			}
			shortestPath, err := controller.FindShortestFlightPath(data)
			Expect(shortestPath).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal(errorconsts.InvalidFlightSchedule))
		})
	})
})
