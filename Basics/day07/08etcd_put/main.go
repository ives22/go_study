package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3

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
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"/Users/liyanjie/Desktop/test/nginx.log","topic":"web_log"},{"path":"/Users/liyanjie/Desktop/test/redis.log","topic":"redis_log"}]`
	// value := `[{"path":"/Users/liyanjie/Desktop/test/nginx.log","topic":"web_log"},{"path":"/Users/liyanjie/Desktop/test/redis.log","topic":"redis_log"},{"path":"/Users/liyanjie/Desktop/test/mysql.log","topic":"mysql_log"}]`
	_, err = cli.Put(ctx, "/logagent/192.168.255.10/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
