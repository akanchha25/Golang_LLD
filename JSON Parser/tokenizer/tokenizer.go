package tokenizer

import "json_parser/data"

type Tokenizer interface {
	Tokenize(string) []data.KeyValuePair
}