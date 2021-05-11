package main

import (
	"fmt"
	"os"
	"regexp"
)

type regexRow struct {
	regex    string
	funcName string
}

func inputApater(inputString string) string {
	output := "未识别的参数"
	regexSlice := []regexRow{
		{regex: `\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`, funcName: "actionIp"},
		{regex: `\d+`, funcName: "timestamp"},
		{regex: `\d{4}-\d{1,2}-\d{1,2}`, funcName: "timestamp"},
		{regex: `now`, funcName: "timestamp"},
	}
	for _, ipObj := range regexSlice {
		//fmt.Println(ipObj.regex)
		match, _ := regexp.MatchString(ipObj.regex, inputString)
		fmt.Println(match)
		if match {
			output = common.Call(ipObj.funcName, inputString)
		}
	}
	return output
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("请输入参数")
		os.Exit(1)
	}

	input := os.Args[1]
	output := inputApater(input)
	fmt.Println(output)
}
