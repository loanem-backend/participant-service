package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/loanem-backend/participant-service/config"
	"github.com/loanem-backend/participant-service/internal/server"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	s := grpc.NewServer()

	server.Start()

	lis := config.InitListener()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed serving grpc: %n\n", err)
	}
}
