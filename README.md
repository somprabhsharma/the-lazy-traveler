
# The Lazy Traveler

A Go app, which finds optimum flight path between two cities.
The app uses Dijkstra's algorithm to calculate the shortest path between two cities.
The Dijkstra's algorithm is implemented by using graph and heap data structures.
There are few assumptions made while implementing the algorithm, which are:
- There can be only one flight with combination of departure city, departure time, arrival city, arrival time. If there are multiple such flights in the request, they all are treated as one.
- If more than one paths are found with same duration, then the shortest path is selected based on the number of cities involved in the path.
- If more than on paths are found with same duration and same number of cities, then first path is selected as shortest path.
- Since there is type in the document provided for the key `prefered_time`, the app uses `preferred_time` key instead.
- If there is gap between arrival and departure time of a flight at an intermediate city then the gap is also included in the duration to calculate shortest path.
- The preferred time is inclusive i.e any flight that is after or at preferred time will be considered.

**APIs**
----
**Find Shortest Flight Path for Lazy Jack**

Returns shortest flight path between origin and destination city for lazy jack using dijkstra's algorithm.

* **URL**

  `/the-lazy-traveler/api/1.0/lazy_jack`

* **Method:**

  `POST`
  
*  **URL Params**

   None

* **Body Params**

  **Required:**
  ```
  {
      "schedules": [
         {
           "departure": {
             "city": "A",
             "timestamp": 2
           },
           "arrival": {
             "city": "Z",
             "timestamp": 10
           }
         }
      ],
      "trip_plan": {
        "start_city": "A",
        "end_city": "Z"
      }
  }
  ```
  **Optional**
  `"preferred_time": 1`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
    ```
    {
        "flight_plan": [
            {
                "city": "A",
                "timestamp": 2
            },
            {
                "city": "Z",
                "timestamp": 10
            }
        ]
    }
    ```
 
* **Error Response:**

  * **Code:** 400 BAD REQUEST <br />
    **Content:**
    ```
    {
          "message": "Invalid request. Please provide all required parameters in the request.",
          "code": 101
    }
    ```

## Built With
* [Gin](https://github.com/gin-gonic/gin) - The web framework
* [Dep](https://github.com/golang/dep) - Dependency Management

## Contributors
* **Som Prabh Sharma** - *Initial Setup, Lazy Jack API* - [somprabhsharma](https://github.com/somprabhsharma)

## Acknowledgments
* Cheers to packages that helped **the-lazy-traveler** achieve everything it has. 