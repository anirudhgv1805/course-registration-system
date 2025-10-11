package redis

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

type CourseCache struct {
    client *redis.Client
}

func NewCourseCache(client *redis.Client) *CourseCache {
    return &CourseCache{client: client}
}

const seatBlockTTL = 2 * time.Minute

func (c *CourseCache) BlockSeat(ctx context.Context, courseID, userID string) (bool, error) {
    key := fmt.Sprintf("seat:block:%s:%s", courseID, userID)
    return c.client.SetNX(ctx, key, "1", seatBlockTTL).Result()
}

func (c *CourseCache) IsSeatBlocked(ctx context.Context, courseID, userID string) (bool, error) {
    key := fmt.Sprintf("seat:block:%s:%s", courseID, userID)
    exists, err := c.client.Exists(ctx, key).Result()
    return exists > 0, err
}

func (c *CourseCache) UnblockSeat(ctx context.Context, courseID, userID string) error {
    key := fmt.Sprintf("seat:block:%s:%s", courseID, userID)
    return c.client.Del(ctx, key).Err()
}
