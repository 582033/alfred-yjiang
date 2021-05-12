package main

import (
	"alfred-yjiang/libs"
	"fmt"
	"os"
	"regexp"
)

const (
	ipUrl = "http://ip-api.com/json/"
)

type regexRow struct {
	regex  string
	action string
}

func inputApater(inputString string) string {
	output := "未识别的参数"

	regexSlice := []regexRow{
		{regex: `\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`, action: libs.ActionIp(inputString)},
		{regex: `\d+`, action: libs.ActionInttotime(inputString)},
		{regex: `\d{4}-\d{1,2}-\d{1,2}`, action: libs.ActionTimetoint(inputString)},
		{regex: `now`, action: libs.ActionGetNow(inputString)},
	}

	for _, regexObj := range regexSlice {
		//fmt.Println(ipObj.regex)
		match, _ := regexp.MatchString(regexObj.regex, inputString)
		//fmt.Println(match)
		if match {
			output = regexObj.action
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
