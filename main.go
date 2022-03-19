package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
)

func main() {
    //Main error catching for the whole system
    defer func(){
        if r != nil {
            var configuration Configuration = LoadConfiguration()

           // If the file doesn't exist, create it, or append to the file
            f, err := os.OpenFile(configuration.LogfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                log.Fatal(err)
            }
            if _, err := f.Write([]byte(r)); err != nil {
                log.Fatal(err)
            }
            if err := f.Close(); err != nil {
                log.Fatal(err)
            }
        }
    }()

    // STARTS THE SHELL AND WAIT FOR INPUT
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("$ ")
        cmdString, err := reader.ReadString('\n')

        if err != nil {
            fmt.Println(os.Stderr, err)
        }

        cmdString = strings.TrimSuffix(cmdString, "\n")

        if cmdString == "exit" {
            return
        }

        result, err := execute(smdString)

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}
