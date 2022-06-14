package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {

	//获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//创建一个chan接受匹配后的结果
	results := make(chan *Result)

	//构建一个waitGroup处理所有的数据源
	var waitGroup sync.WaitGroup

	//每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	//为每一个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["dafault"]
		}

		go func(matcher Matcher, feed *Feed) {

			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}
	//启动一个goroutine监控是否全部完成
	go func() {
		waitGroup.Wait()

		close(results)
	}()

	//显示返回结果
	Display(results)
}

//Register调用时，注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registerd")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
