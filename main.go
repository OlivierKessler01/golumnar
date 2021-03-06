package main

import (
	"bufio"
	"fmt"
	"github.com/OlivierKessler01/golumnar/filesystem"
	"github.com/OlivierKessler01/golumnar/query"
	"log"
	"os"
	"strings"
)

func main() {
	var configuration *filesystem.Configuration
	configuration = filesystem.LoadConfiguration()

	// If the file doesn't exist, create it, or append to the file
	log_file, err := os.OpenFile(configuration.LogfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Can't open log file :")
		fmt.Println(os.Stderr, err)
	}

	log.SetOutput(log_file)

	// STARTS THE SHELL AND WAIT FOR INPUT
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Golumnar database v0.1")
	fmt.Println("Logging in :", configuration.LogfilePath)
	fmt.Println(`Type "exit" to stop`)

	for {
		fmt.Print("$> ")
		cmdString, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
			fmt.Println(os.Stderr, err)
		}

		cmdString = strings.TrimSuffix(cmdString, "\n")

		//The user types "exit", exit the shell
		if cmdString == "exit" {
			return
		}

		//Execute Query
		result, err := query.Execute(cmdString)
		fmt.Println(result)

		if err != nil {
			log.Fatal(err)
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
