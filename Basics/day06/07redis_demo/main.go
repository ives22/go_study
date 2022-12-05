package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	// 声明一个全局的redis变量
	rdb *redis.Client
	// Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	// 它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	ctx = context.Background()
)

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "124.71.33.240:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	// 验证是否连接到服务端
	ret, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("connect failed, err:%#v\n", err)
		return err
	}
	fmt.Printf("connect success, ping => %v\n", ret)
	return nil

}

// doCommand go-redis基本使用示例
func doCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := rdb.Get(ctx, "key1").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := rdb.Get(ctx, "key1")
	fmt.Println(cmder.Val()) // 获取值
	fmt.Println(cmder.Err()) // 获取错误

	// 执行执行命令获取错误
	err = rdb.Set(ctx, "key", 10, time.Hour).Err()
	if err != nil {
		fmt.Printf("set key err, err:%v\n", err)
		return
	}

	// 直接执行命令获取值
	value := rdb.Get(ctx, "key").Val()
	fmt.Println(value)
}

// doDemo rdb.Do 方法执行任意命令示例
func doDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行执行命令获取错误
	err := rdb.Do(ctx, "set", "key2", 10, "EX", 3600).Err()
	if err != nil {
		fmt.Printf("exec do failed, err:%v\n", err)
		return
	}

	// 执行命令获取结果
	val, err := rdb.Do(ctx, "get", "key2").Result()
	if err != nil {
		fmt.Printf("exec do failed, err:%v\n", err)
		return
	} else {
		fmt.Println(val)
	}
}

// getValueFromRedis redis.Nil判断
func getValueFromRedis(key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		// 如果出现其它错误
		return "", err
	}
	return val, nil
}

// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "rank"
	// value
	items := []redis.Z{
		{Score: 90, Member: "PHP"},
		{Score: 96, Member: "Golang"},
		{Score: 97, Member: "Python"},
		{Score: 99, Member: "Java"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 把元素都追加到key中
	err := rdb.ZAdd(ctx, zsetKey, items...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success.")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang score is %f onw.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// scanKeyDemo1 按前缀扫描所有key
func scanKeyDemo1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = rdb.Scan(ctx, cursor, "k*", 0).Result() // 扫描以k开头的所有key
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key:", key)
		}
		if cursor == 0 { // 如果没有keys，就跳出循环
			break
		}
	}
}

// scanKeyDemo2 按前缀扫描key
func scanKeyDemo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 按前缀扫描key
	iter := rdb.Scan(ctx, 0, "k*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err: %v\n", err)
		return
	}

	// doCommand()

	// doDemo()

	// 查询的key不存在
	// ret, err := getValueFromRedis("kkey", "没有值")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("返回的结果是:%v\n", ret)

	// 查询的key存在
	// ret, err = getValueFromRedis("key", "没有值")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("返回的结果是:%v\n", ret)

	// zsetDemo()

	// scanKeyDemo1()
	scanKeyDemo2()
}
