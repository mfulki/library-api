package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	srv := server.New(db, fileLog).Setup()

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
