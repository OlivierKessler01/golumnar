package query

import (
    "regexp"
    "fmt"
)

type Operation int
const (
    createOp Operation = iota
    selectOp
)

const (
	CREATE_SYNTHAX = `CREATE\s+TABLE\s+'(.*)'\s+\((.*)\)\s*;`
    SELECT_SYNTHAX = "SELECT"
    UPDATE_SYNTHAX = "UPDATE"
    DELETE_SYNTHAX = "DELETE"
    TRUNCATE_SYNTHAX = "TRUNCATE"
)

/**
* Check the Synthax of the query against regexps
* Returns the type of operation (CREATE, SELECT) etc
*/
func checkSynthax(query string) (bool, *Operation, *string) {
    matched, _ := regexp.MatchString(CREATE_SYNTHAX, query)
    name := "nothing"
    if matched == true {
        return true, &createOp, name
    } else {
        fmt.PrintLn("Wrong query synthax.")
        return false, nil, nil
    }
}

func parseCreateFields(query string) []Field {
    var fields []Field
    fields = make([]Field, 1, 1)
    return fields
}






