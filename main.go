package main

import (
	"alfred-yjiang/common"
	"alfred-yjiang/libs"
	"fmt"
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
	output := "查询中..."

	switch {
	case regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`).MatchString(inputString):
		output = libs.ActionIp(inputString)
	case regexp.MustCompile(`^\d{10,}$`).MatchString(inputString):
		output = libs.ActionInttotime(inputString)
	case regexp.MustCompile(`^\d{4}\-\d{1,2}\-\d{1,2}.*$`).MatchString(inputString):
		output = libs.ActionTimetoint(inputString)
	case regexp.MustCompile(`^now$`).MatchString(inputString):
		output = libs.ActionGetNow(inputString)
	default:
		output = "未识别的参数"
	}
	return output
}

func run() {
	args := wf.Args()
	input := ""
	output := ""

	if len(args) <= 0 {
		output = "请输入参数"
	} else {
		input = common.Implode(args, " ")
		output = inputApater(input)
	}

	item := wf.NewItem(output)

	//设置icon、副标题等标识
	icon := &aw.Icon{
		Value: "icon.png",
	}
	uid := "com.alfred.yjiang"
	item.Subtitle(input).Icon(icon).UID(uid).Valid(true)

	//设置Commnad键copy作用
	copyInput := fmt.Sprintf("复制结果:%s", output)
	item.Cmd().Subtitle(copyInput).Arg("copy").Valid(true)

	//设置Ctrl键search作用
	searchInput := fmt.Sprintf("搜索结果:%s", input)
	item.Ctrl().Subtitle(searchInput).Arg("search").Valid(true)

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
