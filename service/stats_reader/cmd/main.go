package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/handler"
	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/clickhouse"
	"github.com/MamangRust/microservice-ecommerce-pkg/dotenv"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := dotenv.Viper(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
	}
	log, _ := logger.NewLogger("stats-reader", nil)

	// The ClickHouse database must exist before NewClient can ping.
	if err := clickhouse.EnsureDatabase(log); err != nil {
		log.Fatal("Failed to ensure ClickHouse database", zap.Error(err))
	}
	chConn, err := clickhouse.NewClient(log)
	if err != nil {
		log.Fatal("Failed to connect to ClickHouse", zap.Error(err))
	}

	repo := repository.NewRepository(chConn)

	categoryStatsHandler := handler.NewCategoryStatsHandler(repo, log)
	orderStatsHandler := handler.NewOrderStatsHandler(repo, log)
	transactionStatsHandler := handler.NewTransactionStatsHandler(repo, log)

	grpcServer := grpc.NewServer()

	pb.RegisterCategoryStatsServiceServer(grpcServer, categoryStatsHandler)
	pb.RegisterCategoryStatsByIdServiceServer(grpcServer, categoryStatsHandler)
	pb.RegisterCategoryStatsByMerchantServiceServer(grpcServer, categoryStatsHandler)
	pb.RegisterOrderStatsServiceServer(grpcServer, orderStatsHandler)
	pb.RegisterTransactionStatsServiceServer(grpcServer, transactionStatsHandler)
	pb.RegisterTransactionStatsByMerchantServiceServer(grpcServer, transactionStatsHandler)

	reflection.Register(grpcServer)

	port := ":50070"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen", zap.Error(err))
	}

	log.Info("Stats Reader starting", zap.String("port", port))

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("Failed to serve", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down Stats Reader...")
	grpcServer.GracefulStop()
}
