package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const (
	ipUrl = "http://ip-api.com/json/"
)

func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	return m, nil
}

func checkIp(ipString string) bool {
	regex := `\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`
	match, _ := regexp.MatchString(regex, ipString)
	return match
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("缺少参数IP地址")
		os.Exit(1)
	}

	ip := os.Args[1]
	if checkIp(ip) == false {
		fmt.Println("ip地址格式有误")
		os.Exit(1)
	}

	url := ipUrl + ip
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	data, _ := JsonToMap(string(body))
	//fmt.Println(data["status"])

	output := "获取ip地址位置错误"
	if data["status"] == "success" {
		output = fmt.Sprintf("ip:%v 地理位置:%v %v %v", ip, data["country"], data["regionName"], data["city"])
	}
	fmt.Println(output)
}