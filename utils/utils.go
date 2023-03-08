package utils

import (
	"errors"

	"github.com/antstalepresh/flightpath/types"
)

func GetTrackedPath(paths types.PathList) (types.Path, error) {
	if len(paths) == 1 && paths[0][0] == paths[0][1] {
		return types.Path{}, errors.New("same path")
	}

	var sortedPaths types.PathList
	forwardMap := make(map[string]string)  //map[source => destination]
	backwardMap := make(map[string]string) //map[destination => source]
	for _, data := range paths {
		forwardMap[data[0]] = data[1]
		backwardMap[data[1]] = data[0]
	}

	// find final destination airport starting from the first element of the array
	srcAirport := paths[0][0]
	desAirport := paths[0][1]
	for desAirport != "" {
		sortedPaths = append(sortedPaths, types.Path{srcAirport, desAirport})
		srcAirport = desAirport
		desAirport = forwardMap[desAirport]

		if len(sortedPaths) > len(paths) { // circular dependency is found
			return types.Path{}, errors.New("invalid path")
		}
	}

	// find original source airport starting from the first element of the array
	srcAirport = paths[0][0]
	desAirport = ""
	for srcAirport != "" {
		if desAirport != "" {
			sortedPaths = append(types.PathList{types.Path{srcAirport, desAirport}}, sortedPaths...)
		}
		desAirport = srcAirport
		srcAirport = backwardMap[srcAirport]

		if len(sortedPaths) > len(paths) { // circular dependency is found
			return types.Path{}, errors.New("invalid path")
		}
	}
	return types.Path{sortedPaths[0][0], sortedPaths[len(sortedPaths)-1][1]}, nil
}
