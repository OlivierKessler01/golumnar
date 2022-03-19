package filesystem

import (
    "os"
    "fmt"
	"strings"
)

type Configuration struct {
	string LogfilePath
}

func LoadConfiguration() Configuration {
	dat, err := os.ReadFile("/etc/golumnar/golumnar.cnf")
    check(err)
	components := strings.Fields(cmdString)
	var configuration Configuration;

	for i, v range components {
		if strings.Contains(v, "log_file") {
			configuration.LogfilePath = v[strings.IndexByte(v, '='):]
		}
	}

	return configuration
}
