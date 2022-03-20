package query

type Field struct {
    name string
    field_type Fieldtype
}

type FieldType string
const (
	INT = "INT"
	STRING = "STRING"
	DATE = "DATE"
	DATETIME = ""
	BOOL = ""
)
