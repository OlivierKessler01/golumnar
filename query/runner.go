package query

import (
    "fmt"
    "os"
    "github.com/OlivierKessler01/golumnar/filesystem"
)

/**
* Uses the filesystem library to create the metadata file and the different field files
*/
func createTable(name string, fields []Field) {
    fmt.Println("Creating file", name + ".table_metadata")
    filesystem.Createfile(name + ".table_metadata")
    //TODO : write metadata in the file
    for _, field := range fields {
        fmt.Println("Creating column file, name", "_", field.name)
        filesystem.Createfile(name+"_"+field.name)
    }
}

/**
* Execute a query :
* - Check the synthax
* - Apply the changes on the filesystem
*/
func Execute(query string) (bool, error) {
    result, operation, name := checkSynthax(query)

    if result == false {
        fmt.Fprintln(os.Stderr, result)
    }

    if *operation == 0 {
        var fields []Field = parseCreateFields(query)
        createTable(*name, fields)
    }

    return true, nil
}
