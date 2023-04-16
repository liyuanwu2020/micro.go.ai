package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 101; i < 300; i++ {
		wg.Add(1)
		go func(num int) {
			title := getWebTitle(num)
			if title != "" {
				fmt.Printf("%s%d %s\n", "", num, title)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println("memory init")
}

func getWebTitle(i int) string {
	// 发送 GET 请求
	resp, err := http.Get("" + strconv.Itoa(i))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取网页内容
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)

	// 匹配指定 HTML 标签的正则表达式
	re := regexp.MustCompile(`<title[^>]*>(.*?)</title>`)

	// 提取标签内内容
	matches := re.FindAllStringSubmatch(bodyString, -1)
	for _, match := range matches {
		return match[1]
	}
	return ""
}
