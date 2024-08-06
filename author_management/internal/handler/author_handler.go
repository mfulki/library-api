package handler

import (
	"author-service/internal/dto/response"
	pb "author-service/internal/pb/author"
	"author-service/internal/usecase"
	"context"
)

type AuthorHandler struct {
	authorUsecase usecase.AuthorUsecase
	pb.UnimplementedAuthorServiceServer
}

func NewAuthorHandler(authorUsecase usecase.AuthorUsecase) *AuthorHandler {
	return &AuthorHandler{
		authorUsecase: authorUsecase,
	}
}

func (h *AuthorHandler) GetAuthors(ctx context.Context, in *pb.Empty) (*pb.Authors, error) {
	authors, err := h.authorUsecase.GetAllAuthor(ctx)
	if err != nil {
		return nil, err
	}
	resp := response.NewAuthorsResponse(authors)
	return resp, nil
}

func (h *AuthorHandler) GetSomeAuthor(ctx context.Context, in *pb.Ids) (*pb.Authors, error) {
	authors, err := h.authorUsecase.GetSomeAuthor(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	resp := response.NewAuthorsResponse(authors)
	return resp, nil

}

func (h *AuthorHandler) GetAuthor(xtx context.Context, in *pb.Id) (*pb.Author, error) {
	return nil, nil
}

func (h *AuthorHandler) InsertAuthor(ctx context.Context, in *pb.Author) (*pb.Message, error) {
	return nil, nil
}

func (h *AuthorHandler) DeleteOneAuthor(context.Context, *pb.Id) (*pb.Message, error) {
	return nil, nil
}
