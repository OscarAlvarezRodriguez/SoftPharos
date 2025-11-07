package databases

import (
	"sync"
)

var (
	instance *Client
	once     sync.Once
)

func InitializeDatabase(client *Client) {
	once.Do(func() {
		instance = client
	})
}

func GetInstance() *Client {
	if instance == nil {
		panic("Database not initialized. Call InitializeDatabase first")
	}
	return instance
}

func CloseInstance() {
	if instance != nil {
		instance.Close()
	}
}
