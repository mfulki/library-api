package example

type StoreRequest struct {
	Number int `json:"number" validate:"required,gte=5"`
}

func (req *StoreRequest) Example() Entity {
	return Entity{
		ID: uint(req.Number),
	}
}
