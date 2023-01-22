package complex

import pb "github/ozzy/grpcgo/protoOut"

func DoComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   42,
			Name: "Ozzy Osbourne",
		},
		Dummies: []*pb.Dummy{{Id: 43, Name: "Dave Mustang"}, {Id: 44, Name: "Tom Araya"}, {Id: 45, Name: "SlipKnot"}},
	}
}
