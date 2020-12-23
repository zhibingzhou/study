package common

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{} //列表
}



var GoUrl = "https://www.xiangqinwang.cn"



