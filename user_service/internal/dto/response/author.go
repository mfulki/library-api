package response

type AuthorBook struct {
	AuthorId   uint64 `json:"author_id"`
	AuthorName string `json:"author_name"`
}