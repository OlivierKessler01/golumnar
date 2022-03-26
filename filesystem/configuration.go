package filesystem

import (
	"os"
	"log"
	"strings"
	"regexp"
	"fmt"
)

type Configuration struct {
	LogfilePath string
}

var configuration_cache *Configuration

/**
* Reads the configuration file and load it into a struct.
* This acts as a Singleton in OOP, meaning a global variable is maintained.
* If the configuration as already been parsed during the execution of the program
* return a pointer to this previous struct.
*/
func LoadConfiguration() *Configuration {
	if configuration_cache == nil {
		fmt.Println("Configuration never parsed, parsing...")
		data, err := os.ReadFile("/etc/golumnar/golumnar.cnf")

		if err != nil {
			log.Fatalf("Unable to read configuration file: %v", err)
		}

		var buffer []byte = make([]byte, 0)
		configuration_cache = &Configuration{""}

		for _, v := range data {
			if string(v) == "\n" {
				matched, _ := regexp.Match(`\S+="\S+"`, buffer)
				if matched == false {
					panic("Wrong config line format")
				}

				var buffer_string string = string(buffer)

				if strings.Contains(buffer_string, "log_file") {
					configuration_cache.LogfilePath = buffer_string[strings.Index(buffer_string, "=") + 2:len(buffer_string)-1]
				}
				buffer = make([]byte, 10)
			} else {
				buffer = append(buffer, v)
			}
		}

		if configuration_cache.LogfilePath == "" {
			panic(`Your config file need a -- log_file="" -- entry to specify a log file.`)
		}

		return configuration_cache
	} else {
		fmt.Println("Configuration already parsed, returning cache.")
		return configuration_cache
	}
}
