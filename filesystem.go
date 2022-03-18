package filesystem

import (
    "os"
    "fmt"
)

func Createfile(name string) bool {
    myfile, e := os.Create(name)
    if e != nil {
        fmt.Println("Error")
        return false
    }
    myFile.close()
    return true
}

func Readline() {
}

func Loadfileinmemory() {
}
