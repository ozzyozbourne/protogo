package maps

import pb "github/ozzy/grpcgo/protoOut"

func DoMap() *pb.MapExample {
	return &pb.MapExample{Ids: map[string]*pb.IdWrapper{
		"myId1": {Id: 42},
		"myId2": {Id: 43},
		"myId3": {Id: 44},
	}}
}
