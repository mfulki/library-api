package handler

import (
	"context"
	"log"
	"time"

	"user-service/internal/constant"
	"user-service/internal/dto"
	pb "user-service/internal/pb/books"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookHandler struct {
}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) GetAllBook(ctx *fiber.Ctx) error {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logrus.Errorf("did not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := pb.NewBookServiceClient(conn)
	context, cancel := context.WithTimeout(ctx.Context(), time.Second)
	defer cancel()

	response, err := c.GetBooks(context, &pb.Empty{})
	if err != nil {
		logrus.Errorf("could not request: %v", err)
		return err
	}
	log.Printf("Response: %s", response)

	return ctx.Status(fiber.StatusOK).JSON(dto.Response{
		Message: constant.DataRetrievedMsg,
		Data:    response,
	})
}
