package search

import (
	"fmt"
	"log"
)

//Result保存搜索到的结果
type Result struct {
	Field   string
	Content string
}

//Matcher定义要实现的新搜索类型行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//Match函数为每一个数据源单独启动goroutine来执行这个函数
//并发的进行搜索

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//对特定的匹配器进行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//将结果写入chan
	for _, result := range searchResults {
		results <- result
	}
}

//Display从每个单独的goroutine接受到结果后在终端窗口输出
func Display(results chan *Result) {
	//chan会一直阻塞，直到有结果写入
	//chan一旦被关闭，for循环终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
