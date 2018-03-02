package version

import (
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

var VersionString = "1.0.0"
var CommitSha string = "No Clue!!!"

// GetVersion returns the version info from the parsed string
func GetVersion() plugin.VersionType {
	versionInfo := strings.Split(VersionString, ".")
	return plugin.VersionType{
		Major: getNumber(versionInfo[0], 1),
		Minor: getNumber(versionInfo[1], 0),
		Build: getNumber(versionInfo[2], 0),
	}
}

func getNumber(num string, defaultValue int) int {
	if val, err := strconv.Atoi(num); err == nil {
		return val
	}
	return defaultValue
}
