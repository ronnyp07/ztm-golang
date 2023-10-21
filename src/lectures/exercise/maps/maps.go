//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//Create the server map
	serverMap := make(map[string]int)

	// Assign Onlie status to all serverMap
	for _, item := range servers {
		serverMap[item] = Online
	}

	serverStatus(serverMap)

	//  - change server status of `darkstar` to `Retired`
	serverMap["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverMap["aiur"] = Offline
	//  - call display server info function
	serverStatus(serverMap)
	//  - change server status of all servers to `Maintenance`
	for key, _ := range serverMap {
		serverMap[key] = Maintenance
	}
	//  - call display server info function
	serverStatus(serverMap)
}

func serverStatus(server map[string]int) {
	stats := make(map[int]int)
	for _, value := range server {
		switch value {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		}
	}

	fmt.Println("The total of server is: ", len(server))
	fmt.Println("Online: ", stats[Online], "Offline: ", stats[Offline], "Maitence: ", stats[Maintenance], "Retired: ", stats[Retired])

}
