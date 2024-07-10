package parser

import "json_parser/data"


type NaiveJsonParser struct{}

func(n *NaiveJsonParser) Parse(jsonString string) (*data.JSON, error) {
	
}

func(n *NaiveJsonParser) ToString(json data.JSON) (string, error) {
	
}