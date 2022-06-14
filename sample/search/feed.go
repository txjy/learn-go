package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

//Feed包含需要处理的数据源的信息
type Feed struct {
	Name string `json:"site`
	URI  string `json:"link`
	Type string `json:"type"`
}

//读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {

	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	//函数返回时关闭文件
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
