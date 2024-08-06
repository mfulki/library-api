package main

import (
	"category-service/cmd/app/server"
	"category-service/config"
	"category-service/internal/constant"
	"category-service/internal/database"
	"category-service/pkg/llog"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

const (
	shutdownMsg = "Shutdown server..."
	exitedMsg   = "Server exited..."
	timeoutMsg  = "Timeout of %v"
)

func main() {
	config.Load()

	fileLog, err := llog.NewFileLogger("app")
	if err != nil {
		llog.Fatal(err)
	}

	db, err := database.NewPostgress(config.DB)
	if err != nil {
		llog.Fatal(err)
	}
	netListen, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	}

	grpcServer := grpc.NewServer()

	log.Printf("Server started at %v", netListen.Addr())

	srv := server.New(db, fileLog, grpcServer).Setup()

	if err := srv.Serve(netListen); err != nil {
		log.Fatalf("failed to serve %v", err.Error())
	}

	// grpcServer.GracefulStop()
	// listenWithGracefulShutdown(srv)
}

func listenWithGracefulShutdown(srv *fiber.App) {
	var asyncListen = func() {
		if err := srv.Listen(fmt.Sprintf(":%d", config.App.Port)); err != nil {
			llog.Fatal(err)
		}
	}

	go asyncListen()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	llog.Print(shutdownMsg)
	ctx, cancel := context.WithTimeout(context.Background(), constant.TimeoutShutdown)
	defer cancel()

	if err := srv.ShutdownWithContext(ctx); err != nil {
		llog.Fatal(err)
	}

	select {
	case <-ctx.Done():
		llog.Print(fmt.Sprintf(timeoutMsg, constant.TimeoutShutdown))
	default:
		llog.Print(exitedMsg)
	}
}
