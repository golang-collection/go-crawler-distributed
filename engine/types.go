package engine

type ParserFunc func(
	contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
