package search

//defaultMatcher实现了默认匹配器
type defaultMatcher struct{}

//init将默认匹配器注册到程序中
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

//Search实现了默认匹配器的行为
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
