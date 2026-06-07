package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/loanem-backend/participant-service/config"
	"github.com/loanem-backend/participant-service/internal/messenger"
	"github.com/loanem-backend/participant-service/internal/server"
	"github.com/loanem-backend/participant-service/internal/service"
	"google.golang.org/grpc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	godotenv.Load()

	db := config.InitDB()
	defer db.Close()

	amqpConn, amqpCh := config.InitAMQPChannel()
	defer amqpCh.Close()
	defer amqpConn.Close()

	s := grpc.NewServer()

	crs, cls, ts := service.Initialize(db)

	server.Start(s, cls, ts)

	if err := messenger.Start(ctx, amqpCh, crs); err != nil {
		log.Fatalf("failed starting service: %v\n", err)
	}

	lis := config.InitListener()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed serving grpc: %n\n", err)
	}

	<-ctx.Done()
	log.Println("Stopping server gracefully...")
}
