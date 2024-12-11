package rdb

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var rc *redis.Client
var once sync.Once

func GetRedisClient() *redis.Client {
	once.Do(func() {
		rc = redis.NewClient(&redis.Options{
			Addr: ":6379",
			DB: 9,
		})
	})
	return rc
}

// publist message with xadd
func PublishMessageWithXAdd(ctx context.Context, stream, message string) (string, error) {
	client := GetRedisClient()
	
	return client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{"message": message},
		MaxLen: 5000, // 最多保留5000条消息
	}).Result()
}

// subscribe message with xreadgroup
func SubscribeMessageWithXReadGroup(ctx context.Context, group, consumer, stream string) ([]redis.XStream, error) {
	client := GetRedisClient()
	
	return client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumer,
		Streams:  []string{stream, ">"},
		Block:    0,
		Count:    1,
	}).Result()
}

// ConsumerGroupExists 检查消费者组是否存在
func ConsumerGroupExists(ctx context.Context, stream, group string) (bool, error) {
    client := GetRedisClient()
    
    groups, err := client.XInfoGroups(ctx, stream).Result()
    if err != nil {
        return false, err
    }

    for _, g := range groups {
        if g.Name == group {
            return true, nil
        }
    }
    return false, nil
}

// CreateConsumerGroup 创建消费者组
func CreateConsumerGroup(ctx context.Context, stream, group string) error {
	client := GetRedisClient()
	
	return client.XGroupCreateMkStream(ctx, stream, group, "$").Err()
}

// 确认消息
func AckMessage(ctx context.Context, stream, group, id string) error {
	client := GetRedisClient()
	
	return client.XAck(ctx, stream, group, id).Err()
}

// GetPendingMessages 获取未确认的消息
func GetPendingMessages(ctx context.Context, stream, group, consumer string) ([]redis.XMessage, error) {
    client := GetRedisClient()

    // 获取未确认的消息
    pending, err := client.XPendingExt(ctx, &redis.XPendingExtArgs{
        Stream:   stream,
        Group:    group,
        Start:    "-",
        End:      "+",
        Count:    10,
        Consumer: consumer,
    }).Result()
    if err != nil {
        return nil, err
    }

    var messages []redis.XMessage
    for _, p := range pending {
        msgs, err := client.XClaim(ctx, &redis.XClaimArgs{
            Stream:   stream,
            Group:    group,
            Consumer: consumer,
            MinIdle:  0,
            Messages: []string{p.ID},
        }).Result()
        if err != nil {
            return nil, err
        }
        messages = append(messages, msgs...)
    }

    return messages, nil
}