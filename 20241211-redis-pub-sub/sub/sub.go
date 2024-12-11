package main

import (
	"context"
	"fmt"
	"log"

	"redis-pub-sub/rdb"
)

func main() {
	ctx := context.Background()

	// consumer group
	streamName := "stream"
	groupName := "group"
	consumerName := "consumer"

	// 检查消费者组是否存在
	groupExists, err := rdb.ConsumerGroupExists(ctx, streamName, groupName)
	if err != nil {
		log.Println(err)
	}

	// 如果消费者组不存在，则创建
	if !groupExists {
		if err := rdb.CreateConsumerGroup(ctx, streamName, groupName); err != nil {
			panic(err)
		}
	}

	// 处理未确认的消息
    pendingMessages, err := rdb.GetPendingMessages(ctx, streamName, groupName, consumerName)
    if err != nil {
        panic(err)
    }

    for _, message := range pendingMessages {
        payload, ok := message.Values["message"].(string)
        if ok {
            fmt.Println("pending message:", payload)
            // 确认消息
            if err := rdb.AckMessage(ctx, streamName, groupName, message.ID); err != nil {
                panic(err)
            }
        }
    }

	for {
		// 订阅消息
		streams, err := rdb.SubscribeMessageWithXReadGroup(ctx, groupName, consumerName, streamName)
		if err != nil {
			panic(err)
		}

		// 打印消息
		for _, stream := range streams {
			for _, message := range stream.Messages {
				payload, ok := message.Values["message"].(string)
				if !ok {
					continue
				}
				fmt.Println("message:", payload)
				// 确认消息
				if err := rdb.AckMessage(ctx, "stream", "group", message.ID); err != nil {
					panic(err)
				}
			}
		}
	}
}
