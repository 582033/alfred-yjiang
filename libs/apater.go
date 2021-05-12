package libs

import (
	"alfred-yjiang/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	ipUrl = "http://ip-api.com/json/"
)

func ActionIp(ip string) string {
	url := ipUrl + ip + "?lang=zh-CN"
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	data, _ := common.JsonToMap(string(body))
	//fmt.Println(data["status"])

	output := "获取ip地址位置错误"
	if data["status"] == "success" {
		output = fmt.Sprintf("ip:%v 地理位置:%v %v %v", ip, data["country"], data["regionName"], data["city"])
	}
	return output
}

func ActionTimetoint(timeStr string) string {
	//时间转换模板
	timeLayout := "2006-01-02"
	match, _ := regexp.MatchString(`\d{2}:\d{2}:\d{2}`, timeStr)
	if match {
		timeLayout = "2006-01-02 15:04:05"
	}
	//获取本地location
	loc, _ := time.LoadLocation("Local")                         //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                         //转化为时间戳 类型是int64
	return strconv.Itoa(int(sr))
}

func ActionInttotime(timeStr string) string {
	i, _ := strconv.ParseInt(timeStr, len(timeStr), 64)
	return time.Unix(i, 0).Format("2006-01-02 15:04:05")
}

func ActionGetNow(timeStr string) string {
	timeStamp := time.Now().Unix()
	return strconv.Itoa(int(timeStamp))
}
