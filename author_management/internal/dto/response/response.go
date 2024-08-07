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

func NewAuthorsBooksResp(mapAB map[uint64]entity.AuthorsBooksJson) *pb.AuthorsBooksMap {
	mapAuthorBooks := make(map[uint64]*pb.AuthorsBooks, 0)
	for key, val := range mapAB {
		sliceAuthorBook := []*pb.AuthorsBook{}
		for _, authorsBook := range val {
			pbAuthorBook := pb.AuthorsBook{Id: authorsBook.AuthorId, Name: authorsBook.AuthorName}
			sliceAuthorBook = append(sliceAuthorBook, &pbAuthorBook)
		}
		AuthorBooks := pb.AuthorsBooks{AuthorBookList: sliceAuthorBook}
		mapAuthorBooks[key] = &AuthorBooks
	}
	return &pb.AuthorsBooksMap{
		AuthorBooksMap: mapAuthorBooks,
	}
}
