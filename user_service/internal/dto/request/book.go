package request

type BookId struct {
	Id uint64 `params:"id"`
}

type BookIds struct {
	Ids []uint64 `query:"ids",json:"required"`
}
