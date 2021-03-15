package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Useage: ./v <name>")
	}

	resp, err := http.Get("http://shaobo.fun:8000/api/SearchPriceAPI?s="+os.Args[1])
	if err != nil {
		fmt.Println("网络错误")
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}