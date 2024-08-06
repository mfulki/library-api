package request

import pb "author-service/internal/pb/author"

func NewIdsRequest(in *pb.Ids) []uint64 {
	var ids []uint64
	for _, id := range in.Id {
		ids = append(ids, id)
	}
	return ids
}
