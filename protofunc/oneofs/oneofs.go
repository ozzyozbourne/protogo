package oneofs

import (
	"fmt"
	pb "github/ozzy/grpcgo/protoOut"
)

func DoOneOfs(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		_ = fmt.Errorf("message has unexpected type: %v", x)
	}
}
