package main

import (
	"fmt"
	"net"
	"position_service/config"
	"position_service/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "position_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		"disable",
	)

	// db, err := sqlx.Connect("postgres", conStr)
	// if err != nil {
	// 	log.Error("error while connecting database", logger.Error(err))
	// 	return
	// }

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
