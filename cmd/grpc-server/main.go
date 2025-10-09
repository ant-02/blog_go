package main

import (
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"zll.blog.com/internal/config"
	db "zll.blog.com/internal/database"
	"zll.blog.com/internal/server"
	"zll.blog.com/pkg/logger"
)

func main() {
	logPath, _ := os.Getwd()
	logPath = filepath.Dir(filepath.Dir(logPath))
	config := config.GetConfig()
	logger.InitLogger(config.Log.Level, logPath+config.Log.Path)

	dsn := config.Database.Username + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := db.InitDatabase(dsn, config.Database.Dbname)
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server.RegisterAllServer(grpcServer, db)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
