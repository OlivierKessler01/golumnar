package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
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
