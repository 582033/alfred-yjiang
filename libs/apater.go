package apater

import (
	"common"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
