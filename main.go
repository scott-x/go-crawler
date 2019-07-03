package main

import (
	"fmt"
	"github.com/scott-x/go-crawler/engine"
	"github.com/scott-x/go-crawler/zhenhun/parser"
)

func main(){
	fmt.Println("running.....")
	engine.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}

