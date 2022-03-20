package query

import (
    "fmt"
    "strings"
    "os"
    "github.com/OlivierKessler01/golumnar/filesystem"
)

/**
* Uses the filesystem library to create the metadata file and the different field files
*/
func createTable(name string, fields []Fields)
{
    filesystem.Createfile(name + ".table_metadata")
    //TODO : write metadata in the file
    for index, field := range fields {
        filesystem.Createfile(name+'_'+field.name)
    }
}

func executeQuery(string query) {
    result = checkSynthax(query)

    if result != nil {
        fmt.Fprintln(os.Stderr, result)
    }

    var fields []Fields = parseCreateFields(query)

    queryComponents := strings.Fields(cmdString)
    operation = queryComponents[0]

    if operation == "CREATE" {
        createTable(queryComponents[1], fields)
    }
}
