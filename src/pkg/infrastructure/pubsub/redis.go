package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cry999/pm-projects/pkg/domain/event"
	"github.com/cry999/pm-projects/pkg/interfaces/logger"
	"github.com/cry999/pm-projects/pkg/interfaces/pubsub"
	"github.com/go-redis/redis/v8"
)

// EventMessage ...
type EventMessage struct {
	EventID string      `json:"event_id"`
	Payload interface{} `json:"payload"`
}

// RedisEventBus ...
type RedisEventBus struct {
	cli         *redis.Client
	subscribers map[string][]pubsub.Subscriber
	logger      logger.Logger
	middlewares []pubsub.Middleware
}

// NewRedisEventBus ...
func NewRedisEventBus() (*RedisEventBus, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("PUBSUB_HOST"), os.Getenv("PUBSUB_PORT")),
	})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}
	return &RedisEventBus{
		cli:         cli,
		subscribers: map[string][]pubsub.Subscriber{},
		logger:      logger.NewDefaultLogger(logger.LoggerLevelDebug, "redis.pubsub"),
	}, nil
}

// GlobalUse ...
func (bus *RedisEventBus) GlobalUse(middlewares ...pubsub.Middleware) {
	bus.middlewares = append(bus.middlewares, middlewares...)
}

// Publish ...
func (bus *RedisEventBus) Publish(ctx context.Context, event event.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		bus.logger.Error("failed to serialize event: %v", err)
		return err
	}
	if err := bus.cli.Publish(ctx, event.Type(), data).Err(); err != nil {
		bus.logger.Error("failed to publish: %v", err)
		return err
	}
	return nil
}

// Subscribe ...
func (bus *RedisEventBus) Subscribe(event event.Event, subscriber pubsub.Subscriber) {
	key := event.Type()
	for _, mw := range bus.middlewares {
		subscriber = mw(subscriber)
	}
	bus.subscribers[key] = append(bus.subscribers[key], subscriber)
}

// Start ...
func (bus *RedisEventBus) Start(ctx context.Context) error {
	channels := []string{}
	for channel := range bus.subscribers {
		channels = append(channels, channel)
	}

	var (
		wg  sync.WaitGroup
		sub = bus.cli.Subscribe(ctx, channels...)
	)
	for msg := range sub.Channel() {
		subscribers := bus.subscribers[msg.Channel]
		for _, subscriber := range subscribers {
			wg.Add(1)
			go func(subscriber pubsub.Subscriber) {
				wg.Done()
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				if err := subscriber(ctx, strings.NewReader(msg.Payload)); err != nil {
					bus.logger.Error("failed to subscriber: %v", err)
				}
			}(subscriber)
		}
	}
	wg.Wait()
	return nil
}

// Close ...
func (bus *RedisEventBus) Close() error {
	return bus.cli.Close()
}
