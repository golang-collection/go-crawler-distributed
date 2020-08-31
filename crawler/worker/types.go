package worker

/**
* @Author: super
* @Date: 2020-08-16 08:04
* @Description:
**/

type ParserFunc func(contents []byte, queueName string, url string)

type Parser interface {
	Parse(contents []byte, url string)
}

type Request struct {
	Url    string
	Parser Parser
}

type FuncParser struct {
	parser    ParserFunc
	QueueName string
	Name      string
}

func (f *FuncParser) Parse(contents []byte, url string) {
	f.parser(contents, f.QueueName, url)
}

func NewFuncParser(p ParserFunc, mqName string, name string) *FuncParser {
	return &FuncParser{
		parser:    p,
		QueueName: mqName,
		Name:      name,
	}
}
