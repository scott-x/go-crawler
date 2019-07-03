package parser

import (
	"regexp"

	"github.com/scott-x/go-crawler/engine"
)

const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult{
	re :=regexp.MustCompile(citylistRe)
	matches :=re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matches{
		result.Items=append(result.Items,"City: "+string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParserCity,
		})
	}
	return result
}