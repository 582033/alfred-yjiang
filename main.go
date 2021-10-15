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
	var output string

	switch {
	case regexp.MustCompile(`^\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}$`).MatchString(inputString):
		output = libs.ActionIp(inputString)
	case regexp.MustCompile(`^\d{10,}$`).MatchString(inputString):
		output = libs.ActionInttotime(inputString)
	case regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}$`).MatchString(inputString):
		output = libs.ActionTimetoint(inputString)
	case regexp.MustCompile(`^now$`).MatchString(inputString):
		output = libs.ActionGetNow(inputString)
	default:
		output = "查询中..."
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
