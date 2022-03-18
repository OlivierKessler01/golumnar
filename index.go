package main

import (
    "fmt"
    "bufio"
    "os"
    "os/exec"
    "strings"
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

        cmdComponents := strings.Fields(cmdString)
        cmd := exec.Command(cmdComponents[0], cmdComponents[1:]...)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
        cmd.Run()

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}
