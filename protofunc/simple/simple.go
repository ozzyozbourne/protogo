package simple

import pb "github/ozzy/grpcgo/protoOut"

func DoSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
}
