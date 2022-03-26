package query

type FieldType string
const (
	INT = "INT"
	STRING = "STRING"
	DATE = "DATE"
	DATETIME = ""
	BOOL = ""
)

type Field struct {
    name string
    field_type FieldType
}

