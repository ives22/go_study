package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

// ES demo

// Student ...
type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	// 1. 初始化连接，得到一个client
	client, err := elastic.NewClient(
		elastic.SetURL("http://120.24.222.91:9201"),
		elastic.SetSniff(false), // 将sniff设置为false后，便不会自动转换地址（用于连接docker部署的单节点es时）
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect to es success.")

	p1 := Student{Name: "xiaobai", Age: 18, Married: false}
	// 链式操作
	put1, err := client.Index().
		Index("student").
		Type("doc").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexd student %s to index %s, type:%s\n", put1.Id, put1.Index, put1.Type)
}
