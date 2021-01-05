package worker

/**
* @Author: super
* @Date: 2021-01-05 15:01
* @Description:
**/

type ParserFunc func(contents []byte, url string) ([]string, error)

type Parser interface {
	Parse(contents []byte, url string) ([]string, error)
}

type Request struct {
	Url    string
	Parser Parser
}

type FuncParser struct {
	parser    ParserFunc
	Name      string
}

func (f *FuncParser) Parse(contents []byte, url string) ([]string, error){
	return f.parser(contents, url)
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser:    p,
		Name:      name,
	}
}