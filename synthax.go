package main

import "fmt"

type Operation string
const (
	DATABASE = "CREATE"
    SELECT = "SELECT"
    UPDATE = "UPDATE"
    DELETE = "DELETE"
    TRUNCATE = "TRUNCATE"
)

type Keyword string
const (
    //SQL keywords
    FROM = "FROM"
    VALUES = "VALUES"
    WHERE = "WHERE"
)

type Query struct {
        operation Operation
        queryParts []QueryPart
}

type QueryPart struct {
        keyword Keyword
        rest string
}






