package engine

type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct{
	Requests []Request
	Itmes []interface{}
}

func NilParser([]byte) ParseResult{
	return ParseResult{}
}