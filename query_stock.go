package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"strconv"
)

type Config struct {
	Codes []string `json:"codes"`
}

const (
	path = "config.json"
	sinaurl = "http://hq.sinajs.cn/list="
)

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func loadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stream, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return stream
}

func getMarket(code string) (string, error) {
	var market string
	var queryUrl string
	if strings.HasPrefix(code, "00") {
		queryUrl = sinaurl + "sz" + code
	} else if strings.HasPrefix(code, "60") {
		queryUrl = sinaurl + "sh" + code
	} else {
		return market, errors.New("未识别的代码")
	}
	
	resp, err := http.Get(queryUrl)
	if err != nil {
		return market, errors.New("网络错误")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return market, errors.New("数据流读取失败")
	}
	market = string(body)
	return market, nil
}

func parseData(content string) string {
	items := strings.Split(content, ",")
	pre_close, _ := strconv.ParseFloat(items[2], 64)
	now_price, _ := strconv.ParseFloat(items[3], 64)
	rose := fmt.Sprintf("%.3f", ((now_price-pre_close) / pre_close) * 100)
	return rose
}

func main() {
	codes := []string{"002439"}
	if fileExist(path) == true {
		var config Config
		stream := loadFile(path)
		err := json.Unmarshal([]byte(stream), &config)
		if err != nil {
			fmt.Println("json文件解析错误")
			return
		}
		codes = config.Codes
	}

	var stockRoseList []string
	for _, code := range codes {
		content, err := getMarket(code)
		if err != nil {
			fmt.Println("数据获取失败")
			return
		}
		stockRoseList = append(stockRoseList, parseData(content))
	}
	fmt.Println(stockRoseList)
}
