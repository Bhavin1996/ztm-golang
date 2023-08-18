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

func displayServerInfo(a map[string]int) {
	totalServers := len(a)
	totalOnline := 0
	toatalOffline := 0
	totalMaintenance := 0
	totalRetired := 0

	for _, v := range a {
		switch {
		case v == 0:
			{
				totalOnline += 1
				break
			}
		case v == 1:
			{
				toatalOffline += 1
				break
			}
		case v == 2:
			{
				totalMaintenance += 1
				break
			}
		case v == 3:
			{
				totalRetired += 1
				break
			}
		}
	}
	fmt.Println("Total servers are", totalServers)
	fmt.Println("Total server online are #", totalOnline, "\nTotal server Offline are #", toatalOffline, "\nTotal server Maintainance #", totalMaintenance, "\nTotal server retired #", totalRetired)

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverInfo := make(map[string]int)
	for i := 0; i < len(servers); i++ {
		serverInfo[servers[i]] = Online
	}
	for key, value := range serverInfo {
		fmt.Println(key, value)
	}
	displayServerInfo(serverInfo)
}
