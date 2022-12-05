package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"124.71.33.240:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Printf("create etcd lease failed, err:%v\n", err)
		return
	}

	// 创建一个key，并和租约绑定起来
	_, err = cli.Put(context.TODO(), "testkey", "testvalue", clientv3.WithLease(resp.ID))
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

	// 模拟取值，每1秒去get一下数据，看5s后数据是否过期
	for {
		resp, err := cli.Get(context.TODO(), "testkey")
		if err != nil {
			fmt.Printf("get from etcd failed, err:%v\n", err)
			return
		}
		// 判断是否还有数据
		if resp.Count == 0 {
			fmt.Println("failed to get data, testkey expired.")
			break
		}
		fmt.Printf("get data success, kvs:%s\n", resp.Kvs)
		time.Sleep(time.Second)
	}
}
