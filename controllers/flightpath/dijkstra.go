package flightpath

import (
	"container/heap"
	"github.com/somprabhsharma/the-lazy-traveler/entities/flightpath"
	"strconv"
)

// directPath is a direct path struct between two nodes with duration
type directPath struct {
	duration int64
	nodes    []flightpath.ScheduleDetail
}

// define a type path, that is array of individual direct paths
type path []directPath

// Len gets length of path
func (p path) Len() int {
	return len(p)
}

// Less compares two paths values and tells if a path is less than another path
func (p path) Less(i, j int) bool {
	return p[i].duration < p[j].duration
}

// Swap swaps two paths
func (p path) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
	// p[i], p[j] = p[j], p[i]
}

// Push adds a direct path to path array
func (p *path) Push(x interface{}) {
	*p = append(*p, x.(directPath))
}

// Pop removes the top direct path from path array
func (p *path) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

// heapTree is specialized tree-based data structure that is a complete tree which satisfies the heap property
type heapTree struct {
	Values *path
}

// newHeap creates new heap tree of paths
func newHeap() *heapTree {
	return &heapTree{Values: &path{}}
}

// push adds a new path to the heap tree
func (h *heapTree) push(p directPath) {
	heap.Push(h.Values, p)
}

// pop removes the top path from the heap tree
func (h *heapTree) pop() directPath {
	i := heap.Pop(h.Values)
	return i.(directPath)
}

// edge is a path with details of end node
type edge struct {
	Schedule              flightpath.ScheduleDetail
	Duration              int64
	OriginFlightTimestamp int64
	Reverse               bool //reverse flag to ignore reverse directional edges
}

// graph is a graph data structure to store the various schedules from given origin i.e. neighbouring nodes of a vertex
type graph struct {
	Schedules map[string][]edge
}

// newGraph creates a new graph
func newGraph() *graph {
	return &graph{Schedules: make(map[string][]edge)}
}

// addEdge adds an edge to the graph
func (g *graph) addEdge(origin, destination flightpath.ScheduleDetail, duration int64) {
	g.Schedules[origin.City] = append(g.Schedules[origin.City], edge{Schedule: destination, Duration: duration, OriginFlightTimestamp: origin.Timestamp, Reverse: false})
	g.Schedules[destination.City] = append(g.Schedules[destination.City], edge{Schedule: origin, Duration: duration, OriginFlightTimestamp: destination.Timestamp, Reverse: true})
}

// getEdges gets all the edges of given node
func (g *graph) getEdges(node string) []edge {
	return g.Schedules[node]
}

// getShortestPaths gets the shortest path between origin and destination
func (g *graph) getShortestPaths(origin, destination flightpath.ScheduleDetail) (int64, map[int64][][]flightpath.ScheduleDetail) {
	// create a heap tree starting with the origin city as first node
	heapT := newHeap()
	heapT.push(directPath{duration: 0, nodes: []flightpath.ScheduleDetail{origin}})

	// list of visited nodes to keep track of node
	visitedNode := make(map[string]bool)

	// since there can be multiple shortest path, we will decide later which path is better, hence keeping an array of shortest paths
	shortestPaths := make(map[int64][][]flightpath.ScheduleDetail, 0)
	shortestDuration := int64(0)

	for len(*heapT.Values) > 0 {
		// find the nearest node that is yet to be visitedNode
		p := heapT.pop()
		node := p.nodes[len(p.nodes)-1]

		// if the node is visited then continue
		if visitedNode[node.City+"_"+strconv.FormatInt(node.Timestamp, 10)] {
			continue
		}

		// if we have traversed the complete tree i.e. the last node in the heap is origin node then we have found our shortest path
		// add this shortest path to the shortest paths array
		// update the value of shortestDuration
		if node.City == destination.City {
			if len(shortestPaths) != 0 && p.duration > shortestDuration {
				continue
			}
			storedPaths := shortestPaths[p.duration]
			storedPaths = append(storedPaths, p.nodes)
			shortestPaths[p.duration] = storedPaths
			shortestDuration = p.duration
			continue
		}

		// get all the edges of the given node from the graph
		edges := g.getEdges(node.City)
		originFlightTimestamp := int64(0)
		for _, e := range edges {
			// if any node at the end of edge is not visited yet, then add it to heap with the path duration
			if !visitedNode[e.Schedule.City+"_"+strconv.FormatInt(e.Schedule.Timestamp, 10)] {
				if node.City != origin.City && node.Timestamp > e.OriginFlightTimestamp {
					continue
				}

				gapBetweenFlights := int64(0)
				// handle specific case of origin schedule
				// as there can be multiple flights from the origin, and we don't know from where to start
				// so we added origin in the heap with 0 flight timestamp
				// hence, we are updating the timestamp to its correct value before adding it to the heap
				var updatedNodes []flightpath.ScheduleDetail
				var updatedPNodes []flightpath.ScheduleDetail
				if node.City == origin.City && node.Timestamp == 0 {
					for _, v := range p.nodes {
						if v.City == origin.City && v.Timestamp == 0 {
							v.Timestamp = e.OriginFlightTimestamp
							originFlightTimestamp = e.OriginFlightTimestamp
						}
						updatedPNodes = append(updatedPNodes, v)
					}
				} else {
					updatedPNodes = p.nodes
				}

				// handle case when there is gap between arrival and departure in connecting cities
				if node.Timestamp > 0 && e.OriginFlightTimestamp-node.Timestamp > 0 {
					gapBetweenFlights = e.OriginFlightTimestamp - node.Timestamp
					updatedPNodes = append(updatedPNodes, flightpath.ScheduleDetail{
						City:      node.City,
						Timestamp: e.OriginFlightTimestamp,
					})
				}

				updatedNodes = append(updatedNodes, append(updatedPNodes, e.Schedule)...)
				if originFlightTimestamp == 0 {
					originFlightTimestamp = e.Schedule.Timestamp
				}

				// do not add reverse paths, this is to make sure that only directed paths are added to the heap
				if !e.Reverse {
					heapT.push(directPath{duration: p.duration + e.Duration + gapBetweenFlights, nodes: updatedNodes})
				}
				visitedNode[node.City+"_"+strconv.FormatInt(originFlightTimestamp, 10)] = true
			}
		}

		node.Timestamp = originFlightTimestamp
		visitedNode[node.City+"_"+strconv.FormatInt(node.Timestamp, 10)] = true
	}

	return shortestDuration, shortestPaths
}
