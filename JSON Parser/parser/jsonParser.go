package parser


import (
	"json_parser/data"
)

type JsonParser interface {
	Parse(jsonText string) (*data.JSON, error)
	ToString(jsonObj *data.JSON) (string, error)
}
