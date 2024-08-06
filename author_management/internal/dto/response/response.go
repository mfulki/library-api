package response

import (
	"author-service/internal/entity"
	pb "author-service/internal/pb/author"
)

func NewAuthorsResponse(authors []entity.Author) *pb.Authors {
	var resAuthors []*pb.Author
	for _, author := range authors {
		resAuthor := &pb.Author{
			Id:         author.Id,
			AuthorName: author.Name,
			PhotoUrl:   author.PhotoUrl,
			Gender:     author.Gender,
		}
		resAuthors = append(resAuthors, resAuthor)
	}
	return &pb.Authors{Author: resAuthors}
}
