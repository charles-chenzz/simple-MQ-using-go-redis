package main

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Queue struct {
	pool *redis.Pool
}

// 用于删除执行队列中的消息
func (q *Queue) lrem(queue string, reply interface{}) error {
	conn := q.pool.Get()
	defer conn.Close()
	if _, err := conn.Do("LREM", queue, 1, reply); err != nil {
		fmt.Println("failed to lerm", err)
		return err
	}
	return nil
}

// 读取消息并反序列化为消息结构
func (q *Queue) rpoplpush(imsg IMessage, sourceQueue, destQueue string) (interface{}, IMessage, error) {
	conn := q.pool.Get()
	defer conn.Close()

	r, err := conn.Do("RPOPLPUSH", sourceQueue, destQueue)
	if err != nil {
		return nil, nil, err
	}

	if r == nil {
		return nil, nil, nil
	}

	// 断言为[]uint8 type
	rUint8, ok := r.([]uint8)
	if !ok {
		return nil, nil, errors.New("can't assert reply as type uint8")
	}
	if msg, err := imsg.Unmarshal(rUint8); err != nil {
		return nil, nil, err
	} else if _, ok := msg.(IMessage); ok {
		return r, msg, nil
	} else {
		return nil, nil, errors.New("can't assert msg as interface IMessage")
	}
}

func (q *Queue) Delivery(msg IMessage) error {
	conn := q.pool.Get()
	defer conn.Close()

	// 待执行队列后面加 .prepare
	prepareQueue := fmt.Sprintf("%s.prepare", msg.GetChannel())
	if msgJson, err := msg.Marshal(); err != nil {
		return err
	} else {
		_, err := conn.Do("LPUSH", prepareQueue, msgJson)
		fmt.Println("produced", string(msgJson))
		return err
	}
}

func (q *Queue) InitReceiver(msg IMessage) {
	// 投递目标名称后缀加 .prepare 用于表示执行队列
	prepareQueue := fmt.Sprintf("%s.prepare", msg.GetChannel())

	// 投递目标名称后缀加 .doing 用于表示执行队列
	doingQueue := fmt.Sprintf("%s.doing", msg.GetChannel())

	go func() {
		for {
			q.ack(msg, doingQueue, prepareQueue)

			reply, msg, err := q.rpoplpush(msg, prepareQueue, doingQueue)
			if err != nil {
				fmt.Println("failed to pop msg", err)
				continue
			}
			if msg == nil {
				continue
			}
			if err := msg.Resolve(); err == nil {
				q.lrem(doingQueue, reply)
			} else {
				fmt.Println(err)
			}
		}
	}()

	fmt.Printf("receiver have been initialized")
}

func (q Queue) ack(imsg IMessage, sourceQueue, destQueue string) {
	for {
		reply, _, err := q.rpoplpush(imsg, sourceQueue, destQueue)
		if err != nil {
			fmt.Println("ack failed", err)
			break
		}
		if reply == nil {
			// 空消息 说明无滞留 打断循环
			break
		} else {
			fmt.Printf("got undo msg in the queue %sn", sourceQueue)
		}
	}
}
