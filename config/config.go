package config

import (
	"fmt"
	"net"
	"os"
)

func InitListener() net.Listener {
	listener, err := net.Listen("tcp", ":"+os.Getenv("APP_PORT"))
	if err != nil {
		panic(fmt.Errorf("failed listening: %w", err))
	}

	return listener
}
