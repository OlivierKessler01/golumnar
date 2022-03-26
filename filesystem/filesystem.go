package filesystem

import (
    "os"
    "fmt"
)

func Createfile(name string) bool {
    fmt.Println(configuration_cache)
    myFile, e := os.Create(name)
    if e != nil {
        fmt.Println("Error")
        return false
    }
    myFile.Close()
    return true
}

func Readline() {
}

func Loadfileinmemory() {
}
