package main

import (
	"context"
	"fmt"
	"time"

	"redis-pub-sub/rdb"
)

func main() {
	ctx := context.Background()
	

	// 每隔一秒发布一条消息
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		msg := fmt.Sprintf("message-%d", time.Now().Unix())
		if _, err := rdb.PublishMessageWithXAdd(ctx, "stream", msg); err != nil {
			panic(err)
		}
		fmt.Println("publish message:", msg)
	}
}

