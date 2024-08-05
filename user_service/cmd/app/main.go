package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"user-service/cmd/app/client"
	"user-service/cmd/app/server"
	"user-service/config"
	"user-service/internal/constant"
	"user-service/internal/database"
	"user-service/pkg/llog"

	"github.com/gofiber/fiber/v2"
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
	defer db.Close()

	gRPCConnection, err := client.NewGRPCClient(config.GRPC)
	if err != nil {
		llog.Fatal(err)
	}
	defer gRPCConnection.BookService.Close()
	defer gRPCConnection.AuthorService.Close()
	defer gRPCConnection.CategoryService.Close()

	srv := server.New(db, fileLog, gRPCConnection).Setup()

	listenWithGracefulShutdown(srv)
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
