package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"math/big"
	"time"
)

type IMessage interface {
	Resolve() error
	GetChannel() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) (IMessage, error)
}

type Message struct {
	name    string            // 投递的目标名称
	Content map[string]string `json:"content"` //要序列化的消息内容
}

func (m *Message) GetChannel() string {
	return m.name
}

func (m *Message) Resolve() error {
	// actually can be massive hard
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	if n.Int64()%2 == 0 {
		fmt.Printf("consumed:%v", m.Content)
		time.Sleep(time.Second)
		return nil
	}
	return errors.New("consumed failed")
}

func (m *Message) Marshal() ([]byte, error) {
	return jsoniter.Marshal(m)
}

func (m *Message) Unmarshal(reply []byte) (IMessage, error) {
	var msg Message
	err := jsoniter.Unmarshal(reply, &msg)

	return &msg, err
}
