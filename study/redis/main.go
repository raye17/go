package main

import (
	"context"
	"encoding/json"
	"fmt"
	"study/redis/cache"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	KeyPrefixTaskGroup = "infra:task_group:"
)

func TaskGroupKey(groupID string) string {
	return fmt.Sprintf("%s%s", KeyPrefixTaskGroup, groupID)
}

type TaskGroupCache struct {
	client *RedisClient
}
type RedisClient struct {
	redisClient *redis.Client
}

func NewTaskGroupCache(client *RedisClient) *TaskGroupCache {
	return &TaskGroupCache{
		client: client,
	}
}

type TaskGroupCacheData struct {
	NodeID    int64     `json:"nodeId"`
	NodeIP    string    `json:"nodeIp"`
	Port      int       `json:"port"`
	Status    int8      `json:"status"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (c *TaskGroupCache) Set(ctx context.Context, groupID string, data *TaskGroupCacheData, ttl time.Duration) error {
	key := TaskGroupKey(groupID)
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.client.redisClient.Set(ctx, key, bytes, ttl).Err()
}
func (c *TaskGroupCache) Get(ctx context.Context, groupID string) (*TaskGroupCacheData, error) {
	key := TaskGroupKey(groupID)
	bytes, err := c.client.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	var data TaskGroupCacheData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func main() {
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now())
	//return
	c := cache.NewClientRedis()
	if err := c.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("redis connect failed!:%v", err)
	}
	rc := &RedisClient{
		redisClient: c,
	}
	gc := NewTaskGroupCache(rc)
	groupId := uuid.New().String()
	data := &TaskGroupCacheData{
		NodeID:    1,
		NodeIP:    "192.168.1.2",
		Port:      8080,
		Status:    1,
		ExpiredAt: time.Now().Add(time.Hour * 1),
	}
	if err := gc.Set(context.Background(), groupId, data, time.Hour*1); err != nil {
		fmt.Printf("set task group cache failed!:%v", err)
	}
	d, err := gc.Get(context.Background(), groupId)
	if err != nil {
		fmt.Printf("get task group cache failed!:%v", err)
	}
	fmt.Println("task group cache:", d)
}
