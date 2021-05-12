package main

import (
	"alfred-yjiang/libs"
	"regexp"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

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
		{regex: `^\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}$`, action: libs.ActionIp(inputString)},
		{regex: `^\d{10,}$`, action: libs.ActionInttotime(inputString)},
		{regex: `^\d{4}-\d{1,2}-\d{1,2}$`, action: libs.ActionTimetoint(inputString)},
		{regex: `^now$`, action: libs.ActionGetNow(inputString)},
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

func run() {
	args := wf.Args()
	output := ""
	if len(args) <= 0 {
		output = "请输入参数"
	} else {
		input := args[0]
		output = inputApater(input)
	}
	wf.NewItem(output)
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
