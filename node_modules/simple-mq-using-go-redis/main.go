package main

import (
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	queue := &Queue{
		pool: pool,
	}

	msg := &Message{
		name: "demoQueue",
	}

	queue.InitReceiver(msg)

	go func() {
		for i := 0; i < 10; i++ {
			msg := &Message{
				name: "demoQueue",
				Content: map[string]string{
					"order_no": strconv.FormatInt(time.Now().Unix(), 10),
				},
			}
			_ = queue.Delivery(msg)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT)

	for {
		switch <-quit {
		case syscall.SIGINT:
			return
		}
	}
}
